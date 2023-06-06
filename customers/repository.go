package customers

import "gorm.io/gorm"

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

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}
