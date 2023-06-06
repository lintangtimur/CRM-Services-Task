package accounts

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func (r Repository) Login(a *Actor, lr *LoginRequest) error {
	return r.db.Where("username = ? and verified = 'true' and active = 'true'", lr.Username).First(&a).Error
}

func (r Repository) CreateAdmin(a *Actor) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Create(&a).Error; err != nil {
			return err
		}
		ra := RegisterApproval{
			AdminID:      a.ID,
			SuperAdminID: 1,
			Status:       "",
		}

		if err := tx.Create(&ra).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (r Repository) FindAllApproval() ([]Actor, error) {
	var actor []Actor

	err := r.db.Preload("RA").Where("verified = 'false'").Find(&actor).Error
	if err != nil {
		return nil, err
	}
	return actor, err
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}
