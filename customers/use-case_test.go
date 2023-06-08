package customers

import (
	"CRM-Services-Task/customers/dto"
	"CRM-Services-Task/customers/entity"
	mocks "CRM-Services-Task/mocks/customers"
	"errors"
	"reflect"
	"testing"
)

func TestUseCase_CreateCustomer(t *testing.T) {
	type fields struct {
		repo IRepo
	}
	type args struct {
		c *entity.Customer
	}
	mock := mocks.NewIRepo(t)
	cust := entity.Customer{}
	err := errors.New("error")
	mock.EXPECT().CreateCustomer(&cust).Return(err).Once()

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "sukses create",
			fields: fields{
				repo: mock,
			},
			args: args{
				c: &cust,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UseCase{
				repo: tt.fields.repo,
			}
			if err := u.CreateCustomer(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("CreateCustomer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseCase_DeleteCustomer(t *testing.T) {
	type fields struct {
		repo IRepo
	}
	type args struct {
		c *entity.Customer
		d *dto.DeleteCustomerRequest
	}
	mock := mocks.NewIRepo(t)
	c := entity.Customer{}
	d := dto.DeleteCustomerRequest{}
	err := errors.New("err")
	mock.EXPECT().DeleteCustomer(&c, &d).Return(err).Once()

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "sukses del",
			fields: fields{
				repo: mock,
			},
			args: args{
				c: &c,
				d: &d,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UseCase{
				repo: tt.fields.repo,
			}
			if err := u.DeleteCustomer(tt.args.c, tt.args.d); (err != nil) != tt.wantErr {
				t.Errorf("DeleteCustomer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseCase_GetCustomersByNameAndEmail(t *testing.T) {
	type fields struct {
		repo IRepo
	}
	type args struct {
		firstname string
		email     string
		limit     string
		page      string
	}
	mock := mocks.NewIRepo(t)
	cust := []entity.Customer{}
	err := errors.New("err")
	mock.EXPECT().FindAllCustomers("lintang", "lintang@gmail.com", "1", "1").Return(nil, err).Once()

	mock.EXPECT().FindAllCustomers("lintang", "lintang@gmail.com", "1", "1").Return(cust, nil).Once()

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entity.Customer
		wantErr bool
	}{
		{
			name: "error get all",
			fields: fields{
				repo: mock,
			},
			args: args{
				firstname: "lintang",
				email:     "lintang@gmail.com",
				limit:     "1",
				page:      "1",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "sukses get all",
			fields: fields{
				repo: mock,
			},
			args: args{
				firstname: "lintang",
				email:     "lintang@gmail.com",
				limit:     "1",
				page:      "1",
			},
			want:    cust,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UseCase{
				repo: tt.fields.repo,
			}
			got, err := u.GetCustomersByNameAndEmail(tt.args.firstname, tt.args.email, tt.args.limit, tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCustomersByNameAndEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCustomersByNameAndEmail() got = %v, want %v", got, tt.want)
			}
		})
	}
}
