package tickers

import (
	"time"

	"github.com/puti-projects/puti/internal/common/model"
	"github.com/puti-projects/puti/internal/common/utils"
	"github.com/puti-projects/puti/internal/pkg/logger"

	"github.com/jinzhu/gorm"
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
				if counterCache, found := utils.CounterCache.GetCounterCache(); found {
					for postID, number := range counterCache {
						err := model.DB.Local.Model(&model.PostModel{}).Where("`id` = ?", postID).Update("view_count", gorm.Expr("view_count + ?", number)).Error
						if err != nil {
							logger.Errorf("ticker: post count falied to update into database. %s", err)
						}
						utils.CounterCache.DeleteCounterCache(utils.GetPostCounterIPPoolKey(postID))
					}
					utils.CounterCache.DeleteCounterCache(utils.PostCounterKey)
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
