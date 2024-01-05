package services

import (
	request "github.com/dias-oblivion/PicPay-Simplificado/api/types/request"
	"github.com/dias-oblivion/PicPay-Simplificado/api/utils"
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

func (User) Login(credentials request.Login) (token string, err error) {
	user, err := repositories.User{}.GetUserByEmail(credentials.Email)
	if err != nil {
		return
	}

	if err = utils.CheckPassword(credentials.Password, user.Password); err != nil {
		return
	}

	token, err = utils.CreateJWTToken(user.ID, user.Email)

	if err != nil {
		return
	}

	return
}
