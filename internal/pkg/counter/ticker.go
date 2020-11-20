package counter

import (
	"github.com/puti-projects/puti/internal/pkg/config"
	"time"

	"github.com/puti-projects/puti/internal/model"
	"github.com/puti-projects/puti/internal/pkg/db"
	"github.com/puti-projects/puti/internal/pkg/logger"

	"gorm.io/gorm"
)

const (
	// RepeatTime ticker repeat time
	RepeatTime = time.Minute * 10
)

// CountTickerStopChan chan for stop the count ticker
var CountTickerStopChan = make(chan bool)

// InitCountTicker init the count ticker
func InitCountTicker() {
	countTicker := time.NewTicker(RepeatTime)
	countTickerChan := countTicker.C

	go func() {
		for {
			select {
			case <-countTickerChan:
				// TODO only deal with post view now; (maybe add knowledge base view count)
				if counterCache, found := CounterCache.GetCounterCache(); found {
					for postID, number := range counterCache {
						// TODO make it to one query
						err := db.Engine.Model(&model.Post{}).Where("`id` = ?", postID).Update("view_count", gorm.Expr("view_count + ?", number)).Error
						if err != nil {
							logger.Errorf("ticker: post count failed to update into database. %s", err)
						}
						// delete IP pool cache
						if err := CounterCache.DeleteCounterCache(GetPostCounterIPPoolKey(postID)); err != nil {
							logger.Errorf("deleted IP pool cache failed. %s", err)
						}
					}
					// delete post view number cache
					if err := CounterCache.DeleteCounterCache(config.CachePostCounterKey); err != nil {
						logger.Errorf("deleted counter cache failed. %s", err)
					}
				}
			case <-CountTickerStopChan:
				logger.Info("Ticker will be Stop")
				countTicker.Stop()
				logger.Info("Ticker stopped")
				return
			}
		}
	}()

	logger.Info("start to running the count ticker")
}

// StopCountTicker stop the count ticker
func StopCountTicker() {
	CountTickerStopChan <- true
	return
}
