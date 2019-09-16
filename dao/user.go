package dao

import (
	"vermouth/model"
)

type UserDao interface {
	List(parameter *model.Parameter) (*model.UserList, error)
	Retrieve(id int) (*model.User, error)
	Create(user model.User) (*model.User, error)
}

type UserDaoImpl struct {
	Base
}

func (u UserDaoImpl) List(parameter *model.Parameter) (*model.UserList, error) {
	count := 0
	users := make([]model.User, 0)

	db := u.DB()
	err := db.Model(&users).Count(&count).Error
	if err != nil {
		return nil, err
	}
	db = db.Scopes(paginateScope(parameter))
	db.Find(&users)
	userList := &model.UserList{
		Data:     users,
		Page:     parameter.Page,
		PageSize: parameter.PageSize,
		Count:    count,
	}
	return userList, nil
}

func (u UserDaoImpl) Retrieve(id int) (*model.User, error) {
	user := model.User{}

	db := u.DB()
	err := db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u UserDaoImpl) Create(user model.User) (*model.User, error) {
	db := u.DB()
	err := db.Create(&user).Error
	return &user, err
}
