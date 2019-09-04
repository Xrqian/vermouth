package dao

import (
	"vermouth/model"
)

type UserDao interface {
	List(parameter *model.Parameter) (*model.UserList, error)
}
