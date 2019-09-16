package dao

import (
	"vermouth/model"
)

type PayRecordDao interface {
	List(parameter *model.Parameter) (*model.PayRecordList, error)
	Retrieve(id int) (*model.User, error)
}

type PayRecordDaoImpl struct {
	Base
}

func (u PayRecordDaoImpl) List(parameter *model.Parameter) (*model.PayRecordList, error) {
	count := 0
	records := make([]model.PayRecord, 0)

	db := u.DB()
	err := db.Model(&records).Count(&count).Error
	if err != nil {
		return nil, err
	}
	db = db.Scopes(paginateScope(parameter))
	db.Find(&records)
	recordList := &model.PayRecordList{
		Data:     records,
		Page:     parameter.Page,
		PageSize: parameter.PageSize,
		Count:    count,
	}
	return recordList, nil
}

func (u PayRecordDaoImpl) Retrieve(id int) (*model.PayRecord, error) {
	record := model.PayRecord{}

	db := u.DB()

	err := db.First(&record, id).Error
	if err != nil {
		return nil, err
	}
	return &record, nil
}
