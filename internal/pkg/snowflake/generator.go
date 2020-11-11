package snowflake

import (
	"sync"

	"github.com/puti-projects/puti/internal/pkg/logger"

	"github.com/bwmarrin/snowflake"
)

var once sync.Once

var snowflakeNode *snowflake.Node

// getSnowflakeNodeInstance get singleton instance of snowflake node
func getSnowflakeNodeInstance() *snowflake.Node {
	if snowflakeNode == nil {
		once.Do(func() {
			var err error
			snowflakeNode, err = snowflake.NewNode(1) // just set 1 for now; TODO
			if err != nil {
				logger.Errorf("get snowflake id failed. %s", err)
				return
			}
		})
	}

	return snowflakeNode
}

// GenerateSnowflakeID generate a snowflake id
func GenerateSnowflakeID() int64 {
	sn := getSnowflakeNodeInstance()
	return sn.Generate().Int64()
}
