package database

import (
	"database/sql"

	types "github.com/dias-oblivion/PicPay-Simplificado/api/types/request"
	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage() (*PostgresStorage, error) {
	db, err := sql.Open("postgres", "postgres://postgres:picpay@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStorage{
		db: db,
	}, nil
}

func (s *PostgresStorage) CreateUser(user types.CreateUserRequest) (userId int, err error) {
	query := `INSERT INTO users
		(name, email, document, password, role, balance, created_at)
	VALUES
		($1, $2, $3, $4, $5, 0, Now())
	RETURNING id`
	err = s.db.QueryRow(query, user.Name, user.Email, user.Document, user.Password, user.Role).Scan(&userId)
	return
}
