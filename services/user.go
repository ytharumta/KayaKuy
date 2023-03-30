package services

import (
	"KayaKuy/helper"
	"KayaKuy/models"
	"KayaKuy/repository"
	"fmt"
	"github.com/danilopolani/gocialite/structs"
	"os"
)

type UserService interface {
	GetOrRegisterUser(provider string, user *structs.User) models.User
	Register(user models.User) error
	Login(inputUser models.User, user *models.User) error
	UpdateUser(user models.User) (models.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *userService {
	return &userService{userRepository}
}

func (b *userService) GetOrRegisterUser(provider string, user *structs.User) models.User {
	var userData models.User
	err := b.userRepository.GetUserFromId(provider, user.ID, &userData)
	if err != nil {
		panic(err)
	}

	if userData.ID == 0 {
		newUser := models.User{
			FullName: user.FullName,
			Email:    user.Email,
			UserName: user.Username,
			SocialID: user.ID,
			Provider: provider,
			Avatar:   user.Avatar,
		}
		fmt.Println(newUser)
		err = b.userRepository.InsertUser(newUser)
		if err != nil {
			panic(err)
		}

		return newUser
	} else {
		return userData
	}
}

func (b *userService) Register(inputUser models.User) error {
	var user models.User
	var err error

	user.FullName = inputUser.FullName
	user.UserName = inputUser.UserName
	user.Email = inputUser.Email
	user.Password, err = helper.Encrypt(inputUser.Password, os.Getenv("SECRET"))
	if err != nil {
		return err
	}

	err = b.userRepository.RegisterUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (b *userService) Login(inputUser models.User, user *models.User) error {

	email := inputUser.Email
	password, err := helper.Decrypt(inputUser.Password, os.Getenv("SECRET"))
	if err != nil {
		return err
	}

	err = b.userRepository.CheckUser(email, password, user)
	if err != nil {
		return err
	}

	return nil
}

func (b *userService) UpdateUser(inputUser models.User) (models.User, error) {
	var err error
	inputUser.OldPassword, err = helper.Encrypt(inputUser.OldPassword, os.Getenv("SECRET"))
	inputUser.Password, err = helper.Encrypt(inputUser.Password, os.Getenv("SECRET"))
	user, err := b.userRepository.UpdateUser(inputUser)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
