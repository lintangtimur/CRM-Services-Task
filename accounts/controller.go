package accounts

import (
	"CRM-Services-Task/utils"
	"errors"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	uc *UseCase
}

func (c Controller) Login(lr *LoginRequest) (*LoginResponse, error) {
	actor := Actor{}
	err := c.uc.Login(&actor, lr)

	if err != nil {
		return nil, err
	}
	err = utils.VerifyPassword(actor.Password, lr.Password)
	if err != nil {
		return nil, errors.New("password tidak sama")
	}

	token, err := utils.GenerateJwt(lr.Username, actor.RoleID)
	if err != nil {
		return nil, err
	}

	res := &LoginResponse{
		User: ActorResponse{
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

func (c Controller) CreateAdmin(a *AdminRegisterRequest) (*AdminRegisterResponse, error) {
	hashedPassword, err := utils.HashPassword(a.Password)
	if err != nil {
		return nil, err
	}

	actor := Actor{
		Username: a.Username,
		Password: hashedPassword,
		RoleID:   2,
	}

	err = c.uc.CreateAdmin(&actor)
	if err != nil {
		return nil, err
	}
	res := &AdminRegisterResponse{
		Status: "sukses",
		Data: ActorResponse{
			ID:       actor.ID,
			Username: actor.Username,
			RoleID:   actor.RoleID,
			Verified: actor.Verified,
			Active:   actor.Active,
		},
	}

	return res, nil
}

func (c Controller) GetApproveList() (*ApprovalListResponse, error) {
	acts, err := c.uc.GetApprovalList()
	if err != nil {
		return nil, err
	}
	res := &ApprovalListResponse{}
	for _, act := range acts {
		res.AdminID = append(res.AdminID, act.ID)
	}
	return res, nil
}

func (c Controller) ApproveAdmin(a *Actor, ar *ApproveRequest, ra *RegisterApproval) (*ApproveResponse, error) {
	valueToUpdate := map[string]interface{}{"verified": "true"}
	valueRa := map[string]interface{}{"status": "approved"}

	err := c.uc.ApproveAdmin(a, ar, ra, valueToUpdate, valueRa)

	if err != nil {
		return nil, err
	}
	res := ApproveResponse{
		Message: "admin berhasil di approve",
	}

	return &res, nil
}

func (c Controller) RejectAdmin(a *Actor, ar *ApproveRequest, ra *RegisterApproval) (*RejectResponse, error) {
	valueToUpdate := map[string]interface{}{"verified": "false", "active": "false"}
	valueRa := map[string]interface{}{"status": "rejected"}
	err := c.uc.RejectAdmin(a, ar, ra, valueToUpdate, valueRa)

	if err != nil {
		return nil, err
	}

	res := RejectResponse{
		Message: "admin berhasil di reject",
	}
	return &res, nil
}

func (c Controller) ActivateAdmin(a *ActivateAdminRequest) (*ActivateAdminResponse, error) {
	actor := Actor{}
	activeTrue := map[string]interface{}{"active": "true"}
	err := c.uc.ActivateAdmin(&actor, a, activeTrue)
	if err != nil {
		return nil, err
	}
	res := &ActivateAdminResponse{
		Message: "admin berhasil di aktifkan",
	}
	return res, nil
}

func (c Controller) DeactivateAdmin(d *DeActivateAdminRequest) (*DeActivateAdminResponse, error) {
	actor := Actor{}
	activeTrue := map[string]interface{}{"active": "false"}
	err := c.uc.DeactivateAdmin(&actor, d, activeTrue)
	if err != nil {
		return nil, err
	}
	res := &DeActivateAdminResponse{
		Message: "admin berhasil di non-aktifkan",
	}
	return res, nil
}

func (c Controller) DeleteAdmin(d *DeleteAdminRequest) (*DeleteAdminResponse, error) {
	admin := Actor{}

	//cek dulu apakah dia superadmin atau admin
	//karna admin tidak bisa menghapus admin yang lain atau dirinya sndiri atau superadmin
	err := c.uc.DeleteAdmin(&admin, d)
	if err != nil {
		return nil, err
	}
	res := &DeleteAdminResponse{
		Message: "data admin berhasil dihapus",
	}
	return res, nil
}

func (c Controller) GetAllAdmin(context *gin.Context) (*DataActorResponse, error) {
	username := context.Query("username")

	limit := context.DefaultQuery("limit", "5")
	page := context.DefaultQuery("page", "1")

	actor, err := c.uc.GetActorsByUsername(username, limit, page)
	if err != nil {
		return nil, err
	}

	res := &DataActorResponse{}
	for _, act := range actor {
		item := ActorResponse{
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
