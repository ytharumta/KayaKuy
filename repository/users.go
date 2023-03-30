package repository

import (
	"KayaKuy/models"
	"database/sql"
	"errors"
)

type UserRepository interface {
	GetUserFromId(provide string, social_id string, user *models.User) error
	CheckUser(user_name string, password string, user *models.User) error
	InsertUser(user models.User) error
	RegisterUser(user models.User) error
	UpdateUser(user models.User) (models.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *userRepository {
	return &userRepository{db}
}

func (b *userRepository) GetUserFromId(provider string, social_id string, user *models.User) (err error) {
	sql := "SELECT id, full_name, user_name, email, avatar, social_id, provider, role FROM users where social_id = $1 and provider = $2"

	rows, err := b.db.Query(sql, social_id, provider)
	if err != nil {
		return err
	}

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.FullName, &user.UserName, &user.Email, &user.Avatar, &user.SocialID, &user.Provider, &user.Role)
		if err != nil {
			return err
		}
	}
	return
}

func (b *userRepository) CheckUser(user_name string, password string, user *models.User) (err error) {
	success := false
	sql := "SELECT id, full_name, user_name, email FROM users where (user_name = $1 or email = $1) and password = $2"

	rows, err := b.db.Query(sql, user_name, password)
	if err != nil {
		return err
	}

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.FullName, &user.UserName, &user.Email)
		if err != nil {
			return err
		}
		success = true
	}
	if !success {
		return errors.New("Login Gagal")
	}
	return
}

func (b *userRepository) InsertUser(user models.User) (err error) {
	sql := "INSERT INTO users (full_name, user_name, email, social_id, avatar, provider) VALUES ($1, $2, $3, $4, $5, $6)"

	errs := b.db.QueryRow(sql, user.FullName, user.UserName, user.Email, user.SocialID, user.Avatar, user.Provider)

	return errs.Err()
}

func (b *userRepository) RegisterUser(user models.User) (err error) {
	sql := "INSERT INTO users (full_name, user_name, email, password) VALUES ($1, $2, $3, $4)"

	errs := b.db.QueryRow(sql, user.FullName, user.UserName, user.Email, user.Password)

	return errs.Err()
}

func (a *userRepository) UpdateUser(inputUser models.User) (getUser models.User, err error) {
	var user models.User

	sql := "UPDATE users set full_name = $1, password=$2 WHERE id = $3 and (password=$4 or password is null) returning id, full_name, user_name, email"
	//fmt.Println(sql, inputUser.FullName, inputUser.Password, inputUser.ID, inputUser.OldPassword)
	errs := a.db.QueryRow(sql, inputUser.FullName, inputUser.Password, inputUser.ID, inputUser.OldPassword).Scan(&user.ID, &user.FullName, &user.UserName, &user.Email)
	if errs != nil {
		return models.User{}, errs
	}
	return user, err
}
