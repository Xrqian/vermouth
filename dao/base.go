package dao

import (
	"github.com/jinzhu/gorm"
	"vermouth/model"
)

type Base struct {
	db *gorm.DB
}

func (b Base) DB() *gorm.DB {
	if b.db != nil {
		return b.db
	}
	return DB
}

func paginateScope(parameter *model.Parameter) (scope func(*gorm.DB) *gorm.DB) {
	page, pageSize := parameter.Page, parameter.PageSize

	if pageSize == -1 {
		return func(db *gorm.DB) *gorm.DB { return db.Limit(200) }
	}

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	offset, limit := (page-1)*pageSize, pageSize

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(limit)
	}
}
