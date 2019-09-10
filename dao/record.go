package dao

import (
	"vermouth/model"
)

type PayRecordDao interface {
	List(parameter *model.Parameter) (*model.PayRecordList, error)
}

type PayRecordDaoImpl struct {
	Base
}

func (u PayRecordDaoImpl) List(parameter *model.Parameter) (*model.PayRecordList, error) {
	count := 0
	records := make([]model.PayRecord, 0)

	db := u.DB()
	db = db.Scopes(paginateScope(parameter))
	db.Model(&records).Count(&count)
	db.Find(&records)

	recordList := &model.PayRecordList{
		Data:     records,
		Page:     parameter.Page,
		PageSize: parameter.PageSize,
		Count:    count,
	}
	return recordList, nil
}
