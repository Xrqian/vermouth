package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	OpenID     string      `gorm:"type:varchar(128);unique;not null",json:"open_id"`
	SessionKey string      `gorm:"type:varchar(128);not null",json:"session_key"`
	Nickname   string      `gorm:"type:varchar(128);not null",json:"nickname"`
	Avatar     string      `gorm:"type:varchar(255)",json:"avatar"`
	Province   string      `gorm:"type:varchar(255)",json:"province"`
	Country    string      `gorm:"type:varchar(255)",json:"country"`
	Gender     int         `gorm:"type:int(1)",json:"gender"`
	PayRecords []PayRecord `gorm:"ForeignKey:UserID;"`
}

func (User) TableName() string {
	return "user"
}

type PayRecord struct {
	gorm.Model
	Cny      float32 `gorm:"type:float;not null",json:"cny"`
	Mode     string  `gorm:"type:varchar(255)",json:"mode"`
	Category string  `gorm:"type:varchar(255)",json:"category"`
	Comment  string  `gorm:"type:varchar(255)",json:"comment"`
	UserID   int     `gorm:"type:int(10);not null",json:"user_id"`
}

func (PayRecord) TableName() string {
	return "pay_record"
}
