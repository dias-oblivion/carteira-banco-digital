package services

import (
	request "github.com/dias-oblivion/PicPay-Simplificado/api/types/request"
	"github.com/dias-oblivion/PicPay-Simplificado/database"
)

type User struct{}

func CreateUser(user request.CreateUser) (userId int, err error) {
	query := `INSERT INTO users
		(name, email, document, password, role, balance, created_at)
	VALUES
		($1, $2, $3, $4, $5, 0, Now())
	RETURNING id`
	err = database.DB.QueryRow(query, user.Name, user.Email, user.Document, user.Password, user.Role).Scan(&userId)
	return
}
