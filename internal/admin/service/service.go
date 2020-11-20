package service

import (
	"github.com/puti-projects/puti/internal/pkg/cache"
)

// Engine service layer's service
type Engine struct {
	Cache *cache.Cache
}

// SrvEngine service engine instance
var SrvEngine = &Engine{
	Cache: cache.GetInstance(),
}

// DeleteCache delete cache by key during service engine
func (s *Engine) DeleteCache(key string) error {
	return s.Cache.DeleteCache(key)
}
