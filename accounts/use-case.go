package accounts

import (
	"CRM-Services-Task/accounts/dto"
	"CRM-Services-Task/accounts/entity"
)

type IUseCase interface {
	Login(a *entity.Actor, lr *dto.LoginRequest) error
	CreateAdmin(a *entity.Actor) error
	GetApprovalList() ([]entity.Actor, error)
	ApproveAdmin(a *entity.Actor, ar *dto.ApproveRequest, ra *entity.RegisterApproval, update map[string]interface{}, ra2 map[string]interface{}) error
	RejectAdmin(a *entity.Actor, ar *dto.ApproveRequest, ra *entity.RegisterApproval, update map[string]interface{}, valueRa map[string]interface{}) error
	ActivateAdmin(a *entity.Actor, aar *dto.ActivateAdminRequest, activeTrue map[string]interface{}) error
	DeactivateAdmin(a *entity.Actor, d *dto.DeActivateAdminRequest, val map[string]interface{}) error
	DeleteAdmin(a *entity.Actor, d *dto.DeleteAdminRequest) error
	GetActorsByUsername(username string, limit string, page string) ([]entity.Actor, error)
}

type UseCase struct {
	repo IRepo
}

func (u UseCase) Login(a *entity.Actor, lr *dto.LoginRequest) error {
	return u.repo.Login(a, lr)
}

func (u UseCase) CreateAdmin(a *entity.Actor) error {
	return u.repo.CreateAdmin(a)
}

func (u UseCase) GetApprovalList() ([]entity.Actor, error) {
	res, err := u.repo.FindAllApproval()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u UseCase) ApproveAdmin(a *entity.Actor, ar *dto.ApproveRequest, ra *entity.RegisterApproval, update map[string]interface{}, ra2 map[string]interface{}) error {
	err := u.repo.UpdateActor(a, ar, update)
	if err != nil {
		return err
	}
	err = u.repo.UpdateStatusRA(ra, ar, ra2)
	if err != nil {
		return err
	}
	return err
}

func (u UseCase) RejectAdmin(a *entity.Actor, ar *dto.ApproveRequest, ra *entity.RegisterApproval, update map[string]interface{}, valueRa map[string]interface{}) error {
	err := u.repo.UpdateActor(a, ar, update)
	if err != nil {
		return err
	}
	err = u.repo.UpdateStatusRA(ra, ar, valueRa)
	if err != nil {
		return err
	}
	return err
}

func (u UseCase) ActivateAdmin(a *entity.Actor, aar *dto.ActivateAdminRequest, activeTrue map[string]interface{}) error {
	return u.repo.ActivateAdmin(a, aar, activeTrue)
}

func (u UseCase) DeactivateAdmin(a *entity.Actor, d *dto.DeActivateAdminRequest, val map[string]interface{}) error {
	return u.repo.DeactivateAdmin(a, d, val)
}

func (u UseCase) DeleteAdmin(a *entity.Actor, d *dto.DeleteAdminRequest) error {
	return u.repo.DeleteAdmin(a, d)
}

func (u UseCase) GetActorsByUsername(username string, limit string, page string) ([]entity.Actor, error) {
	return u.repo.FindAllActors(username, limit, page)
}

func NewUseCase(repo *Repository) *UseCase {
	return &UseCase{repo: repo}
}
