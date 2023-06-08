package customers

import (
	"CRM-Services-Task/customers/dto"
	"CRM-Services-Task/customers/entity"
)

type ICustomer interface {
	CreateCustomer(c *entity.Customer) error
	DeleteCustomer(c *entity.Customer, d *dto.DeleteCustomerRequest) error
	GetCustomersByNameAndEmail(firstname string, email string, limit string, page string) ([]entity.Customer, error)
}

type UseCase struct {
	repo IRepo
}

func (u UseCase) CreateCustomer(c *entity.Customer) error {
	return u.repo.CreateCustomer(c)
}

func (u UseCase) DeleteCustomer(c *entity.Customer, d *dto.DeleteCustomerRequest) error {
	return u.repo.DeleteCustomer(c, d)
}

func (u UseCase) GetCustomersByNameAndEmail(firstname string, email string, limit string, page string) ([]entity.Customer, error) {
	res, err := u.repo.FindAllCustomers(firstname, email, limit, page)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func NewUseCase(repo *Repository) *UseCase {
	return &UseCase{repo: repo}
}
