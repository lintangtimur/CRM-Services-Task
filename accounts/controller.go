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

func NewController(uc *UseCase) *Controller {
	return &Controller{
		uc: uc,
	}
}
