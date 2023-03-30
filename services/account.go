package services

import (
	"KayaKuy/models"
	"KayaKuy/repository"
)

type AccountService interface {
	GetAllAccount(UserId int64) ([]models.Account, error)
	InsertAccount(account models.Account) error
	UpdateAccount(inputAccount models.Account, id int64) (int64, error)
	DeleteAccount(inputAccount models.Account, id int64) (int64, error)
}

type accountService struct {
	accountRepository repository.AccountRepository
}

func NewAccountService(accountRepository repository.AccountRepository) *accountService {
	return &accountService{accountRepository}
}

func (a *accountService) GetAllAccount(UserId int64) ([]models.Account, error) {
	err, account := a.accountRepository.GetAllAccount(UserId)
	if err != nil {
		return account, err
	}

	return account, nil
}

func (a *accountService) InsertAccount(inputAccount models.Account) error {
	var account models.Account

	account.UserID = inputAccount.UserID
	account.Name = inputAccount.Name

	err := a.accountRepository.InsertAccount(account)
	if err != nil {
		return err
	}

	return nil
}

func (a *accountService) UpdateAccount(inputAccount models.Account, id int64) (int64, error) {
	var account models.Account

	account.ID = id
	account.Name = inputAccount.Name
	account.UserID = inputAccount.UserID

	newAccount, err := a.accountRepository.UpdateAccount(account)
	if err != nil {
		return newAccount, err
	}
	return newAccount, nil
}

func (a *accountService) DeleteAccount(inputAccount models.Account, id int64) (int64, error) {
	var account models.Account

	account.ID = id
	account.Name = inputAccount.Name
	account.UserID = inputAccount.UserID

	newAccount, err := a.accountRepository.DeleteAccount(account)
	if err != nil {
		return newAccount, err
	}
	return newAccount, nil
}
