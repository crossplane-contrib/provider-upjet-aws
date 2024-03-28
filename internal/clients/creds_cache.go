// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"context"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/pkg/errors"

	"github.com/upbound/provider-aws/apis/v1beta1"
)

const (
	errGetAccountID = "cannot retrieve the AWS account ID"
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
		cache:   map[string]*awsCredentialsProviderCacheEntry{},
		maxSize: 100,
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
	// and ResourceVersion and additional fields depending on the auth method
	// (currently only IRSA temporary credential caching is supported).
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

// RetrieveCredentials returns a Credentials either from the credential cache.
// If the authentication scheme is IRSA and the supplied
// aws.CredentialsProvider implementation is an aws.CredentialsCache, then the
// retrieved credentials and the account ID are cached for future requests.
// Otherwise, this function returns the AWS credentials by calling
// the downstream aws.CredentialsProvider.Retrieve, and for now, does *not*
// call the given AccountIDFn because in that case, a separate identity cache
// should be used to retrieve the caller identity.
func (c *AWSCredentialsProviderCache) RetrieveCredentials(ctx context.Context, pc *v1beta1.ProviderConfig, region string, credsProvider aws.CredentialsProvider, accountIDFn AccountIDFn) (Credentials, error) {
	// Only IRSA credentials are cached currently and
	// only aws.CredentialsCache is supported as the underlying
	// credential provider.
	awsCredsCache, ok := credsProvider.(*aws.CredentialsCache)
	if !ok {
		c.logger.Debug("Configured aws.CredentialsProvider is not an aws.CredentialsCache, cannot utilize the provider credential cache...")
	}
	if pc.Spec.Credentials.Source != authKeyIRSA || !ok {
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
	// generate unique cache keys through UUID and ResourceVersion of
	// the ProviderConfig's k8s object, as they change when the provider
	// config is modified.
	//
	// Any other external parameter that have an effect on the resulting
	// credentials and does not appear in the ProviderConfig directly
	// (i.e. the same provider config content produces a different config),
	// should be included in the cache key.
	cacheKeyParams := []string{
		string(pc.UID),
		pc.ResourceVersion,
		region,
		string(pc.Spec.Credentials.Source),
	}
	tokenHash, err := hashTokenFile(os.Getenv("AWS_WEB_IDENTITY_TOKEN_FILE"))
	if err != nil {
		return Credentials{}, errors.Wrap(err, "cannot calculate the hash for the credentials file")
	}
	cacheKeyParams = append(cacheKeyParams, tokenHash, os.Getenv("AWS_WEB_IDENTITY_TOKEN_FILE"), os.Getenv("AWS_ROLE_ARN"))
	cacheKey := strings.Join(cacheKeyParams, ":")
	c.logger.Debug("Checking cache entry", "cacheKey", cacheKey, "pc", pc.GroupVersionKind().String())
	c.mu.RLock()
	cacheEntry, ok := c.cache[cacheKey]
	c.mu.RUnlock()

	// TODO: consider implementing a TTL even though the cached entry is valid
	// cache hit
	if ok {
		c.logger.Debug("Cache hit", "cacheKey", cacheKey, "pc", pc.GroupVersionKind().String())
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
		c.logger.Debug("Cache miss", "cacheKey", cacheKey, "pc", pc.GroupVersionKind().String(), "cacheSize", len(c.cache))
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
