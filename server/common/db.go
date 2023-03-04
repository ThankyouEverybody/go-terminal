package common

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"github.com/go-terminal-server/common/config"
	"github.com/go-terminal-server/model/orm"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var DB *gorm.DB

func init() {
	//初始化链接数据库
	fmt.Println("========setup database========")
	var logMod logger.Interface
	if config.GlobalConfig.Server.Log.Level == "debug" {
		logMod = logger.Default.LogMode(logger.Info)
	} else {
		logMod = logger.Default.LogMode(logger.Silent)
	}
	var err error
	var db *gorm.DB
	//if config.GlobalConfig.Server.DB == "sqlite" {
	if true {
		dsn := fmt.Sprintf("file:%s?mode=rwc&cache=shared", config.GlobalConfig.Sqlite.File)
		db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{
			Logger:                 logMod,
			SkipDefaultTransaction: true,
		})
	}
	if err != nil {
		panic(fmt.Errorf("连接数据库异常: %v", err.Error()).(any))
	}
	err = db.AutoMigrate(&orm.User{})

	if err != nil {
		panic(fmt.Errorf("初始化数据库表结构异常: %v", err.Error()).(any))
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = db
}
