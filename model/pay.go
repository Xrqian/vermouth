package model

import "time"

type User struct {
	Id         int64     `json:"id,omitempty",gorm:"primary_key;type:int(11);not null" `
	OpenID     string    `json:"open_id,omitempty",gorm:"type:varchar(255);unique;not null"`
	SessionKey string    `json:"session_key,omitempty",gorm:"type:varchar(255)"`
	Avatar     string    `json:"avatar,omitempty",gorm:"type:varchar(255)"`
	Nickname   string    `json:"nickname,omitempty",gorm:"type:varchar(255)"`
	Gender     int       `json:"gender,omitempty",gorm:"type:int(1)"`
	Province   string    `json:"province,omitempty",gorm:"type:varchar(255)"`
	Country    string    `json:"country,omitempty",gorm:"type:varchar(255)"`
	DateJoined time.Time `json:"date_joined,omitempty",gorm:"type:datetime;ASSOCIATION_AUTOUPDATE"`
	LastLogin  time.Time `json:"last_login,omitempty",gorm:"type:datetime;ASSOCIATION_AUTOUPDATE"`
}

type PayRecord struct {
}

func (User) TableName() string {
	return "user"
}

