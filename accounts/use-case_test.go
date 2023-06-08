package accounts

import (
	"CRM-Services-Task/accounts/dto"
	"CRM-Services-Task/accounts/entity"
	mocks "CRM-Services-Task/mocks/accounts"
	"errors"
	"reflect"
	"testing"
)

func TestUseCase_Login(t *testing.T) {
	type fields struct {
		repo IRepo
	}
	type args struct {
		a  *entity.Actor
		lr *dto.LoginRequest
	}
	mockRepo := mocks.NewIRepo(t)
	req := dto.LoginRequest{
		Username: "lintang",
		Password: "pass",
	}
	err := errors.New("login error")
	actor := &entity.Actor{}
	mockRepo.EXPECT().Login(actor, &req).Return(err).Once()

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "login sukses",
			fields: fields{
				repo: mockRepo,
			},
			args: args{
				a:  actor,
				lr: &req,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UseCase{
				repo: tt.fields.repo,
			}
			if err := u.Login(tt.args.a, tt.args.lr); (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseCase_CreateAdmin(t *testing.T) {
	type fields struct {
		repo IRepo
	}
	type args struct {
		a *entity.Actor
	}
	mockRepo := mocks.NewIRepo(t)
	actor := entity.Actor{}
	mockRepo.EXPECT().CreateAdmin(&actor).Return(errors.New("ada error")).Once()
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "create return error",
			fields: fields{
				repo: mockRepo,
			},
			args: args{
				a: &actor,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UseCase{
				repo: tt.fields.repo,
			}
			if err := u.CreateAdmin(tt.args.a); (err != nil) != tt.wantErr {
				t.Errorf("CreateAdmin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseCase_GetApprovalList(t *testing.T) {
	type fields struct {
		repo IRepo
	}
	mockRepo := mocks.NewIRepo(t)
	actor := []entity.Actor{}
	err := errors.New("ada error")
	mockRepo.EXPECT().FindAllApproval().Return(actor, err).Once()

	mockRepo.EXPECT().FindAllApproval().Return(actor, nil).Once()

	tests := []struct {
		name    string
		fields  fields
		want    []entity.Actor
		wantErr bool
	}{
		{
			name: "get all with error",
			fields: fields{
				repo: mockRepo,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "sukses",
			fields: fields{
				repo: mockRepo,
			},
			want:    actor,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UseCase{
				repo: tt.fields.repo,
			}
			got, err := u.GetApprovalList()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetApprovalList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetApprovalList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_ApproveAdmin(t *testing.T) {
	type fields struct {
		repo IRepo
	}
	mockRepo := mocks.NewIRepo(t)

	//update actor gagal
	actor := entity.Actor{}
	ar := dto.ApproveRequest{}
	ra := entity.RegisterApproval{}
	valueToUpdate := map[string]interface{}{"verified": "true"}
	status := map[string]interface{}{"status": "approved"}
	err := errors.New("ada error")
	mockRepo.EXPECT().UpdateActor(&actor, &ar, valueToUpdate).Return(err).Once()

	mockRepo.EXPECT().UpdateActor(&actor, &ar, valueToUpdate).Return(nil).Once()
	mockRepo.EXPECT().UpdateStatusRA(&ra, &ar, status).Return(err).Once()

	mockRepo.EXPECT().UpdateActor(&actor, &ar, valueToUpdate).Return(nil).Once()
	mockRepo.EXPECT().UpdateStatusRA(&ra, &ar, status).Return(nil).Once()

	type args struct {
		a      *entity.Actor
		ar     *dto.ApproveRequest
		ra     *entity.RegisterApproval
		update map[string]interface{}
		ra2    map[string]interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "gagal update actor",
			fields: fields{
				repo: mockRepo,
			},
			args: args{
				a:      &actor,
				ar:     &ar,
				ra:     &ra,
				update: valueToUpdate,
				ra2:    status,
			},
			wantErr: true,
		},
		{
			name: "gagal update status approval",
			fields: fields{
				repo: mockRepo,
			},
			args: args{
				a:      &actor,
				ar:     &ar,
				ra:     &ra,
				update: valueToUpdate,
				ra2:    status,
			},
			wantErr: true,
		},
		{
			name: "sukses",
			fields: fields{
				repo: mockRepo,
			},
			args: args{
				a:      &actor,
				ar:     &ar,
				ra:     &ra,
				update: valueToUpdate,
				ra2:    status,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UseCase{
				repo: tt.fields.repo,
			}
			if err := u.ApproveAdmin(tt.args.a, tt.args.ar, tt.args.ra, tt.args.update, tt.args.ra2); (err != nil) != tt.wantErr {
				t.Errorf("ApproveAdmin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseCase_RejectAdmin(t *testing.T) {
	type fields struct {
		repo IRepo
	}
	type args struct {
		a       *entity.Actor
		ar      *dto.ApproveRequest
		ra      *entity.RegisterApproval
		update  map[string]interface{}
		valueRa map[string]interface{}
	}

	mockRepo := mocks.NewIRepo(t)
	actor := entity.Actor{}
	ar := dto.ApproveRequest{}
	ra := entity.RegisterApproval{}
	val := map[string]interface{}{"verif": "true"}
	err := errors.New("ada error")
	mockRepo.EXPECT().UpdateActor(&actor, &ar, val).Return(err).Once()

	mockRepo.EXPECT().UpdateActor(&actor, &ar, val).Return(nil).Once()
	mockRepo.EXPECT().UpdateStatusRA(&ra, &ar, val).Return(err).Once()

	mockRepo.EXPECT().UpdateActor(&actor, &ar, val).Return(nil).Once()
	mockRepo.EXPECT().UpdateStatusRA(&ra, &ar, val).Return(nil).Once()

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "update aktor gagal",
			fields: fields{
				repo: mockRepo,
			},
			args: args{
				a:       &actor,
				ar:      &ar,
				ra:      &ra,
				update:  val,
				valueRa: nil,
			},
			wantErr: true,
		},
		{
			name: "update RA gagal",
			fields: fields{
				repo: mockRepo,
			},
			args: args{
				a:       &actor,
				ar:      &ar,
				ra:      &ra,
				update:  val,
				valueRa: val,
			},
			wantErr: true,
		},
		{
			name: "update sukses",
			fields: fields{
				repo: mockRepo,
			},
			args: args{
				a:       &actor,
				ar:      &ar,
				ra:      &ra,
				update:  val,
				valueRa: val,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UseCase{
				repo: tt.fields.repo,
			}
			if err := u.RejectAdmin(tt.args.a, tt.args.ar, tt.args.ra, tt.args.update, tt.args.valueRa); (err != nil) != tt.wantErr {
				t.Errorf("RejectAdmin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseCase_ActivateAdmin(t *testing.T) {
	type fields struct {
		repo IRepo
	}
	type args struct {
		a          *entity.Actor
		aar        *dto.ActivateAdminRequest
		activeTrue map[string]interface{}
	}
	mockRepo := mocks.NewIRepo(t)
	actor := entity.Actor{}
	aar := dto.ActivateAdminRequest{}
	val := map[string]interface{}{"active": true}
	err := errors.New("ada error")
	mockRepo.EXPECT().ActivateAdmin(&actor, &aar, val).Return(err).Once()

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "sukses update",
			fields: fields{
				repo: mockRepo,
			},
			args: args{
				a:          &actor,
				aar:        &aar,
				activeTrue: val,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UseCase{
				repo: tt.fields.repo,
			}
			if err := u.ActivateAdmin(tt.args.a, tt.args.aar, tt.args.activeTrue); (err != nil) != tt.wantErr {
				t.Errorf("ActivateAdmin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseCase_DeactivateAdmin(t *testing.T) {
	type fields struct {
		repo IRepo
	}
	type args struct {
		a   *entity.Actor
		d   *dto.DeActivateAdminRequest
		val map[string]interface{}
	}
	mockRepo := mocks.NewIRepo(t)
	actor := entity.Actor{}
	dar := dto.DeActivateAdminRequest{}
	val := map[string]interface{}{"active": true}
	err := errors.New("ada error")
	mockRepo.EXPECT().DeactivateAdmin(&actor, &dar, val).Return(err).Once()

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "sukses deactivate",
			fields: fields{
				repo: mockRepo,
			},
			args: args{
				a:   &actor,
				d:   &dar,
				val: val,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UseCase{
				repo: tt.fields.repo,
			}
			if err := u.DeactivateAdmin(tt.args.a, tt.args.d, tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("DeactivateAdmin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseCase_DeleteAdmin(t *testing.T) {
	type fields struct {
		repo IRepo
	}
	type args struct {
		a *entity.Actor
		d *dto.DeleteAdminRequest
	}
	mockRepo := mocks.NewIRepo(t)
	actor := entity.Actor{}
	dar := dto.DeleteAdminRequest{}

	err := errors.New("ada error")
	mockRepo.EXPECT().DeleteAdmin(&actor, &dar).Return(err).Once()

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "delete admin sukses",
			fields: fields{
				repo: mockRepo,
			},
			args: args{
				a: &actor,
				d: &dar,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UseCase{
				repo: tt.fields.repo,
			}
			if err := u.DeleteAdmin(tt.args.a, tt.args.d); (err != nil) != tt.wantErr {
				t.Errorf("DeleteAdmin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseCase_GetActorsByUsername(t *testing.T) {
	type fields struct {
		repo IRepo
	}
	type args struct {
		username string
		limit    string
		page     string
	}
	mockRepo := mocks.NewIRepo(t)

	err := errors.New("ada error")
	actor := []entity.Actor{}
	mockRepo.EXPECT().FindAllActors("lintang", "1", "1").Return(actor, err).Once()
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entity.Actor
		wantErr bool
	}{
		{
			name: "get all sukses",
			fields: fields{
				repo: mockRepo,
			},
			args: args{
				username: "lintang",
				limit:    "1",
				page:     "1",
			},
			want:    actor,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UseCase{
				repo: tt.fields.repo,
			}
			got, err := u.GetActorsByUsername(tt.args.username, tt.args.limit, tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetActorsByUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetActorsByUsername() got = %v, want %v", got, tt.want)
			}
		})
	}
}
