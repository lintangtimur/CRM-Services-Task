package accounts

import (
	"CRM-Services-Task/accounts/dto"
	"CRM-Services-Task/accounts/entity"
	"CRM-Services-Task/utils"
	"errors"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	uc    IUseCase
	auth  utils.IAuth
	token utils.IToken
}

func (c Controller) Login(lr *dto.LoginRequest) (*dto.LoginResponse, error) {
	actor := entity.Actor{}
	err := c.uc.Login(&actor, lr)

	if err != nil {
		return nil, err
	}
	err = c.auth.VerifyPassword(actor.Password, lr.Password)
	if err != nil {
		return nil, errors.New("password tidak sama")
	}
	token, err := c.token.GenerateJwt(lr.Username, actor.RoleID)
	if err != nil {
		return nil, err
	}

	res := &dto.LoginResponse{
		User: dto.ActorResponse{
			ID:       actor.ID,
			Username: actor.Username,
			RoleID:   actor.RoleID,
			Verified: actor.Verified,
			Active:   actor.Active,
		},
		Token: token,
	}
	return res, nil
}

func (c Controller) CreateAdmin(a *dto.AdminRegisterRequest) (*dto.AdminRegisterResponse, error) {

	hashedPassword, err := c.auth.HashPassword(a.Password)
	if err != nil {
		return nil, err
	}

	actor := entity.Actor{
		Username: a.Username,
		Password: hashedPassword,
		RoleID:   2,
	}

	err = c.uc.CreateAdmin(&actor)
	if err != nil {
		return nil, err
	}
	res := &dto.AdminRegisterResponse{
		Status: "sukses",
		Data: dto.ActorResponse{
			ID:       actor.ID,
			Username: actor.Username,
			RoleID:   actor.RoleID,
			Verified: actor.Verified,
			Active:   actor.Active,
		},
	}

	return res, nil
}

func (c Controller) GetApproveList() (*dto.ApprovalListResponse, error) {
	acts, err := c.uc.GetApprovalList()
	if err != nil {
		return nil, err
	}
	res := &dto.ApprovalListResponse{}
	for _, act := range acts {
		res.AdminID = append(res.AdminID, act.ID)
	}
	return res, nil
}

func (c Controller) ApproveAdmin(a *entity.Actor, ar *dto.ApproveRequest, ra *entity.RegisterApproval) (*dto.ApproveResponse, error) {
	valueToUpdate := map[string]interface{}{"verified": "true"}
	valueRa := map[string]interface{}{"status": "approved"}

	err := c.uc.ApproveAdmin(a, ar, ra, valueToUpdate, valueRa)

	if err != nil {
		return nil, err
	}
	res := dto.ApproveResponse{
		Message: "admin berhasil di approve",
	}

	return &res, nil
}

func (c Controller) RejectAdmin(a *entity.Actor, ar *dto.ApproveRequest, ra *entity.RegisterApproval) (*dto.RejectResponse, error) {
	valueToUpdate := map[string]interface{}{"verified": "false", "active": "false"}
	valueRa := map[string]interface{}{"status": "rejected"}
	err := c.uc.RejectAdmin(a, ar, ra, valueToUpdate, valueRa)

	if err != nil {
		return nil, err
	}

	res := dto.RejectResponse{
		Message: "admin berhasil di reject",
	}
	return &res, nil
}

func (c Controller) ActivateAdmin(a *dto.ActivateAdminRequest) (*dto.ActivateAdminResponse, error) {
	actor := entity.Actor{}
	activeTrue := map[string]interface{}{"active": "true"}
	err := c.uc.ActivateAdmin(&actor, a, activeTrue)
	if err != nil {
		return nil, err
	}
	res := &dto.ActivateAdminResponse{
		Message: "admin berhasil di aktifkan",
	}
	return res, nil
}

func (c Controller) DeactivateAdmin(d *dto.DeActivateAdminRequest) (*dto.DeActivateAdminResponse, error) {
	actor := entity.Actor{}
	activeTrue := map[string]interface{}{"active": "false"}
	err := c.uc.DeactivateAdmin(&actor, d, activeTrue)
	if err != nil {
		return nil, err
	}
	res := &dto.DeActivateAdminResponse{
		Message: "admin berhasil di non-aktifkan",
	}
	return res, nil
}

func (c Controller) DeleteAdmin(d *dto.DeleteAdminRequest) (*dto.DeleteAdminResponse, error) {
	admin := entity.Actor{}

	//cek dulu apakah dia superadmin atau admin
	//karna admin tidak bisa menghapus admin yang lain atau dirinya sndiri atau superadmin
	err := c.uc.DeleteAdmin(&admin, d)
	if err != nil {
		return nil, err
	}
	res := &dto.DeleteAdminResponse{
		Message: "data admin berhasil dihapus",
	}
	return res, nil
}

func (c Controller) GetAllAdmin(context *gin.Context) (*dto.DataActorResponse, error) {
	username := context.Query("username")

	limit := context.DefaultQuery("limit", "5")
	page := context.DefaultQuery("page", "1")

	actor, err := c.uc.GetActorsByUsername(username, limit, page)
	if err != nil {
		return nil, err
	}

	res := &dto.DataActorResponse{}
	for _, act := range actor {
		item := dto.ActorResponse{
			ID:       act.ID,
			Username: act.Username,
			Verified: act.Verified,
			Active:   act.Active,
			RoleID:   act.RoleID,
		}

		res.Data = append(res.Data, item)
	}

	return res, nil
}

func NewController(uc *UseCase) *Controller {
	return &Controller{
		uc: uc,
	}
}
