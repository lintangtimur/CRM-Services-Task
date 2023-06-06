package accounts

import (
	"CRM-Services-Task/utils"
	"errors"
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

func NewController(uc *UseCase) *Controller {
	return &Controller{
		uc: uc,
	}
}
