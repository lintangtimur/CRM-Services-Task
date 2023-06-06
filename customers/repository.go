package customers

import (
	"gorm.io/gorm"
	"strconv"
)

type Repository struct {
	db *gorm.DB
}

func (r Repository) CreateCustomer(c *Customer) error {
	return r.db.Create(c).Error
}

func (r Repository) DeleteCustomer(c *Customer, d *DeleteCustomerRequest) error {
	if err := r.db.First(&c, d.CustomerID).Error; err != nil {
		return err
	}

	return r.db.Delete(&c).Error
}

func (r Repository) FindAllCustomers(firstname string, email string, limit string, page string) ([]Customer, error) {
	var customer []Customer
	limits, _ := strconv.Atoi(limit)
	pages, _ := strconv.Atoi(page)
	offset := (pages - 1) * limits

	err := r.db.Limit(limits).Offset(offset).Find(&customer).Error
	search := r.db
	if firstname != "" {
		search = search.Where("first_name = ?", firstname)
	}
	if email != "" {
		search = search.Where("email = ?", email)
	}
	err = search.Limit(limits).Offset(offset).Find(&customer).Error
	if err != nil {
		return nil, err
	}

	return customer, err
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}
