package repository

import (
	"KayaKuy/models"
	"database/sql"
	"fmt"
)

type JournalRepository interface {
	GetAllJournal(UserId int64) (error, []models.Journal_entry_select)
	InsertJournal(journal models.Journal_entry) error
	UpdateJournal(journal models.Journal_entry) (int64, error)
	DeleteJournal(journal models.Journal_entry) (int64, error)
}

type journalRepository struct {
	db *sql.DB
}

func NewJournalRepo(db *sql.DB) *journalRepository {
	return &journalRepository{db}
}

func (a *journalRepository) GetAllJournal(UserID int64) (err error, results []models.Journal_entry_select) {
	sql := `SELECT je.id, je.code, c.name, a.name, je.value, je.note, u.user_name, je.transaction_type, je.created_at, je.updated_at FROM journal_entries je
         join customers c on c.id = je.customer_id
         join accounts a on a.id = je.account_id
         join users u on u.id = je.user_id
         where je.user_id = $1`

	rows, err := a.db.Query(sql, UserID)
	if err != nil {
		return err, []models.Journal_entry_select{}
	}

	for rows.Next() {
		var journal = models.Journal_entry_select{}
		err = rows.Scan(&journal.ID, &journal.Code, &journal.CustomerId, &journal.AccountId, &journal.Value, &journal.Note, &journal.UserID, &journal.TransactionType, &journal.CreatedAt, &journal.UpdatedAt)
		if err != nil {
			return err, []models.Journal_entry_select{}
		}
		results = append(results, journal)
	}
	return
}

func (a *journalRepository) InsertJournal(journal models.Journal_entry) (err error) {
	fmt.Println(journal.UserID)
	sql := "INSERT INTO journal_entries (code, customer_id, account_id, value, note, user_id, transaction_type, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, now())"

	errs := a.db.QueryRow(sql, journal.Code, journal.CustomerId, journal.AccountId, journal.Value, journal.Note, journal.UserID, journal.TransactionType)

	return errs.Err()
}

func (a *journalRepository) UpdateJournal(journal models.Journal_entry) (ct int64, err error) {
	sql := "UPDATE journal_entries set customer_id = $1, account_id = $2, value = $3, note = $4, user_id = $5, transaction_type = $6, updated_at = now() WHERE id = $7 and user_id = $8"

	res, errs := a.db.Exec(sql, journal.CustomerId, journal.AccountId, journal.Value, journal.Note, journal.UserID, journal.TransactionType, journal.ID, journal.UserID)
	if errs != nil {
		return 0, errs
	}
	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, err
}

func (a *journalRepository) DeleteJournal(customer models.Journal_entry) (ct int64, err error) {
	sql := "DELETE FROM journal_entries WHERE id = $1 and user_id = $2"

	res, errs := a.db.Exec(sql, customer.ID, customer.UserID)

	if errs != nil {
		return 0, errs
	}
	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, err
}
