// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"context"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/crossplane/crossplane-runtime/v2/pkg/logging"
	"github.com/pkg/errors"
	"golang.org/x/sync/singleflight"

	"github.com/upbound/provider-aws/v2/apis/namespaced/v1beta1"
)

const (
	errGetAccountID = "cannot retrieve the AWS account ID"

	// defaultCacheMaxSize is the default maximum number of entries in the
	// AWS credentials provider cache.
	defaultCacheMaxSize = 2000

	// Reasons for which a credential provider is not cached. They are logged
	// as the "reason" of a "Not caching the credential provider" debug line,
	// so keep them stable: they are meant to be filtered and counted on.
	skipNotACredentialsCache     = "Configured aws.CredentialsProvider is not an aws.CredentialsCache"
	skipSourceNotCacheable       = "CredentialsSource is not cacheable"
	skipNoCredentialsFingerprint = "No credentials fingerprint"
)

// AWSCredentialsProviderCacheOption lets you configure
// a *GlobalAWSCredentialsProviderCache.
type AWSCredentialsProviderCacheOption func(cache *AWSCredentialsProviderCache)

// WithCacheMaxSize lets you override the default MaxSize for
// AWS CredentialsProvider cache.
func WithCacheMaxSize(n int) AWSCredentialsProviderCacheOption {
	return func(c *AWSCredentialsProviderCache) {
		c.maxSize = n
	}
}

// WithCacheStore lets you bootstrap AWS CredentialsProvider Cache with
// your own cache. It must be keyed the way RetrieveCredentials keys it, i.e.
// by the ProviderConfig's UID and the region, and the supplied entries must
// carry everything RetrieveCredentials serves from one, i.e. an account ID, an
// access time and the version of the credential material, as the cache only
// ever holds fully initialized entries.
func WithCacheStore(cache map[string]*awsCredentialsProviderCacheEntry) AWSCredentialsProviderCacheOption {
	return func(c *AWSCredentialsProviderCache) {
		c.cache = cache
	}
}

// WithCacheLogger lets you configure the logger for the cache.
func WithCacheLogger(l logging.Logger) AWSCredentialsProviderCacheOption {
	return func(c *AWSCredentialsProviderCache) {
		c.logger = l
	}
}

// NewAWSCredentialsProviderCache returns a new empty
// *AWSCredentialsProviderCache with the default GetAWSConfig method.
func NewAWSCredentialsProviderCache(opts ...AWSCredentialsProviderCacheOption) *AWSCredentialsProviderCache {
	c := &AWSCredentialsProviderCache{
		cache: map[string]*awsCredentialsProviderCacheEntry{},
		// entries are keyed by the provider config and the region, and a
		// rotation of the credentials replaces an entry instead of adding one,
		// so the count is bounded by the number of provider configs times the
		// number of regions in use. Only the entries of the provider configs
		// that are gone, or of the regions that stopped being reconciled, are
		// left for this ceiling to reclaim. Evicting an entry means paying for
		// its STS calls again, and the entries are small, so keep it generous.
		maxSize: defaultCacheMaxSize,
		mu:      &sync.RWMutex{},
		logger:  logging.NewNopLogger(),
	}
	for _, f := range opts {
		f(c)
	}
	return c
}

// AWSCredentialsProviderCache holds aws.CredentialsProvider objects in memory
// so that we don't need to make API calls to AWS in every reconciliation of
//
//	every resource. It has a maximum size that when it's reached, the entry
//	that has the oldest access time will be removed from the cache,
//	i.e. FIFO on last access time.
//
// The credentials an entry holds never change, so a served entry never needs
// to be invalidated. A ProviderConfig does however produce a new version of
// them whenever it is edited or its credential material rotates, and the cache
// holds a single entry per ProviderConfig and region, so the new version
// replaces the old one in place. Replacing swaps in a whole new entry rather
// than mutating the stored one, hence it does not invalidate the entry for
// whoever is already holding it.
type AWSCredentialsProviderCache struct {
	// cache holds one entry per ProviderConfig and region, keyed by the
	// ProviderConfig's UID and the region. Everything else that the resulting
	// credentials depend on, i.e. the ProviderConfig's generation, its
	// credential source and a fingerprint of the credential material, is
	// carried in the entry instead of in the key: those designate the version
	// of the credentials, of which only one can ever be current, so a new
	// version replaces its predecessor rather than accumulating next to it.
	// Finding an entry is therefore not enough to serve it, see
	// awsCredentialsProviderCacheEntry.matches.
	//
	// Which credential sources are cached at all is decided by
	// cacheableSource.
	cache map[string]*awsCredentialsProviderCacheEntry

	// maxSize is the maximum number of elements this cache can ever have.
	maxSize int

	// mu is used to make sure the cache map is concurrency-safe. It is never
	// held across an AWS API call, see RetrieveCredentials.
	mu *sync.RWMutex

	// inflight collapses the concurrent initializations of one version of the
	// credentials into a single one, so that a rotation observed by many
	// reconciliations at once costs one set of AWS API calls instead of one
	// per reconciliation. It is keyed per version rather than per cache slot,
	// see RetrieveCredentials.
	inflight singleflight.Group

	// logger is the logger for cache operations.
	logger logging.Logger
}

