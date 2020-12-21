package service

import (
	"github.com/puti-projects/puti/internal/pkg/cache"
	"github.com/puti-projects/puti/internal/pkg/db"
	"github.com/puti-projects/puti/internal/pkg/logger"
	"github.com/puti-projects/puti/internal/web/dao"

	jsoniter "github.com/json-iterator/go"
)

// Engine service engine
type Engine struct {
	Cache *cache.Cache
	JSON  jsoniter.API
	dao   *dao.Dao
}

// SrvEngine service engine instance
var SrvEngine *Engine

// NewServiceEngine return a new Service engine for frontend web
// this method will only be called once at the beginning for initialization
func NewServiceEngine() error {
	if SrvEngine == nil {
		SrvEngine = &Engine{
			Cache: cache.GetInstance(),
			JSON:  jsoniter.ConfigCompatibleWithStandardLibrary,
			dao:   dao.New(db.Engine),
		}
	}
	return nil
}

// GetCache get cache by key during service engine
func (svc *Engine) GetCache(key string) ([]byte, bool) {
	return svc.Cache.GetCacheWithByte(key)
}

// SetCache set cache by key during service engine
func (svc *Engine) SetCache(key string, value []byte) error {
	return svc.Cache.SetCacheWithByte(key, value)
}

// JSONUnmarshal json unmarshal
func (svc *Engine) JSONUnmarshal(data []byte, v interface{}) {
	if err := svc.JSON.Unmarshal(data, v); err != nil {
		logger.Errorf("found cache, but the conversion failed.")
	}
}

// MarshalAndSetCache json marshal and set cache
func (svc *Engine) MarshalAndSetCache(key string, v interface{}) {
	byteData, err := svc.JSON.Marshal(v)
	if err != nil {
		logger.Errorf("json convert failed before set cache. %s", err)
	}
	if err := svc.SetCache(key, byteData); err != nil {
		logger.Errorf("set cache failed. %s", err)
	}
}
