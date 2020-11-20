package service

import (
	"github.com/puti-projects/puti/internal/pkg/cache"

	jsoniter "github.com/json-iterator/go"
)

// Engine service engine
type Engine struct {
	Cache *cache.Cache
	JSON  jsoniter.API
}

// SrvEngine service engine instance
var SrvEngine = &Engine{
	Cache: cache.GetInstance(),
	JSON:  jsoniter.ConfigCompatibleWithStandardLibrary,
}

// GetCache get cache by key during service engine
func (s *Engine) GetCache(key string) ([]byte, bool) {
	return s.Cache.GetCacheWithByte(key)
}

// GetCache set cache by key during service engine
func (s *Engine) SetCache(key string, value []byte) error {
	return s.Cache.SetCacheWithByte(key, value)
}
