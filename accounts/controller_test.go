package accounts

import (
	"CRM-Services-Task/accounts/dto"
	"CRM-Services-Task/accounts/entity"
	mocks "CRM-Services-Task/mocks/accounts"
	mocksUtils "CRM-Services-Task/mocks/utils"
	"CRM-Services-Task/utils"
	"errors"
	"reflect"
	"testing"
)

func TestController_CreateAdmin(t *testing.T) {
	type fields struct {
		uc    IUseCase
		auth  utils.IAuth
		token utils.IToken
	}
	type args struct {
		a *dto.AdminRegisterRequest
	}
	mockUseCase := mocks.NewIUseCase(t)
	mockauth := mocksUtils.NewIAuth(t)

	req := dto.AdminRegisterRequest{
		Username: "lintang",
		Password: "",
	}
	reqbenar := dto.AdminRegisterRequest{
		Username: "lintang",
		Password: "qwerty",
	}

	actor := entity.Actor{
		Username: req.Username,
		Password: "$2a$10$oSmtc1CDNlZwnCl/BBy9auzaEDCPVbslhpkghimkhCaKvrHbHcrvW",
		RoleID:   2,
	}

	err := errors.New("hash password gagal")
	mockauth.EXPECT().HashPassword(req.Password).Return("", err).Once()

	mockauth.EXPECT().HashPassword("qwerty").Return("$2a$10$oSmtc1CDNlZwnCl/BBy9auzaEDCPVbslhpkghimkhCaKvrHbHcrvW", nil).Once()
	errdb := errors.New("gagal insert ke db")
	mockUseCase.EXPECT().CreateAdmin(&actor).Return(errdb).Once()

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
	mockauth.EXPECT().HashPassword("qwerty").Return("$2a$10$oSmtc1CDNlZwnCl/BBy9auzaEDCPVbslhpkghimkhCaKvrHbHcrvW", nil).Once()
	mockUseCase.EXPECT().CreateAdmin(&actor).Return(nil).Once()

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dto.AdminRegisterResponse
		wantErr bool
	}{
		{
			name: "error on hashpassword",
			fields: fields{
				uc:    mockUseCase,
				auth:  mockauth,
				token: nil,
			},
			args: args{
				a: &req,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "create admin gagal",
			fields: fields{
				uc:   mockUseCase,
				auth: mockauth,
			},
			args: args{
				a: &reqbenar,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "sukses insert",
			fields: fields{
				uc:   mockUseCase,
				auth: mockauth,
			},
			args: args{
				a: &reqbenar,
			},
			want:    res,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Controller{
				uc:    tt.fields.uc,
				auth:  tt.fields.auth,
				token: tt.fields.token,
			}
			got, err := c.CreateAdmin(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateAdmin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateAdmin() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestController_Login(t *testing.T) {
	type fields struct {
		uc    IUseCase
		auth  utils.IAuth
		token utils.IToken
	}
	type args struct {
		lr *dto.LoginRequest
	}
	mockUseCase := mocks.NewIUseCase(t)
	mockToken := mocksUtils.NewIToken(t)
	mockauth := mocksUtils.NewIAuth(t)

	req := dto.LoginRequest{
		Username: "lintang",
		Password: "qwerty",
	}
	actor := entity.Actor{}
	err := errors.New("gagal login")
	mockUseCase.EXPECT().Login(&actor, &req).Return(err).Once()

	err = errors.New("password tidak sama")
	mockUseCase.EXPECT().Login(&actor, &req).Return(nil).Once()
	mockauth.EXPECT().VerifyPassword(actor.Password, req.Password).Return(err).Once()

	mockUseCase.EXPECT().Login(&actor, &req).Return(nil).Once()
	mockauth.EXPECT().VerifyPassword(actor.Password, req.Password).Return(nil).Once()
	err = errors.New("error jwt")
	mockToken.EXPECT().GenerateJwt(req.Username, actor.RoleID).Return("", err).Once()

	mockUseCase.EXPECT().Login(&actor, &req).Return(nil).Once()
	mockauth.EXPECT().VerifyPassword(actor.Password, req.Password).Return(nil).Once()
	mockToken.EXPECT().GenerateJwt(req.Username, actor.RoleID).Return("123", nil).Once()

	res := &dto.LoginResponse{
		User: dto.ActorResponse{
			ID:       actor.ID,
			Username: actor.Username,
			RoleID:   actor.RoleID,
			Verified: actor.Verified,
			Active:   actor.Active,
		},
		Token: "123",
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dto.LoginResponse
		wantErr bool
	}{
		{
			name: "login gagal",
			fields: fields{
				uc:    mockUseCase,
				auth:  mockauth,
				token: mockToken,
			},
			args: args{
				lr: &req,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "password tidak sama",
			fields: fields{
				uc:    mockUseCase,
				auth:  mockauth,
				token: mockToken,
			},
			args: args{
				lr: &req,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "generate jwt error",
			fields: fields{
				uc:    mockUseCase,
				auth:  mockauth,
				token: mockToken,
			},
			args: args{
				lr: &req,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "login sukses",
			fields: fields{
				uc:    mockUseCase,
				auth:  mockauth,
				token: mockToken,
			},
			args: args{
				lr: &req,
			},
			want:    res,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Controller{
				uc:    tt.fields.uc,
				auth:  tt.fields.auth,
				token: tt.fields.token,
			}
			got, err := c.Login(tt.args.lr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Login() got = %v, want %v", got, tt.want)
			}
		})
	}
}
