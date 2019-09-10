package dao

import (
	"vermouth/model"
)

type UserDao interface {
	List(parameter *model.Parameter) (*model.UserList, error)
}

type UserDaoImpl struct {
	Base
}

func (u UserDaoImpl) List(parameter *model.Parameter) (*model.UserList, error) {
	count := 0
	users := make([]model.User, 0)

	db := u.DB()
	db = db.Scopes(paginateScope(parameter))
	db.Model(&users).Count(&count)
	db.Find(&users)

	userList := &model.UserList{
		Data:     users,
		Page:     parameter.Page,
		PageSize: parameter.PageSize,
		Count:    count,
	}
	return userList, nil
}
