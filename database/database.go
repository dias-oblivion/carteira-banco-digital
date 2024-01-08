package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func NewPostgresStorage() error {
	db, err := sql.Open("postgres", "postgres://admin:admin@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}
	DB = db
	return nil
}

// func (s *PostgresStorage) GetTransactionHistoryListByUserId() (transactions []response.TransactionHistory, err error) {
// 	query := `SELECT u.name, u.document, th.description, th.value, 'RECEIVED'
// 	FROM transactions_history th
// 	INNER JOIN users u ON (th.fk_payee_id=u.id)
// 		UNION
// 	SELECT u.name, u.document, th.description, th.value, 'SENT'
// 	FROM transactions_history th
// 	INNER JOIN users u ON (th.fk_payer_id=u.id)
// 	`
// 	rows, err := s.db.Query(query)

// 	for rows.Next() {
// 		transaction := new(response.TransactionHistory)
// 		err := rows.Scan(
// 			&transaction.UserName,
// 			&
// 		)
// 	}
// }
