package counter

// Counter cache utils
// It looks forward to create a cache buffer before calculate the post views
// default to set a expiration time of 10 minutes, in this time one IP will be count as one view for one article or page
// Note: For simple implementation, the calculation may be repeated in a short time, because this count does not need to be accurate to a certain extent.

import (
	"strconv"
	"strings"

	"github.com/puti-projects/puti/internal/pkg/cache"
	"github.com/puti-projects/puti/internal/pkg/config"
	"github.com/puti-projects/puti/internal/pkg/logger"
	"github.com/puti-projects/puti/internal/utils"
)

// counterCache struct of counter cache
type counterCache struct {
	cacheBody *cache.Cache
}

// CounterCache instance of counter cache
var CounterCache = &counterCache{
	cacheBody: cache.GetInstance(),
}

// GetPostCounterIPPoolKey get the key of IP Pool cache
func GetPostCounterIPPoolKey(postID string) string {
	var builder strings.Builder
	builder.WriteString(config.CachePostCounterIPPoolKeyPrefix)
	builder.WriteString(postID)
	postCounterIPPoolKey := builder.String()
	return postCounterIPPoolKey
}

// CountOne count the number for the post by post id
func (c *counterCache) CountOne(IP string, postID uint64) {
	postIDStr := strconv.Itoa(int(postID))
	postCounterIPPoolKey := GetPostCounterIPPoolKey(postIDStr)

	counterIPPoolCache := make(map[string]interface{})
	iPPoolJSON, found := c.cacheBody.GetCacheWithByte(postCounterIPPoolKey)
	if found {
		if err := utils.JSON2Map(iPPoolJSON, &counterIPPoolCache); err != nil {
			logger.Errorf("an error occurred when converting JSON to map. %s", err)
		}
		if _, ok := counterIPPoolCache[IP]; ok {
			return
		}
	}

	counterIPPoolCache[IP] = true
	counterIPPoolCacheByte, err := utils.Map2JSON(counterIPPoolCache)
	if err != nil {
		logger.Errorf("an error occurred when converting map to JSON. %s", err)
	}
	if err := c.cacheBody.SetCacheWithByte(postCounterIPPoolKey, counterIPPoolCacheByte); err != nil {
		logger.Errorf("an error occurred when setting cache. %s", err)
	}

	// count ++
	counterCache := make(map[string]interface{})
	if postCounter, found := c.cacheBody.GetCacheWithByte(config.CachePostCounterKey); found {
		if err := utils.JSON2Map(postCounter, &counterCache); err != nil {
			logger.Errorf("an error occurred when converting JSON to map. %s", err)
		}

		if number, ok := counterCache[postIDStr]; ok {
			counterCache[postIDStr] = number.(int64) + 1
		} else {
			counterCache[postIDStr] = 1
		}
	} else {
		counterCache[postIDStr] = 1
	}

	counterCacheByte, err := utils.Map2JSON(counterCache)
	if err != nil {
		logger.Errorf("an error occurred when converting map to JSON. %s", err)
	}
	if err := c.cacheBody.SetCacheWithByte(config.CachePostCounterKey, counterCacheByte); err != nil {
		logger.Errorf("an error occurred when setting cache. %s", err)
	}
}

// GetCounterCache get counter cache data
func (c *counterCache) GetCounterCache() (map[string]interface{}, bool) {
	counterCache := make(map[string]interface{})
	if x, found := c.cacheBody.GetCacheWithByte(config.CachePostCounterKey); found {
		if err := utils.JSON2Map(x, &counterCache); err != nil {
			logger.Errorf("an error occurred when converting JSON to map. %s", err)
		}
		return counterCache, true
	}
	return counterCache, false
}

// DeleteCounterCache delete the counter cache after database action
func (c *counterCache) DeleteCounterCache(key string) error {
	return c.cacheBody.DeleteCache(key)
}
