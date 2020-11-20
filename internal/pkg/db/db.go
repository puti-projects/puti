package db

import (
	"fmt"
	"time"

	"github.com/puti-projects/puti/internal/pkg/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Engine *gorm.DB
)

// InitDB init db connection pool
func InitDB() error {
	var err error
	Engine, err = openDB(config.Db.Username, config.Db.Password, config.Db.Addr, config.Db.Name)
	if err != nil {
		return err
	}

	// set up config of db connection pool
	err = setupDB(Engine)
	if err != nil {
		return err
	}

	return nil
}

// openDB creates the DB connection
// It sets the location to UTC time
func openDB(username, password, addr, name string) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		"UTC",
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// setupDB sets the DB settings
func setupDB(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	// 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(config.Db.MaxIdleConns)
	// 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(config.Db.MaxOpenConns)
	// 设置连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)

	return nil
}
