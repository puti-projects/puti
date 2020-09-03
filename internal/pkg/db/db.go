package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/puti-projects/puti/internal/pkg/config"
)

var (
	DBEngine *gorm.DB
)

func InitDB() error {
	var err error
	DBEngine, err = openDB(
		config.Db.Username,
		config.Db.Password,
		config.Db.Addr,
		config.Db.Name,
	)
	if err != nil {
		return err
	}

	// set for db connection
	setupDB(DBEngine)

	return nil
}

// openDB creates the DB connection
// It sets the location to UTC time
func openDB(username, password, addr, name string) (*gorm.DB, error) {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4,utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		"UTC")
	db, err := gorm.Open("mysql", config)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// setupDB sets the DB settings
func setupDB(db *gorm.DB) {
	// gorm log mode
	if config.Server.Runmode == "debug" {
		db.LogMode(true)
	} else {
		db.LogMode(false)
	}

	// connection pool setting
	db.DB().SetMaxOpenConns(config.Db.MaxOpenConns) // 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	db.DB().SetMaxIdleConns(config.Db.MaxIdleConns) // 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
}