type awsCredentialsProviderCacheEntry struct {
	awsCredCache *aws.CredentialsCache
	accessedAt   atomic.Value
	accountID    atomic.Value

	// generation, source and credsFingerprint designate the version of the
	// credential material this entry was built from. They are immutable for
	// the lifetime of the entry: a new version replaces the whole entry
	// instead of mutating this one, so that the readers holding it, which do
	// so without the cache's lock, always observe a consistent version.
	generation       int64
	source           string
	credsFingerprint string
}

// matches reports whether the entry was built from the credential material
// that the supplied ProviderConfig and provenance designate. The cache holds a
// single entry per ProviderConfig and region, so finding one is not enough to
// serve it: it may belong to a superseded generation, credential source or
// credential material. Every read of a stored entry must go through here, the
// lookups in RetrieveCredentials included.
func (e *awsCredentialsProviderCacheEntry) matches(pc *v1beta1.ClusterProviderConfig, cfgMeta awsConfigProvenanceMeta) bool {
	return e.generation == pc.Generation &&
		e.source == string(pc.Spec.Credentials.Source) &&
		e.credsFingerprint == cfgMeta.credsFingerprint
}

// AccountIDFn is a function for retrieving the account ID.
type AccountIDFn func(ctx context.Context) (string, error)

func accountIDFromCacheEntry(e *awsCredentialsProviderCacheEntry) AccountIDFn {
	return func(context.Context) (string, error) {
		// return the cached account ID
		return e.accountID.Load().(string), nil
	}
}

// Credentials holds the aws.Credentials and the associated AWS account ID for
// these credentials. It's possible that the account ID is not resolved and
// only the aws.Credentials are available in a successful result.
type Credentials struct {
	creds     aws.Credentials
	accountID string
}

// newCredentials returns the Credentials whose credentials are retrieved
// using the given aws.CredentialsProvider and whose account ID is set using
// the given AccountIDFn.
func newCredentials(ctx context.Context, credsProvider aws.CredentialsProvider, accountIDFn AccountIDFn) (Credentials, error) {
	var result Credentials
	// try to retrieve the credentials if a retriever has been supplied
	if credsProvider != nil {
		var err error
		if result.creds, err = credsProvider.Retrieve(ctx); err != nil {
			return Credentials{}, errors.Wrap(err, "cannot retrieve the AWS credentials")
		}
	}
	// try to get the account ID
	if accountIDFn != nil {
		var err error
		if result.accountID, err = accountIDFn(ctx); err != nil {
			return Credentials{}, errors.Wrap(err, errGetAccountID)
		}
	}
	return result, nil
}

// cacheableSource reports whether the credentials resulting from the supplied
// ProviderConfig can and are worth being cached.
func cacheableSource(pc *v1beta1.ClusterProviderConfig) bool {
	switch pc.Spec.Credentials.Source {
	case authKeyIRSA:
		return true
	case authKeyWebIdentity:
		// Retrieving the credentials calls sts:AssumeRoleWithWebIdentity, so
		// caching pays off regardless of whether a role chain is configured on
		// top. The token, which is the only credential material outside of the
		// spec here, is covered by fingerprintIdentityToken.
		return true
	case authKeyPodIdentity, authKeyUpbound:
		// TODO: these authentication methods are not supported by the cache
		// yet. They need their own out-of-spec key material, similar to the
		// IRSA token hash.
		return false
	default:
		// Static credential sources, i.e. Secret, Fs and Environment, involve
		// no AWS API call while retrieving the credentials, so caching them
		// only pays off when a role chain needs to be assumed on top of them.
		return len(pc.Spec.AssumeRoleChain) > 0
	}
}

