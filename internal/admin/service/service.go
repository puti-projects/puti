package service

import (
	"context"
	"strconv"

	"github.com/puti-projects/puti/internal/admin/dao"
	"github.com/puti-projects/puti/internal/pkg/cache"
	"github.com/puti-projects/puti/internal/pkg/config"
	"github.com/puti-projects/puti/internal/pkg/db"
	"github.com/puti-projects/puti/internal/pkg/logger"
)

// Service service layer's service
type Service struct {
	ctx   context.Context
	dao   *dao.Dao
	cache *cache.Cache
}

// New return a new Service with context
func New(ctx context.Context) Service {
	return Service{
		ctx:   ctx,
		dao:   dao.New(db.Engine),  // global db instance
		cache: cache.GetInstance(), // global cache instance
	}
}

// DeleteCache delete cache by key during service engine
func (svc *Service) DeleteCache(key string) {
	if err := svc.cache.DeleteCache(key); err != nil {
		logger.Errorf("error deleting cache. %s", err)
	}
}

// CleanCacheAfterEditArticle clean cache after update article
func (svc *Service) CleanCacheAfterEditArticle(articleID uint64) {
	// article detail cache
	svc.DeleteCache(config.CacheArticleDetailPrefix + strconv.Itoa(int(articleID)))
}

// CleanCacheAfterEditPage clean cache after update page
func (svc *Service) CleanCacheAfterEditPage(pageID uint64) {
	// page detail cache
	svc.DeleteCache(config.CachePageDetailPrefix + strconv.Itoa(int(pageID)))
}

// CleanCacheAfterUpdateKnowledge clean cache after update knowledge info
func (svc *Service) CleanCacheAfterUpdateKnowledge(slug string) {
	// knowledge info cache
	svc.DeleteCache(config.CacheKnowledgeInfoPrefix + slug)
}

// CleanCacheKnowledgeItemList clean cache of knowledge item list
func (svc *Service) CleanCacheKnowledgeItemList(kID uint64) {
	// knowledge item list cache
	svc.DeleteCache(config.CacheKnowledgeItemListPrefix + strconv.Itoa(int(kID)))
}

// CleanCacheAfterUpdateKnowledgeItemContent clean cache after update knowledge item content
func (svc *Service) CleanCacheAfterUpdateKnowledgeItemContent(kiSymbol uint64) {
	// knowledge item content cache by item symbol
	svc.DeleteCache(config.CacheKnowledgeItemContentPrefix + strconv.Itoa(int(kiSymbol)))
}
