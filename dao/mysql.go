package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
	"time"
	"vermouth/conf"
	"vermouth/model"
)

var (
	once sync.Once
	DB   *gorm.DB
)

func Init() {
	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.Cfg.MysqlUserName,
		conf.Cfg.MysqlDbPass,
		conf.Cfg.MysqlDbHost,
		conf.Cfg.MysqlDbPort,
		conf.Cfg.MysqlDbName,
	)
	once.Do(func() {
		db, err := gorm.Open("mysql", dbSource)
		db.SingularTable(true)
		db.DB().SetMaxOpenConns(100)
		if err != nil {
			fmt.Println(err.Error(), " 连接数据库失败")
		}
		db.DB().SetMaxOpenConns(conf.Cfg.MysqlMaxOpenConns)
		db.DB().SetConnMaxLifetime(10 * time.Minute)
		db.AutoMigrate(&model.User{}, &model.PayRecord{})
		// init DB
		DB = db
	})
}
