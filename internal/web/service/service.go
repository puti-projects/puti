package service

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/puti-projects/puti/internal/pkg/cache"
	"github.com/puti-projects/puti/internal/pkg/logger"
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

// JSONUnmarshal json unmarshal
func (s *Engine) JSONUnmarshal(data []byte, v interface{}) {
	if err := s.JSON.Unmarshal(data, v); err != nil {
		logger.Errorf("found cache, but the conversion failed.")
	}
}

// MarshalAndSetCache json marshal and set cache
func (s *Engine) MarshalAndSetCache(key string, v interface{}) {
	byteData, err := s.JSON.Marshal(v)
	if err != nil {
		logger.Errorf("json convert failed before set cache. %s", err)
	}
	if err := s.SetCache(key, byteData); err != nil {
		logger.Errorf("set cache failed. %s", err)
	}
}
