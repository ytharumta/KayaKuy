package models

type Account struct {
	ID     int64  `json:"id""`
	Name   string `json:"name""`
	UserID int64  `json:"user_id"`
}
