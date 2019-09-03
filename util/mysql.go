package util

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
)

type MysqlConnect struct {
	DB *gorm.DB
}

var mysqlConnect *MysqlConnect

var once sync.Once

func GetConnect(username string, password string, host string, port string, db string) *MysqlConnect {
	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		db,
	)
	once.Do(func() {
		db, err := gorm.Open("mysql", dbSource)
		db.SingularTable(true)
		db.DB().SetMaxOpenConns(100)
		if err != nil {
			fmt.Println(err.Error(), " 连接数据库失败")
		}
		mysqlConnect = &MysqlConnect{db}
	})
	return mysqlConnect
}