// RetrieveCredentials returns a Credentials either from the credential cache.
// If the authentication scheme is cacheable, i.e. IRSA, WebIdentity or a static
// credential source with an assume role chain, and the supplied aws.CredentialsProvider
// implementation is an aws.CredentialsCache, then the retrieved credentials and
// the account ID are cached for future requests.
// Otherwise, this function returns the AWS credentials by calling
// the downstream aws.CredentialsProvider.Retrieve, and for now, does *not*
// call the given AccountIDFn because in that case, a separate identity cache
// should be used to retrieve the caller identity.
func (c *AWSCredentialsProviderCache) RetrieveCredentials(ctx context.Context, pc *v1beta1.ClusterProviderConfig, region string, credsProvider aws.CredentialsProvider, cfgMeta awsConfigProvenanceMeta, accountIDFn AccountIDFn) (Credentials, error) { //nolint:gocyclo // mostly the cache key calculation
	// Only aws.CredentialsCache is supported as the underlying credential
	// provider, as the whole point of the cache is to keep the SDK's own
	// credential cache object, which refreshes the credentials as they expire,
	// alive across reconciliations.
	awsCredsCache, isCredsCache := credsProvider.(*aws.CredentialsCache)
	var skipReason string
	switch {
	case !isCredsCache:
		skipReason = skipNotACredentialsCache
	case !cacheableSource(pc):
		skipReason = skipSourceNotCacheable
	case cfgMeta.credsFingerprint == "":
		// The credential material lives outside of the ProviderConfig, e.g. in
		// a Secret, so its generation alone cannot detect a rotation. Without a
		// fingerprint of the material we cannot construct a safe cache key,
		// hence skip the cache instead of risking serving credentials derived
		// from stale material.
		skipReason = skipNoCredentialsFingerprint
	}
	if skipReason != "" {
		c.logger.Debug("Cannot utilize the provider credential cache",
			"reason", skipReason,
			"source", string(pc.Spec.Credentials.Source),
			"providerConfigName", pc.Name, "providerConfigUID", string(pc.UID))
		// if this cache manager is not going to be employed, do not call
		// the given accountIDFn because there's a separate identity cache
		// implementation.
		// TODO: Replace the identity cache with this cache.
		return newCredentials(ctx, credsProvider, nil)
	}
	// Every parameter that could change the resulting AWS credentials has to
	// be accounted for, either by the slot key or by the version carried in
	// the entry:
	//
	//   - the parameters that live in the ProviderConfig spec are covered by
	//     its generation, which changes whenever the spec does;
	//   - the parameters that live outside of it, e.g. in a Secret or in a
	//     projected token file, produce the same generation with different
	//     credentials, so they are covered by cfgMeta.credsFingerprint;
	//   - the region and the identity of the ProviderConfig itself make up
	//     the slot key.
	//
	// The cache holds a single entry per ProviderConfig and region, hence only
	// the latter is in the key. The rest designates the version of the
	// credentials, of which only one can ever be current, so it is checked
	// against the stored entry with matches and a new version takes the slot
	// of the one it supersedes instead of accumulating next to it.
	slotKey := string(pc.UID) + ":" + region
	c.logger.Debug("Checking cache entry", "cacheKey", slotKey, "providerConfigName", pc.Name, "providerConfigUID", string(pc.UID))
	c.mu.RLock()
	cacheEntry, ok := c.cache[slotKey]
	// snapshot the size for the miss log below: it must not be read outside of
	// this critical section, as the initializations running concurrently write
	// to the map, and a concurrent map read is fatal rather than merely torn.
	cacheSize := len(c.cache)
	c.mu.RUnlock()

	// TODO: consider implementing a TTL even though the cached entry is valid
	// cache hit
	if ok && cacheEntry.matches(pc, cfgMeta) {
		c.logger.Debug("Cache hit", "cacheKey", slotKey, "providerConfigName", pc.Name, "providerConfigUID", string(pc.UID))
		// since this is a hot-path in the execution, do not always update
		// the last access times, it is fine to evict the LRU entry on a less
		// granular precision.
		if time.Since(cacheEntry.accessedAt.Load().(time.Time)) > 10*time.Minute {
			cacheEntry.accessedAt.Store(time.Now())
		}
		return newCredentials(ctx, cacheEntry.awsCredCache, accountIDFromCacheEntry(cacheEntry))
	}

	// cache miss, i.e. the slot is empty or holds a superseded version.
	// Initializing an entry costs an sts:GetCallerIdentity plus, on the first
	// retrieval, an sts:AssumeRole*, so it must not happen while holding the
	// global lock: that would serialize the reconciliations of every other
	// ProviderConfig, the cache hits included, since a pending writer also
	// locks the readers out.
	//
	// The initializations are deduplicated per version of the credentials and
	// not per slot, as two versions can legitimately be initialized at the
	// same time while a rotation is in progress, and a reconciliation must
	// never join a flight that is minting a version other than the one it
	// asked for. Note that a flight runs on a goroutine of its own and
	// captures the context, the account ID lookup and the credentials provider
	// of whichever reconciliation won the race for the version. The
	// reconciliations that join one therefore observe its outcome, including
	// its failures, rather than making the same calls themselves.
	flightKey := strings.Join([]string{
		slotKey,
		strconv.FormatInt(pc.Generation, 10),
		string(pc.Spec.Credentials.Source),
		cfgMeta.credsFingerprint, // empty fingerprints are rejected above
	}, ":")
	c.logger.Debug("Cache miss", "cacheKey", slotKey, "providerConfigName", pc.Name, "providerConfigUID", string(pc.UID), "cacheSize", cacheSize)
	inflight := c.inflight.DoChan(flightKey, func() (any, error) {
		// another reconciliation may have populated the slot with this very
		// version between our lookup above and our turn here.
		c.mu.RLock()
		entry, ok := c.cache[slotKey]
		c.mu.RUnlock()
		if ok && entry.matches(pc, cfgMeta) {
			return entry, nil
		}
		id, err := accountIDFn(ctx)
		if err != nil {
			// nothing is cached, hence the next reconciliation retries.
			return nil, errors.Wrap(err, errGetAccountID)
		}
		entry = &awsCredentialsProviderCacheEntry{
			awsCredCache:     awsCredsCache,
			generation:       pc.Generation,
			source:           string(pc.Spec.Credentials.Source),
			credsFingerprint: cfgMeta.credsFingerprint,
		}
		entry.accountID.Store(id)
		// accessedAt must be populated before the entry is published, as
		// makeRoom loads it from every entry in the map.
		entry.accessedAt.Store(time.Now())

		c.mu.Lock()
		defer c.mu.Unlock()
		existing, occupied := c.cache[slotKey]
		switch {
		case occupied && existing.generation > pc.Generation:
			// a straggling reconciliation of an older generation must not
			// replace the entry of a newer one: the generation only ever moves
			// forward, so the stored entry is the current one and ours is
			// already obsolete. Serve this reconciliation, but leave the cache
			// alone, or the two generations would evict each other for as long
			// as the stragglers keep arriving.
			c.logger.Debug("Not caching the credentials of a superseded ProviderConfig generation",
				"cacheKey", slotKey, "providerConfigName", pc.Name, "providerConfigUID", string(pc.UID),
				"generation", pc.Generation, "cachedGeneration", existing.generation)
		case occupied:
			// replacing a superseded version does not grow the cache, hence no
			// room needs to be made for it.
			c.cache[slotKey] = entry
		default:
			c.makeRoom()
			c.cache[slotKey] = entry
		}
		return entry, nil
	})
	select {
	case res := <-inflight:
		if res.Err != nil {
			return Credentials{}, res.Err
		}
		cacheEntry = res.Val.(*awsCredentialsProviderCacheEntry)
	case <-ctx.Done():
		// only this reconciliation gives up waiting. The initialization keeps
		// running and still populates the cache, unless the context it
		// captured is the one that was cancelled, in which case it fails and
		// leaves nothing behind.
		return Credentials{}, errors.Wrap(ctx.Err(), "cannot wait for the AWS credentials provider to be initialized")
	}
	return newCredentials(ctx, cacheEntry.awsCredCache, accountIDFromCacheEntry(cacheEntry))
}

