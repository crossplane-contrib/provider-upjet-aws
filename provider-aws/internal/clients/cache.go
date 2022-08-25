/*
Copyright 2022 Upbound Inc.
*/

package clients

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/pkg/errors"
)

// GlobalCallerIdentityCache is a global cache to be used by all controllers.
var GlobalCallerIdentityCache = NewCallerIdentityCache(100)

// NewCallerIdentityCache returns a new empty *CallerIdentityCache.
func NewCallerIdentityCache(maxSize int) *CallerIdentityCache {
	return &CallerIdentityCache{
		cache:   map[string]*sts.GetCallerIdentityOutput{},
		maxSize: maxSize,
	}
}

// CallerIdentityCache holds GetCallerIdentityOutput objects in memory so that
// we don't need to make API calls to AWS in every reconciliation of every
// resource. It has a maximum size that when it's reached, random entries from
// the cache will be deleted to keep a ceiling on the memory it uses.
type CallerIdentityCache struct {
	// cache holds caller identity with a key whose format is the following:
	// <access_key>:<secret_key>:<token>
	// Any of the variables could be empty.
	cache map[string]*sts.GetCallerIdentityOutput

	// maxSize is the maximum number of elements this cache can have. When it
	// reaches above this number, randomly selected entries will be deleted.
	maxSize int
}

// GetCallerIdentity returns the identity of the caller.
func (c *CallerIdentityCache) GetCallerIdentity(ctx context.Context, cfg aws.Config, creds aws.Credentials) (*sts.GetCallerIdentityOutput, error) {
	key := fmt.Sprintf("%s:%s:%s",
		creds.AccessKeyID,
		creds.SecretAccessKey,
		creds.SessionToken,
	)
	if i, ok := c.cache[key]; ok {
		return i, nil
	}
	i, err := sts.NewFromConfig(cfg).GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		return nil, errors.Wrap(err, "GetCallerIdentity query failed")
	}
	c.cache[key] = i
	c.checkCache()
	return i, nil
}

// NOTE(muvaf): I considered deleting the entries that are not accessed in
// a given timeframe or are older than a given expiration window. However, the
// only worry we have for this cache is to use too much memory because
// invalidation is not necessary since the value is always the same for a given
// caller. The maximum size method guarantees that we don't have ever-growing
// memory, which may be possible with other methods.

// checkCache checks whether we reached the maximum number of elements and
// remove randomly selected entries till we reduce it back to the maximum. Since
// the value for a given key never changes, we don't really need to invalidate
// entries. It's only to put a limit on how much memory it can ever use.
func (c *CallerIdentityCache) checkCache() {
	n := len(c.cache) - c.maxSize
	for key := range c.cache {
		if n <= 0 {
			return
		}
		delete(c.cache, key)
		n--
	}
}
