package cache

import (
	"errors"
	"time"

	"github.com/puti-projects/puti/internal/pkg/logger"

	"github.com/allegro/bigcache"
)

type Cache struct {
	*bigcache.BigCache
}

// CacheInstance
var Instance *Cache

// LoadCache load cache
func LoadCache() error {
	config := bigcache.Config{
		// number of shards (must be a power of 2)
		Shards: 1024,

		// time after which entry can be evicted
		LifeWindow: 60 * time.Minute,

		// Interval between removing expired entries (clean up).
		// If set to <= 0 then no action is performed.
		// Setting to < 1 second is counterproductive — bigcache has a one second resolution.
		CleanWindow: 10 * time.Minute,

		// rps * lifeWindow, used only in initial memory allocation
		MaxEntriesInWindow: 1000 * 10 * 60,

		// max entry size in bytes, used only in initial memory allocation
		MaxEntrySize: 500,

		// prints information about additional memory allocation
		Verbose: true,

		// cache will not allocate more memory than this limit, value in MB
		// if value is reached then the oldest entries can be overridden for the new ones
		// 0 value means no size limit
		HardMaxCacheSize: 204800,

		// callback fired when the oldest entry is removed because of its expiration time or no space left
		// for the new entry, or because delete was called. A bitmask representing the reason will be returned.
		// Default value is nil which means no callback and it prevents from unwrapping the oldest entry.
		OnRemove: nil,

		// OnRemoveWithReason is a callback fired when the oldest entry is removed because of its expiration time or no space left
		// for the new entry, or because delete was called. A constant representing the reason will be passed through.
		// Default value is nil which means no callback and it prevents from unwrapping the oldest entry.
		// Ignored if OnRemove is specified.
		OnRemoveWithReason: nil,
	}

	c, err := bigcache.NewBigCache(config)
	if err != nil {
		return err
	}

	Instance = &Cache{
		c,
	}
	return nil
}

// GetInstance get instance
func GetInstance() *Cache {
	if Instance == nil {
		if err := LoadCache(); err != nil {
			logger.Errorf("get cache instance failed. init cache failed. %s", err)
		}
	}
	return Instance
}

// SetCacheWithByte set cache by key using byte data
func (c *Cache) SetCacheWithByte(key string, value []byte) error {
	return c.Set(key, value)
}

// SetCache set cache by key
func (c *Cache) SetCache(key, value string) error {
	return c.Set(key, []byte(value))
}

// GetCacheWithByte get cache by key. It returns the cache value in byte and ifFound
func (c *Cache) GetCacheWithByte(key string) ([]byte, bool) {
	entry, err := c.Get(key)
	if err != nil {
		if errors.Is(err, bigcache.ErrEntryNotFound) {
			return nil, false
		}

		logger.Errorf("error when getting cache. key: %s. err: %s", key, err)
		return nil, false
	}

	return entry, true
}

// GetCache get cache by key. It returns the cache value and ifFound
func (c *Cache) GetCache(key string) (string, bool) {
	entry, found := c.GetCacheWithByte(key)
	if found {
		return string(entry), found
	}
	return "", found
}

// DeleteCache delete cache by key
func (c *Cache) DeleteCache(key string) error {
	if err := c.Delete(key); err != nil && !errors.Is(err, bigcache.ErrEntryNotFound) {
		return err
	}
	return nil
}

// Flush delete all cache
func (c *Cache) Flush() error {
	return c.Reset()
}