// makeRoom ensures that there is at most maxSize-1 elements in the cache map
// so that a new entry can be added. It deletes the object that
// was last accessed before all others.
// This implementation is not thread safe. Callers must properly synchronize.
func (c *AWSCredentialsProviderCache) makeRoom() {
	if 1+len(c.cache) <= c.maxSize {
		return
	}
	var dustiest string
	for key, val := range c.cache {
		if dustiest == "" {
			dustiest = key
			continue
		}
		if val.accessedAt.Load().(time.Time).Before(c.cache[dustiest].accessedAt.Load().(time.Time)) {
			dustiest = key
		}
	}
	delete(c.cache, dustiest)
}

// fingerprintKey keys the credential fingerprints. It is randomly generated
// per process, so that a fingerprint is meaningless outside of the process
// that produced it: fingerprints reach the credential cache keys, which are
// logged at debug level, and an unkeyed digest of a credential would be a
// stable identifier that anyone holding a candidate credential could confirm
// by recomputing it. The credential cache never outlives the process, so the
// fingerprints need no stability across restarts.
var fingerprintKey = []byte(rand.Text())

// Domains for fingerprintMaterial. All credential sources contribute their
// fingerprint to the very same cache key component, so their digests must not
// be able to collide across sources even when they digest the same bytes.
const (
	fpDomainStaticCreds      = "static-creds"
	fpDomainWebIdentityToken = "web-identity-token"
	fpDomainIRSA             = "irsa"
)

