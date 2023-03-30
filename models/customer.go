package models

type Customer struct {
	ID       int64  `json:"id""`
	Name     string `json:"name""`
	IsVendor int64  `json:"is_vendor"`
	UserID   int64  `json:"user_id"`
}
