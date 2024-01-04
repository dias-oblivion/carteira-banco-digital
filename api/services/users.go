package services

import (
	request "github.com/dias-oblivion/PicPay-Simplificado/api/types/request"
	"github.com/dias-oblivion/PicPay-Simplificado/database/repositories"
	"golang.org/x/crypto/bcrypt"
)

type User struct{}

func (User) CreateUser(user request.CreateUser) (userId int, err error) {
	password := []byte(user.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	user.Password = string(hashedPassword)
	userId, err = repositories.User{}.CreateUser(user)
	if err != nil {
		return 0, err
	}

	return
}
