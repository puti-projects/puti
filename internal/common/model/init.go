package model

import (
	"fmt"

	"github.com/puti-projects/puti/internal/pkg/logger"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql for gorm
	"github.com/spf13/viper"
)

// Database database struct
type Database struct {
	Local *gorm.DB
}

// DB definds the database
var DB *Database

// Init is the databases init function
func (db *Database) Init() {
	DB = &Database{
		Local: GetLocalDB(),
	}
}

// GetLocalDB gets the main DB
func GetLocalDB() *gorm.DB {
	return InitLocalDB()
}

// InitLocalDB inits the main DB
func InitLocalDB() *gorm.DB {
	return openDB(
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"),
	)
}

// openDB creates the DB connection
// It sets the location to UTC time
func openDB(username, password, addr, name string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4,utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		"UTC")

	db, err := gorm.Open("mysql", config)
	if err != nil {
		logger.Errorf("sql.Open() error(%v)", err)
		logger.Errorf("Database connection failed. Database name: %s", name)
	}

	// set for db connection
	setupDB(db)

	return db
}

// setupDB sets the DB settings
func setupDB(db *gorm.DB) {
	db.LogMode(viper.GetBool("gormlog"))
	db.DB().SetMaxOpenConns(150) // 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	db.DB().SetMaxIdleConns(20)  // 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
}

// Close close DB
func (db *Database) Close() {
	db.Close()
}
