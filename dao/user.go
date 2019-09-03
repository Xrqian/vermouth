package dao

import (
	"github.com/jinzhu/gorm"
	"vermouth/model"
)

type UserDao interface {
	List(parameter *model.Parameter) (*model.UserList, error)
}

type UserDapImpl struct {
	Mysql *gorm.DB
}

func (p UserDapImpl) List(parameter *model.Parameter) (*model.UserList, error) {
	count := 0
	var users = make([]model.User, 0)

	p.Mysql.Find(&users)

	results := &model.UserList{
		Page:     parameter.Page,
		PageSize: parameter.PageSize,
		Count:    count,
		Data:     users,
	}

	return results, nil
}
