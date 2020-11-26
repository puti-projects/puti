package service

import (
	"github.com/puti-projects/puti/internal/pkg/cache"
	"github.com/puti-projects/puti/internal/pkg/config"
	"github.com/puti-projects/puti/internal/pkg/logger"
	"strconv"
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
func (s *Engine) DeleteCache(key string) {
	if err := s.Cache.DeleteCache(key); err != nil {
		logger.Errorf("error deleting cache. %s", err)
	}
}

// CleanCacheAfterEditArticle clean cache after update article
func (s *Engine) CleanCacheAfterEditArticle(articleID uint64) {
	// article detail cache
	s.DeleteCache(config.CacheArticleDetailPrefix + strconv.Itoa(int(articleID)))
}

// CleanCacheAfterEditPage clean cache after update page
func (s *Engine) CleanCacheAfterEditPage(pageID uint64) {
	// page detail cache
	s.DeleteCache(config.CachePageDetailPrefix + strconv.Itoa(int(pageID)))
}

// CleanCacheAfterUpdateKnowledge clean cache after update knowledge info
func (s *Engine) CleanCacheAfterUpdateKnowledge(slug string) {
	// knowledge info cache
	s.DeleteCache(config.CacheKnowledgeInfoPrefix + slug)
}

// CleanCacheKnowledgeItemList
func (s *Engine) CleanCacheKnowledgeItemList(kID uint64) {
	// knowledge item list cache
	s.DeleteCache(config.CacheKnowledgeItemListPrefix + strconv.Itoa(int(kID)))
}

// CleanCacheAfterUpdateKnowledgeItemContent clean cache after update knowledge item content
func (s *Engine) CleanCacheAfterUpdateKnowledgeItemContent(kiSymbol uint64) {
	// knowledge item content cache by item symbol
	s.DeleteCache(config.CacheKnowledgeItemContentPrefix + strconv.Itoa(int(kiSymbol)))
}
