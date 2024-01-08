package repositories

import (
	"errors"

	request "github.com/dias-oblivion/carteira-banco-digital/api/types/request"
	"github.com/dias-oblivion/carteira-banco-digital/database"
	"github.com/dias-oblivion/carteira-banco-digital/database/models"
)

type Transfer struct{}

func (Transfer) TransferBalance(userID int64, transfer request.TransferRequest) error {
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

func (Transfer) GetTransfersHistory(userID int64) (transfersHistory []models.Transfer, err error) {
	query := `SELECT th.id, u.name, u.document, th.description, th.value, 'RECEIVED' as transfer_type
	FROM transfers_history th
	INNER JOIN users u ON (th.fk_payee_id=u.id)
	where u.id = $1
	
	UNION
	
	SELECT th.id, u.name, u.document, th.description, th.value, 'SENT' as transfer_type
	FROM transfers_history th
	INNER JOIN users u ON (th.fk_payer_id=u.id)
	where u.id = $1`

	data, err := database.DB.Query(query, userID)

	if err != nil {
		return []models.Transfer{}, err
	}

	for data.Next() {
		var t models.Transfer
		data.Scan(
			&t.ID,
			&t.UserName,
			&t.UserDocument,
			&t.Description,
			&t.Value,
			&t.Type,
		)
		transfersHistory = append(transfersHistory, t)
	}

	return
}
