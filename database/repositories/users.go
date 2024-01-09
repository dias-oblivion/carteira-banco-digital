package repositories

import (
	request "github.com/dias-oblivion/carteira-banco-digital/api/types/request"
	"github.com/dias-oblivion/carteira-banco-digital/database"
	"github.com/dias-oblivion/carteira-banco-digital/database/models"
)

type User struct{}

func (User) CreateUser(user request.CreateUser) (userId int, err error) {
	query := `INSERT INTO users
		(name, email, document, password, role, balance, created_at)
	VALUES
		($1, $2, $3, $4, $5, 0, Now())
	RETURNING id`
	err = database.DB.QueryRow(query, user.Name, user.Email, user.Document, user.Password, user.Role).Scan(&userId)
	return
}

func (User) GetUserByEmail(email string) (user models.User, err error) {
	query := `SELECT
				    id,
					name,
					document,
					email,
					password,
					role,
					balance,
					created_at
			FROM 
			    	users
			WHERE
					email = $1
			LIMIT 1`
	data, err := database.DB.Query(query, email)

	if err != nil {
		return
	}

	for data.Next() {
		data.Scan(
			&user.ID,
			&user.Name,
			&user.Document,
			&user.Email,
			&user.Password,
			&user.Role,
			&user.Balance,
			&user.Created_at,
		)
	}

	return
}

func (User) GetUserByDocument(document string) (user models.User, err error) {
	query := `SELECT
					id,
					name,
					document,
					email,
					password,
					role,
					balance,
					created_at
				FROM 
					users
				WHERE
					document = $1
				LIMIT 1`

	data, err := database.DB.Query(query, document)

	if err != nil {
		return
	}

	for data.Next() {
		data.Scan(
			&user.ID,
			&user.Name,
			&user.Document,
			&user.Email,
			&user.Password,
			&user.Role,
			&user.Balance,
			&user.Created_at,
		)
	}

	return
}
