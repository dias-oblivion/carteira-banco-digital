package database

import types "github.com/dias-oblivion/PicPay-Simplificado/api/types/request"

type Store interface {
	CreateUser(userName types.CreateUserRequest) (int, error)
}
