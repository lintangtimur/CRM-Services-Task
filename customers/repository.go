package customers

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func (r Repository) CreateCustomer(c *Customer) error {
	return r.db.Create(c).Error
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}
