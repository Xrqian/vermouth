package model

import (
	"github.com/jinzhu/gorm"
)

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
