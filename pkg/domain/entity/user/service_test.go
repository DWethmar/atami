package user

import (
	"errors"
	"reflect"
	"testing"

	"github.com/dwethmar/atami/pkg/domain/entity"
	"github.com/dwethmar/atami/pkg/memstore"
)

func Test_errValidate_Valid(t *testing.T) {
	type fields struct {
		Errors []error
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "create valid errValidate",
			fields: fields{
				Errors: []error{},
			},
			want: true,
		},
		{
			name: "create invalid errValidate",
			fields: fields{
				Errors: []error{errors.New("error")},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := errValidate{
				Errors: tt.fields.Errors,
			}
			if got := err.Valid(); got != tt.want {
				t.Errorf("errValidate.Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errValidate_Error(t *testing.T) {
	type fields struct {
		Errors []error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "create valid errValidate",
			fields: fields{
				Errors: []error{},
			},
			want: "",
		},
		{
			name: "create invalid errValidate",
			fields: fields{
				Errors: []error{errors.New("error1"), errors.New("error2")},
			},
			want: "error1. error2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := errValidate{
				Errors: tt.fields.Errors,
			}
			if got := err.Error(); got != tt.want {
				t.Errorf("errValidate.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_Create(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		e *Create
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.ID
		wantErr bool
	}{
		{
			name: "create new user",
			fields: fields{
				repo: NewInMemoryRepo(memstore.NewStore()),
			},
			args: args{
				e: &Create{
					UID       : entity.NewUID(),
					Username  : "username",
					Email     :"test@test.nl",
					Password  :"lkasjDSlsjk*(&^*&(jhggjkh11",
					Biography :"bio"	,
					CreatedAt :entity.Now(),
					UpdatedAt :entity.Now(),
				},
			},
			want: 1,
			wantErr: false,
		},
		{
			name: "create invalid user with invalid password",
			fields: fields{
				repo: NewInMemoryRepo(memstore.NewStore()),
			},
			args: args{
				e: &Create{
					UID       : entity.NewUID(),
					Username  : "kipsate",
					Email     :"test@test.nl",
					Password  :"12",
					Biography :"bio"	,
					CreatedAt :entity.Now(),
					UpdatedAt :entity.Now(),
				},
			},
			want: 0,
			wantErr: true,
		},
		{
			name: "create invalid user with invalid username",
			fields: fields{
				repo: NewInMemoryRepo(memstore.NewStore()),
			},
			args: args{
				e: &Create{
					UID       : entity.NewUID(),
					Username  : "a@",
					Email     :"test@test.nl",
					Password  :"lkasjDSlsjk*(&^*&(jhggjkh11",
					Biography :"bio"	,
					CreatedAt :entity.Now(),
					UpdatedAt :entity.Now(),
				},
			},
			want: 0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repo: tt.fields.repo,
			}
			got, err := s.Create(tt.args.e)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_Get(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		ID entity.ID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repo: tt.fields.repo,
			}
			got, err := s.Get(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_List(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		limit  uint
		offset uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repo: tt.fields.repo,
			}
			got, err := s.List(tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_Delete(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		id entity.ID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repo: tt.fields.repo,
			}
			if err := s.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Service.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_Update(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		ID entity.ID
		e  *Update
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repo: tt.fields.repo,
			}
			if err := s.Update(tt.args.ID, tt.args.e); (err != nil) != tt.wantErr {
				t.Errorf("Service.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_Authenticate(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		email    string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Authenticated
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repo: tt.fields.repo,
			}
			got, err := s.Authenticate(tt.args.email, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.Authenticate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.Authenticate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_ValidateCreate(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		c *Create
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repo: tt.fields.repo,
			}
			if err := s.ValidateCreate(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Service.ValidateCreate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_ValidateUpdate(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		u *Update
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repo: tt.fields.repo,
			}
			if err := s.ValidateUpdate(tt.args.u); (err != nil) != tt.wantErr {
				t.Errorf("Service.ValidateUpdate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
