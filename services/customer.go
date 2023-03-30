package services

import (
	"KayaKuy/models"
	"KayaKuy/repository"
)

type CustomerService interface {
	GetAllCustomer(UserId int64) ([]models.Customer, error)
	InsertCustomer(account models.Customer) error
	UpdateCustomer(inputAccount models.Customer, id int64) (int64, error)
	DeleteCustomer(inputAccount models.Customer, id int64) (int64, error)
}

type customerService struct {
	customerRepository repository.CustomerRepository
}

func NewCustomerService(customerRepository repository.CustomerRepository) *customerService {
	return &customerService{customerRepository}
}

func (a *customerService) GetAllCustomer(UserId int64) ([]models.Customer, error) {
	err, customer := a.customerRepository.GetAllCustomer(UserId)
	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (a *customerService) InsertCustomer(inputCustomer models.Customer) error {
	var customer models.Customer

	customer.UserID = inputCustomer.UserID
	customer.Name = inputCustomer.Name
	customer.IsVendor = inputCustomer.IsVendor

	err := a.customerRepository.InsertCustomer(customer)
	if err != nil {
		return err
	}

	return nil
}

func (a *customerService) UpdateCustomer(inputCustomer models.Customer, id int64) (int64, error) {
	var customer models.Customer

	customer.ID = id
	customer.Name = inputCustomer.Name
	customer.IsVendor = inputCustomer.IsVendor
	customer.UserID = inputCustomer.UserID

	newCustomer, err := a.customerRepository.UpdateCustomer(customer)
	if err != nil {
		return newCustomer, err
	}
	return newCustomer, nil
}

func (a *customerService) DeleteCustomer(inputCustomer models.Customer, id int64) (int64, error) {
	var customer models.Customer

	customer.ID = id
	customer.Name = inputCustomer.Name
	customer.UserID = inputCustomer.UserID

	newCustomer, err := a.customerRepository.DeleteCustomer(customer)
	if err != nil {
		return newCustomer, err
	}
	return newCustomer, nil
}
