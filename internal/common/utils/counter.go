package utils

import (
	"fmt"
	"time"

	gocache "github.com/patrickmn/go-cache"
)

// Counter cache utils
// It looks forward to create a cache buffer before calculate the post views
// default to set a expiration time of 10 minutes, in this time one IP will be count as one view for one article or page

const (
	// PostCounterKey key of post counter cache, store all post's count number as a map
	PostCounterKey = "PUTI_POST_VIEWS_CACHE"
	// PostCounterIPPoolKeyPrefix key prefix, IP Pool was be designed as a single key/value cache
	PostCounterIPPoolKeyPrefix = "PUTI_POST_VIEWS_CACHE_IP_"
)

// CounterCacheExpiration default expiration time
// The cache will store for 15 minutes, but these is a ticker runing will clean the data for 10 minutes expiration
var CounterCacheExpiration = 15 * time.Minute

// CounterCachePurgesExpiration default purges expired items
// Set for clean the cache, actually cache will be clean after ticker consume
var CounterCachePurgesExpiration = 10 * time.Minute

// CounterCache instance of counter cache
var CounterCache = &counterCache{
	cacheBody: gocache.New(CounterCacheExpiration, CounterCachePurgesExpiration),
}

// counterCache struct of counter cache
type counterCache struct {
	cacheBody *gocache.Cache
}

// GetPostCounterIPPoolKey get the key of IP Pool cache
func GetPostCounterIPPoolKey(postID uint64) string {
	postCounterIPPoolKey := fmt.Sprintf("%s%v", PostCounterIPPoolKeyPrefix, postID)
	return postCounterIPPoolKey
}

// CountOne count the number for the post by post id
func (cache *counterCache) CountOne(IP string, postID uint64) {
	postCounterIPPoolKey := GetPostCounterIPPoolKey(postID)
	counterIPPoolCache := make(map[string]bool)
	if x, found := cache.cacheBody.Get(postCounterIPPoolKey); found {
		counterIPPoolCachePointer := x.(*map[string]bool)
		counterIPPoolCache = *counterIPPoolCachePointer
		if _, ok := counterIPPoolCache[IP]; ok {
			return
		}
	}
	counterIPPoolCache[IP] = true
	cache.cacheBody.Set(postCounterIPPoolKey, &counterIPPoolCache, CounterCacheExpiration)

	// count ++
	counterCache := make(map[uint64]int)
	if postCounter, found := cache.cacheBody.Get(PostCounterKey); found {
		counterCachePointer := postCounter.(*map[uint64]int)
		counterCache = *counterCachePointer
		number := 1
		if number, ok := counterCache[postID]; ok {
			number++
		}
		counterCache[postID] = number
	} else {
		counterCache[postID] = 1
	}
	cache.cacheBody.Set(PostCounterKey, &counterCache, CounterCacheExpiration)
}

// GetCounterCache get counter cache data
func (cache *counterCache) GetCounterCache() (map[uint64]int, bool) {
	var counterCache map[uint64]int
	var found bool
	if x, found := cache.cacheBody.Get(PostCounterKey); found {
		counterCachePointer := x.(*map[uint64]int)
		counterCache = *counterCachePointer

		return counterCache, found
	}
	return counterCache, found
}

// DeleteCounterCache delete the counter cache after database action
func (cache *counterCache) DeleteCounterCache(key string) {
	cache.cacheBody.Delete(key)
	return
}
