package repositories

import (
	"errors"

	request "github.com/dias-oblivion/PicPay-Simplificado/api/types/request"
	"github.com/dias-oblivion/PicPay-Simplificado/database"
)

type Transfer struct{}

func (Transfer) TransferBalance(userID int64, transfer request.Transfer) error {
	sqlTransaction, err := database.DB.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			sqlTransaction.Rollback()
			panic(p)
		} else if err != nil {
			sqlTransaction.Rollback()
		} else {
			err = sqlTransaction.Commit()
		}
	}()

	var payerBalance, payeeBalance float64

	if err = sqlTransaction.QueryRow("SELECT balance from users where id = $1", userID).Scan(&payerBalance); err != nil {
		return err
	}
	if err = sqlTransaction.QueryRow("SELECT balance from users where id = $1", transfer.PayeeID).Scan(&payeeBalance); err != nil {
		return err
	}

	if payerBalance < transfer.Value {
		return errors.New("insufficient balance")
	}

	payerBalance -= transfer.Value
	payeeBalance += transfer.Value

	if _, err = sqlTransaction.Exec("UPDATE users SET balance = $1 where id = $2", payerBalance, userID); err != nil {
		return err
	}
	if _, err = sqlTransaction.Exec("UPDATE users SET balance = $1 where id = $2", payeeBalance, transfer.PayeeID); err != nil {
		return err
	}

	if _, err = sqlTransaction.Exec("INSERT INTO transfers_history(fk_payer_id, fk_payee_id, description, value) VALUES ($1, $2, $3, $4)",
		userID, transfer.PayeeID, transfer.Description, transfer.Value); err != nil {
		return err
	}

	return nil
}
