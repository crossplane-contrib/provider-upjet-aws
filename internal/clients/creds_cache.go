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
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/crossplane/crossplane-runtime/v2/pkg/logging"
	"github.com/pkg/errors"

	"github.com/upbound/provider-aws/v2/apis/namespaced/v1beta1"
)

const (
	errGetAccountID = "cannot retrieve the AWS account ID"

	// defaultCacheMaxSize is the default maximum number of entries in the
	// AWS credentials provider cache.
	defaultCacheMaxSize = 1000

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
// your own cache.
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
		// entries are keyed by, among others, the region and the credential
		// material, so their count scales with the number of provider configs
		// times the number of regions in use, plus the rotated-out entries of
		// both. Evicting an entry means paying for its STS calls again, and the
		// entries are small, so keep the ceiling generous.
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
// Note that there is no need to invalidate the values in the cache because
// they never change, so we don't need concurrency-safety to prevent access
// to an invalidated entry.
type AWSCredentialsProviderCache struct {
	// cache holds the AWS Config with a unique cache key per
	// provider configuration. Key content includes the ProviderConfig's UUID
	// and Generation and additional fields depending on the auth method
	// (currently IRSA, and static credentials with an assume role chain, are
	// supported. See cacheableSource).
	cache map[string]*awsCredentialsProviderCacheEntry

	// maxSize is the maximum number of elements this cache can ever have.
	maxSize int

	// mu is used to make sure the cache map is concurrency-safe.
	mu *sync.RWMutex

	// logger is the logger for cache operations.
	logger logging.Logger
}

type awsCredentialsProviderCacheEntry struct {
	awsCredCache *aws.CredentialsCache
	accessedAt   atomic.Value
	accountID    atomic.Value
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
	case authKeyWebIdentity, authKeyPodIdentity, authKeyUpbound:
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
// If the authentication scheme is cacheable, i.e. IRSA or a static credential
// source with an assume role chain, and the supplied aws.CredentialsProvider
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
	// cache key calculation tries to capture any parameter that
	// could cause changes in the resulting AWS credentials,
	// to ensure unique keys.
	//
	// Parameters that are directly available in the provider config will
	// generate unique cache keys through UUID and Generation of
	// the ProviderConfig's k8s object, as they change when the provider
	// config is modified.
	//
	// Any other external parameter that have an effect on the resulting
	// credentials and does not appear in the ProviderConfig directly
	// (i.e. the same provider config content produces a different config),
	// should be included in the cache key via `cfgMeta.credsFingerprint`
	cacheKeyParams := []string{ // nolint:prealloc
		string(pc.UID),
		strconv.FormatInt(pc.Generation, 10),
		region,
		string(pc.Spec.Credentials.Source),
		cfgMeta.credsFingerprint, // empty fingerprints are rejected above
	}
	cacheKey := strings.Join(cacheKeyParams, ":")
	c.logger.Debug("Checking cache entry", "cacheKey", cacheKey, "providerConfigName", pc.Name, "providerConfigUID", string(pc.UID))
	c.mu.RLock()
	cacheEntry, ok := c.cache[cacheKey]
	c.mu.RUnlock()

	// TODO: consider implementing a TTL even though the cached entry is valid
	// cache hit
	if ok {
		c.logger.Debug("Cache hit", "cacheKey", cacheKey, "providerConfigName", pc.Name, "providerConfigUID", string(pc.UID))
		// since this is a hot-path in the execution, do not always update
		// the last access times, it is fine to evict the LRU entry on a less
		// granular precision.
		if time.Since(cacheEntry.accessedAt.Load().(time.Time)) > 10*time.Minute {
			cacheEntry.accessedAt.Store(time.Now())
		}
		return newCredentials(ctx, cacheEntry.awsCredCache, accountIDFromCacheEntry(cacheEntry))
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	// we need to recheck the cache because it might have already been
	// populated.
	cacheEntry, ok = c.cache[cacheKey]
	if !ok {
		// cache miss
		c.logger.Debug("Cache miss", "cacheKey", cacheKey, "providerConfigName", pc.Name, "providerConfigUID", string(pc.UID), "cacheSize", len(c.cache))
		c.makeRoom()
		cacheEntry = &awsCredentialsProviderCacheEntry{
			awsCredCache: awsCredsCache,
		}
		id, err := accountIDFn(ctx)
		if err != nil {
			return Credentials{}, errors.Wrap(err, errGetAccountID)
		}
		cacheEntry.accountID.Store(id)
		cacheEntry.accessedAt.Store(time.Now())
		c.cache[cacheKey] = cacheEntry
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

// hashTokenFile calculates the sha256 checksum of the token file content at
// the supplied file path
func hashTokenFile(filename string) (string, error) {
	if filename == "" {
		return "", errors.New("token file name cannot be empty")
	}
	file, err := os.Open(filepath.Clean(filename))
	if err != nil {
		return "", err
	}
	defer func() {
		_ = file.Close()
	}()

	hash := sha256.New()
	if _, err = io.Copy(hash, file); err != nil {
		return "", err
	}

	checksum := hash.Sum(nil)
	return fmt.Sprintf("%x", checksum), nil
}

func fingerprintIRSACreds() (string, error) {
	tokenHash, err := hashTokenFile(os.Getenv(envWebIdentityTokenFile))
	if err != nil {
		return "", errors.Wrap(err, "cannot calculate the hash for the credentials file")
	}
	return fmt.Sprintf("%s:%s:%s", tokenHash, os.Getenv(envWebIdentityTokenFile), os.Getenv(envWebIdentityRoleARN)), nil
}

// fingerprintKey keys the credential fingerprints. It is randomly generated
// per process, so that a fingerprint is meaningless outside of the process
// that produced it: fingerprints reach the credential cache keys, which are
// logged at debug level, and an unkeyed digest of a credential would be a
// stable identifier that anyone holding a candidate credential could confirm
// by recomputing it. The credential cache never outlives the process, so the
// fingerprints need no stability across restarts.
var fingerprintKey = []byte(rand.Text())

// fingerprintStaticCreds returns a non-reversible digest of the supplied
// aws.Credentials, suitable for use as a cache key component. The digest is
// only comparable against the digests produced by the same process, see
// fingerprintKey.
func fingerprintStaticCreds(creds aws.Credentials) string {
	h := hmac.New(sha256.New, fingerprintKey)
	for _, s := range []string{creds.AccessKeyID, creds.SecretAccessKey, creds.SessionToken} {
		// length-prefix framing, so that different credentials cannot digest
		// to the same value by concatenating to the same byte string. A mere
		// separator byte would leave the fields ambiguous if one of them ever
		// contained that byte itself.
		_ = binary.Write(h, binary.BigEndian, uint64(len(s)))
		_, _ = h.Write([]byte(s))
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}
