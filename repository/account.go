package repository

import (
	"KayaKuy/models"
	"database/sql"
)

type AccountRepository interface {
	GetAllAccount(UserId int64) (error, []models.Account)
	InsertAccount(account models.Account) error
	UpdateAccount(account models.Account) (int64, error)
	DeleteAccount(account models.Account) (int64, error)
}

type accountRepository struct {
	db *sql.DB
}

func NewAccountRepo(db *sql.DB) *accountRepository {
	return &accountRepository{db}
}

func (a *accountRepository) GetAllAccount(UserID int64) (err error, results []models.Account) {
	sql := "SELECT * FROM accounts where user_id = $1"

	rows, err := a.db.Query(sql, UserID)
	if err != nil {
		return err, []models.Account{}
	}

	for rows.Next() {
		var account = models.Account{}
		err = rows.Scan(&account.ID, &account.Name, &account.UserID)
		if err != nil {
			return err, []models.Account{}
		}
		results = append(results, account)
	}
	return
}

func (a *accountRepository) InsertAccount(account models.Account) (err error) {
	sql := "INSERT INTO accounts (name, user_id) VALUES ($1, $2)"

	errs := a.db.QueryRow(sql, account.Name, account.UserID)

	return errs.Err()
}

func (a *accountRepository) UpdateAccount(account models.Account) (ct int64, err error) {
	sql := "UPDATE accounts set name = $1 WHERE id = $2 and user_id = $3"

	res, errs := a.db.Exec(sql, account.Name, account.ID, account.UserID)
	if errs != nil {
		return 0, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, err
}

func (a *accountRepository) DeleteAccount(account models.Account) (ct int64, err error) {
	sql := "DELETE FROM accounts WHERE id = $1 and user_id = $2"

	res, errs := a.db.Exec(sql, account.ID, account.UserID)

	if errs != nil {
		return 0, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, err
}
