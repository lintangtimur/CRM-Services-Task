package accounts

import (
	"CRM-Services-Task/accounts/dto"
	"CRM-Services-Task/accounts/entity"
	"gorm.io/gorm"
	"strconv"
)

type IRepo interface {
	Login(a *entity.Actor, lr *dto.LoginRequest) error
	CreateAdmin(a *entity.Actor) error
	FindAllApproval() ([]entity.Actor, error)
	UpdateActor(a *entity.Actor, ar *dto.ApproveRequest, update map[string]interface{}) error
	UpdateStatusRA(ra *entity.RegisterApproval, ar *dto.ApproveRequest, val map[string]interface{}) error
	ActivateAdmin(a *entity.Actor, aar *dto.ActivateAdminRequest, activeTrue map[string]interface{}) error
	DeactivateAdmin(a *entity.Actor, d *dto.DeActivateAdminRequest, val map[string]interface{}) error
	DeleteAdmin(a *entity.Actor, d *dto.DeleteAdminRequest) error
	FindAllActors(username string, limit string, page string) ([]entity.Actor, error)
}

type Repository struct {
	db *gorm.DB
}

func (r Repository) Login(a *entity.Actor, lr *dto.LoginRequest) error {
	return r.db.Where("username = ? and verified = 'true' and active = 'true'", lr.Username).First(&a).Error
}

func (r Repository) CreateAdmin(a *entity.Actor) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Create(&a).Error; err != nil {
			return err
		}
		ra := entity.RegisterApproval{
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

func (r Repository) FindAllApproval() ([]entity.Actor, error) {
	var actor []entity.Actor

	err := r.db.Preload("RA").Where("verified = 'false'").Find(&actor).Error
	if err != nil {
		return nil, err
	}
	return actor, err
}

func (r Repository) UpdateActor(a *entity.Actor, ar *dto.ApproveRequest, update map[string]interface{}) error {
	return r.db.Model(&a).Where("id = ?", ar.AdminID).First(&a).Updates(update).Error
}

func (r Repository) UpdateStatusRA(ra *entity.RegisterApproval, ar *dto.ApproveRequest, val map[string]interface{}) error {
	return r.db.Model(&ra).Where("admin_id = ?", ar.AdminID).Updates(val).Error
}

func (r Repository) ActivateAdmin(a *entity.Actor, aar *dto.ActivateAdminRequest, activeTrue map[string]interface{}) error {
	return r.db.Model(&a).Where("id = ?", aar.AdminID).First(&a).Updates(activeTrue).Error
}

func (r Repository) DeactivateAdmin(a *entity.Actor, d *dto.DeActivateAdminRequest, val map[string]interface{}) error {
	return r.db.Model(&a).Where("id = ?", d.AdminID).First(&a).Updates(val).Error
}

func (r Repository) DeleteAdmin(a *entity.Actor, d *dto.DeleteAdminRequest) error {
	if err := r.db.First(&a, d.AdminID).Error; err != nil {
		return err
	}
	return r.db.Delete(&a).Error
}

func (r Repository) FindAllActors(username string, limit string, page string) ([]entity.Actor, error) {
	var actor []entity.Actor
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
