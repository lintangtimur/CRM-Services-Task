package accounts

import (
	"gorm.io/gorm"
	"strconv"
)

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

func (r Repository) UpdateActor(a *Actor, ar *ApproveRequest, update map[string]interface{}) error {
	return r.db.Model(&a).Where("id = ?", ar.AdminID).First(&a).Updates(update).Error
}

func (r Repository) UpdateStatusRA(ra *RegisterApproval, ar *ApproveRequest, val map[string]interface{}) error {
	return r.db.Model(&ra).Where("admin_id = ?", ar.AdminID).Updates(val).Error
}

func (r Repository) ActivateAdmin(a *Actor, aar *ActivateAdminRequest, activeTrue map[string]interface{}) error {
	return r.db.Model(&a).Where("id = ?", aar.AdminID).First(&a).Updates(activeTrue).Error
}

func (r Repository) DeactivateAdmin(a *Actor, d *DeActivateAdminRequest, val map[string]interface{}) error {
	return r.db.Model(&a).Where("id = ?", d.AdminID).First(&a).Updates(val).Error
}

func (r Repository) DeleteAdmin(a *Actor, d *DeleteAdminRequest) error {
	if err := r.db.First(&a, d.AdminID).Error; err != nil {
		return err
	}
	return r.db.Delete(&a).Error
}

func (r Repository) FindAllActors(username string, limit string, page string) ([]Actor, error) {
	var actor []Actor
	limits, _ := strconv.Atoi(limit)
	pages, _ := strconv.Atoi(page)
	offset := (pages - 1) * limits

	err := r.db.Limit(limits).Offset(offset).Find(&actor).Error
	search := r.db
	if username != "" {
		search = search.Where("username = ?", username)
	}

	err = search.Limit(limits).Offset(offset).Find(&actor).Error
	if err != nil {
		return nil, err
	}

	return actor, err
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}