// fingerprintMaterial returns a non-reversible, domain-separated digest of the
// supplied credential material, suitable for use as a cache key component. The
// digest is only comparable against the digests produced by the same process,
// see fingerprintKey.
func fingerprintMaterial(domain string, parts ...[]byte) string {
	h := hmac.New(sha256.New, fingerprintKey)
	for _, p := range append([][]byte{[]byte(domain)}, parts...) {
		// length-prefix framing, so that different material cannot digest to
		// the same value by concatenating to the same byte string. A mere
		// separator byte would leave the parts ambiguous if one of them ever
		// contained that byte itself.
		_ = binary.Write(h, binary.BigEndian, uint64(len(p)))
		_, _ = h.Write(p)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

// fingerprintStaticCreds returns a non-reversible digest of the supplied
// aws.Credentials.
func fingerprintStaticCreds(creds aws.Credentials) string {
	return fingerprintMaterial(fpDomainStaticCreds,
		[]byte(creds.AccessKeyID), []byte(creds.SecretAccessKey), []byte(creds.SessionToken))
}

// fingerprintIdentityToken returns a non-reversible digest of the web identity
// token that the supplied retriever currently yields.
//
// The token is read eagerly, because the cache key must be known before the
// cache is consulted, whereas the AWS SDK only calls GetIdentityToken lazily,
// on a credential refresh. Note that the retriever handed to the SDK must stay
// the live one: a cached credential provider needs to re-read the token source
// on every refresh, otherwise it pins a token that expires while the cache
// entry is still in use.
//
// This eager read happens synchronously within the caller's reconciliation, so
// it is bound to the supplied context whenever the retriever supports it. The
// retriever handed to the SDK, in contrast, deliberately holds a context
// detached from the reconciliation, see xpWebIdentityTokenRetriever.
func fingerprintIdentityToken(ctx context.Context, r stscreds.IdentityTokenRetriever) (string, error) {
	if cr, ok := r.(contextBoundTokenRetriever); ok {
		r = cr.withContext(ctx)
	}
	token, err := r.GetIdentityToken()
	if err != nil {
		return "", errors.Wrap(err, "cannot read the web identity token")
	}
	return fingerprintMaterial(fpDomainWebIdentityToken, token), nil
}

// contextBoundTokenRetriever is implemented by the
// stscreds.IdentityTokenRetriever implementations whose token reads observe a
// context, so that a caller reading the token synchronously can bind the read
// to its own context. The SDK's own IdentityTokenRetriever interface has no
// notion of a context.
type contextBoundTokenRetriever interface {
	withContext(ctx context.Context) stscreds.IdentityTokenRetriever
}

// fingerprintIRSACreds returns a non-reversible digest of the material behind
// the IRSA credentials, i.e. the projected service account token together with
// the environment that selects it. Unlike the WebIdentity source, IRSA takes
// both the token path and the role ARN from the environment instead of from the
// ProviderConfig spec, so neither is covered by the spec's generation and both
// belong in the digest.
func fingerprintIRSACreds() (string, error) {
	tokenFile := os.Getenv(envWebIdentityTokenFile)
	if tokenFile == "" {
		return "", errors.Errorf("environment variable %s must be set for the IRSA credential source", envWebIdentityTokenFile)
	}
	token, err := stscreds.IdentityTokenFile(filepath.Clean(tokenFile)).GetIdentityToken()
	if err != nil {
		return "", errors.Wrap(err, "cannot read the IRSA web identity token")
	}
	return fingerprintMaterial(fpDomainIRSA, token, []byte(tokenFile), []byte(os.Getenv(envWebIdentityRoleARN))), nil
}
