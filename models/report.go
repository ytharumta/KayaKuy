package models

import "time"

type Report struct {
	History        []History
	AccountBalance []AccountBalance
	Total          []Total
}

type History struct {
	Code            string    `json:"code"`
	Note            string    `json:"note"`
	ToFrom          string    `json:"to_from"`
	Account         string    `json:"account"`
	TransactionType string    `json:"transaction_type"`
	Value           float64   `json:"value"`
	CreatedAt       time.Time `json:"created_at"`
}

type AccountBalance struct {
	Account     string  `json:"account"`
	TotalDebit  float64 `json:"total_debit"`
	TotalKredit float64 `json:"total_kredit"`
	Total       float64 `json:"total"`
}

type Total struct {
	TotalDebit  float64 `json:"total_debit"`
	TotalKredit float64 `json:"total_kredit"`
	Precentage  float64 `json:"precentage"`
	Note        string  `json:"note"`
}
