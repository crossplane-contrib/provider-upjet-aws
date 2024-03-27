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
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/pkg/errors"

	"github.com/upbound/provider-aws/apis/v1beta1"
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
func WithCacheStore(cache map[string]awsCredentialsProviderCacheEntry) AWSCredentialsProviderCacheOption {
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
		cache:   map[string]awsCredentialsProviderCacheEntry{},
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
	cache map[string]awsCredentialsProviderCacheEntry

	// maxSize is the maximum number of elements this cache can ever have.
	maxSize int

	// mu is used to make sure the cache map is concurrency-safe.
	mu *sync.RWMutex

	// logger is the logger for cache operations.
	logger logging.Logger
}

type awsCredentialsProviderCacheEntry struct {
	awsCredCache *aws.CredentialsCache
	AccessedAt   time.Time
}

func (c *AWSCredentialsProviderCache) RetrieveCredentials(ctx context.Context, pc *v1beta1.ProviderConfig, region string, awsCredCache *aws.CredentialsCache) (aws.Credentials, error) {
	// Only IRSA authentication method credentials are cached currently
	if pc.Spec.Credentials.Source != authKeyIRSA {
		// skip cache for other/unimplemented credential types
		return awsCredCache.Retrieve(ctx)
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
		return aws.Credentials{}, errors.Wrap(err, "cannot calculate the hash for the credentials file")
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
		if time.Since(cacheEntry.AccessedAt) > 10*time.Minute {
			c.mu.Lock()
			cacheEntry.AccessedAt = time.Now()
			c.cache[cacheKey] = cacheEntry
			c.mu.Unlock()
		}
		return cacheEntry.awsCredCache.Retrieve(ctx)
	}

	// cache miss
	c.logger.Debug("Cache miss", "cacheKey", cacheKey, "pc", pc.GroupVersionKind().String())
	c.mu.Lock()
	c.makeRoom()
	c.cache[cacheKey] = awsCredentialsProviderCacheEntry{
		awsCredCache: awsCredCache,
		AccessedAt:   time.Now(),
	}
	c.mu.Unlock()
	return awsCredCache.Retrieve(ctx)
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
		if val.AccessedAt.Before(c.cache[dustiest].AccessedAt) {
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
	defer file.Close() //nolint:errcheck

	hash := sha256.New()
	if _, err = io.Copy(hash, file); err != nil {
		return "", err
	}

	checksum := hash.Sum(nil)
	return fmt.Sprintf("%x", checksum), nil
}
