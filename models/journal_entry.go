package models

import "database/sql"

type Journal_entry struct {
	ID              int64        `json:"id""`
	Code            string       `json:"code"`
	CustomerId      int64        `json:"customer_id"`
	AccountId       int64        `json:"account_id"`
	Value           float64      `json:"value"`
	Note            string       `json:"note"`
	TransactionType string       `json:"transaction_type"`
	UserID          int64        `json:"user_id"`
	CreatedAt       sql.NullTime `json:"created_at"`
	UpdatedAt       sql.NullTime `json:"updated_at"`
}

type Journal_entry_select struct {
	ID              int64        `json:"id""`
	Code            string       `json:"code"`
	CustomerId      string       `json:"customer_id"`
	AccountId       string       `json:"account_id"`
	Value           float64      `json:"value"`
	Note            string       `json:"note"`
	TransactionType string       `json:"transaction_type"`
	UserID          string       `json:"user_id"`
	CreatedAt       sql.NullTime `json:"created_at"`
	UpdatedAt       sql.NullTime `json:"updated_at"`
}
