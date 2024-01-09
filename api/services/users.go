package services

import (
	"errors"

	request "github.com/dias-oblivion/carteira-banco-digital/api/types/request"
	"github.com/dias-oblivion/carteira-banco-digital/api/utils"
	"github.com/dias-oblivion/carteira-banco-digital/database/repositories"
	"golang.org/x/crypto/bcrypt"
)

type User struct{}

func (User) CreateUser(user request.CreateUser) (userId int, err error) {

	userRegistered, err := repositories.User{}.GetUserByEmail(user.Email)

	if err != nil {
		return 0, err
	}

	if userRegistered.Email == user.Email {
		return 0, errors.New("error: email already registered")
	}

	userRegistered, err = repositories.User{}.GetUserByDocument(user.Document)

	if userRegistered.Document == user.Document {
		return 0, errors.New("error: document number already registered")
	}

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

	if user.ID == 0 {
		return "", errors.New("error: user not found")
	}

	if err = utils.CheckPassword(credentials.Password, user.Password); err != nil {
		return
	}

	token, err = utils.CreateJWTToken(user.ID, user.Email, user.Role)

	if err != nil {
		return
	}

	return
}
