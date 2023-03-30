package repository

import (
	"KayaKuy/models"
	"database/sql"
)

type CustomerRepository interface {
	GetAllCustomer(UserId int64) (error, []models.Customer)
	InsertCustomer(account models.Customer) error
	UpdateCustomer(account models.Customer) (int64, error)
	DeleteCustomer(account models.Customer) (int64, error)
}

type customerRepository struct {
	db *sql.DB
}

func NewCustomerRepo(db *sql.DB) *customerRepository {
	return &customerRepository{db}
}

func (a *customerRepository) GetAllCustomer(UserID int64) (err error, results []models.Customer) {
	sql := "SELECT * FROM customers where user_id = $1"

	rows, err := a.db.Query(sql, UserID)
	if err != nil {
		return err, []models.Customer{}
	}

	for rows.Next() {
		var customer = models.Customer{}
		err = rows.Scan(&customer.ID, &customer.Name, &customer.UserID, &customer.IsVendor)
		if err != nil {
			return err, []models.Customer{}
		}
		results = append(results, customer)
	}
	return
}

func (a *customerRepository) InsertCustomer(customer models.Customer) (err error) {
	sql := "INSERT INTO customers (name, user_id, is_vendor) VALUES ($1, $2, $3)"

	errs := a.db.QueryRow(sql, customer.Name, customer.UserID, customer.IsVendor)

	return errs.Err()
}

func (a *customerRepository) UpdateCustomer(customer models.Customer) (ct int64, err error) {
	sql := "UPDATE customers set name = $1, is_vendor = $2 WHERE id = $3 and user_id = $4"

	res, errs := a.db.Exec(sql, customer.Name, customer.IsVendor, customer.ID, customer.UserID)
	if errs != nil {
		return 0, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, err
}

func (a *customerRepository) DeleteCustomer(customer models.Customer) (ct int64, err error) {
	sql := "DELETE FROM customers WHERE id = $1 and user_id = $2"

	res, errs := a.db.Exec(sql, customer.ID, customer.UserID)

	if errs != nil {
		return 0, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, err
}
