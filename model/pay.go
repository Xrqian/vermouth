package model

import "time"

type User struct {
	Id         int64     `json:"id",gorm:"primary_key;type:int(11);not null" `
	OpenID     string    `json:"open_id",gorm:"type:varchar(255);unique;not null"`
	SessionKey string    `json:"session_key",gorm:"type:varchar(255)"`
	Avatar     string    `json:"avatar",gorm:"type:varchar(255)"`
	Nickname   string    `json:"nickname",gorm:"type:varchar(255)"`
	Gender     int       `json:"gender",gorm:"type:int(1)"`
	Province   string    `json:"province",gorm:"type:varchar(255)"`
	Country    string    `json:"country",gorm:"type:varchar(255)"`
	DateJoined time.Time `json:"date_joined",gorm:"type:datetime;ASSOCIATION_AUTOUPDATE"`
	LastLogin  time.Time `json:"last_login",gorm:"type:datetime;ASSOCIATION_AUTOUPDATE"`
}

func (User) TableName() string {
	return "user"
}

type PayRecord struct {
	Id       int64   `json:"id",gorm:"primary_key;type:int(11);not null" `
	Cny      float32 `json:"cny",gorm:"type:float;not null" `
	Mode     string  `json:"mode",gorm:"type:varchar(255)" `
	Category string  `json:"category",gorm:"type:varchar(255)" `
	Comment  string  `json:"comment",gorm:"type:varchar(255)" `
	User     User    `gorm:"ForeignKey:PayRecordId"`
}

func (PayRecord) TableName() string {
	return "pay_record"
}
