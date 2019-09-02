package main

import (
	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"sync"
	"encoding/json"
)

type User struct {
	Id          int64 `json:"id,omitempty",gorm:"primary_key,type:int(11),not null" `
	Username    string
	FirstName   string
	LastName    string
	Email       string
	IsActive    bool
	IsStaff     bool
	IsSuperuser bool
	DateJoined  time.Time
	LastLogin   time.Time
}

func (User) TableName() string {
	return "auth_user"
}

type MysqlConnect struct {
	Db *gorm.DB
}

var instance *MysqlConnect

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
		instance = &MysqlConnect{db}
	})
	return instance
}

func Json(value interface{}) []byte {
	b, err := json.Marshal(value)
	if err != nil {
		fmt.Println("Umarshal failed:", err)
		return nil
	}
	return b
}

func main() {

}
