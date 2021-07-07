package user

import (
	"errors"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/dwethmar/atami/pkg/domain/entity"
	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/stretchr/testify/assert"
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
				t.Errorf("user errValidate.Valid() = %v, want %v", got, tt.want)
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
				t.Errorf("user errValidate.Error() = %v, want %v", got, tt.want)
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
					UID:       entity.NewUID(),
					Username:  "username",
					Email:     "test@test.nl",
					Password:  "lkasjDSlsjk*(&^*&(jhggjkh11",
					Biography: "bio",
					CreatedAt: entity.Now(),
					UpdatedAt: entity.Now(),
				},
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "create invalid user with invalid password",
			fields: fields{
				repo: NewInMemoryRepo(memstore.NewStore()),
			},
			args: args{
				e: &Create{
					UID:       entity.NewUID(),
					Username:  "kipsate",
					Email:     "test@test.nl",
					Password:  "12",
					Biography: "bio",
					CreatedAt: entity.Now(),
					UpdatedAt: entity.Now(),
				},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "create invalid user with invalid username",
			fields: fields{
				repo: NewInMemoryRepo(memstore.NewStore()),
			},
			args: args{
				e: &Create{
					UID:       entity.NewUID(),
					Username:  "a@",
					Email:     "test@test.nl",
					Password:  "lkasjDSlsjk*(&^*&(jhggjkh11",
					Biography: "bio",
					CreatedAt: entity.Now(),
					UpdatedAt: entity.Now(),
				},
			},
			want:    0,
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
				t.Errorf("user Service.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("user Service.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_Get(t *testing.T) {
	fixtures := newTestFixtures()
	setup := func() *memstore.Memstore {
		s := memstore.NewStore()
		for _, u := range fixtures.users {
			s.GetUsers().Put(u.ID, *toMemory(u))
		}
		return s
	}

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
		{
			name: "get user by ID",
			fields: fields{
				repo: NewInMemoryRepo(setup()),
			},
			args: args{
				ID: 1,
			},
			want: func() *User {
				f := fixtures.users[0]
				f.Password = ""
				return f
			}(),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repo: tt.fields.repo,
			}
			got, err := s.Get(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("user Service.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("user Service.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_List(t *testing.T) {
	fixtures := newTestFixtures()
	setup := func() *memstore.Memstore {
		s := memstore.NewStore()
		for _, u := range fixtures.users {
			s.GetUsers().Put(u.ID, *toMemory(u))
		}
		return s
	}

	// remove password from fixtures
	prepareUsers := func() []*User {
		rr := make([]*User, len(fixtures.users))
		for i, u := range fixtures.users {
			p := *u
			p.Password = ""
			rr[i] = &p
		}
		return rr
	}

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
		{
			name: "List users",
			fields: fields{
				repo: NewInMemoryRepo(setup()),
			},
			args: args{
				limit:  10,
				offset: 0,
			},
			want: func() []*User {
				u := prepareUsers()
				return []*User{
					u[4],
					u[3],
					u[2],
					u[1],
					u[0],
				}
			}(),
			wantErr: false,
		},
		{
			name: "No error on no users",
			fields: fields{
				repo: NewInMemoryRepo(memstore.NewStore()),
			},
			args: args{
				limit:  10,
				offset: 0,
			},
			want:    []*User{},
			wantErr: false,
		},
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

			for i, u := range got {
				if !assert.Equal(t, u, tt.want[i]) {
					t.Errorf("Service.List() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestService_Delete(t *testing.T) {
	fixtures := newTestFixtures()
	setup := func() *memstore.Memstore {
		s := memstore.NewStore()
		for _, u := range fixtures.users {
			s.GetUsers().Put(u.ID, *toMemory(u))
		}
		return s
	}

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
		wantErr bool
	}{
		{
			name: "successfully delete user",
			fields: fields{
				repo: NewInMemoryRepo(setup()),
			},
			args: args{
				ID: 1,
			},
			wantErr: false,
		},
		{
			name: "error on none existing user",
			fields: fields{
				repo: NewInMemoryRepo(setup()),
			},
			args: args{
				ID: 999,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repo: tt.fields.repo,
			}
			if err := s.Delete(tt.args.ID); (err != nil) != tt.wantErr {
				t.Errorf("Service.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_Update(t *testing.T) {
	fixtures := newTestFixtures()
	setup := func() *memstore.Memstore {
		s := memstore.NewStore()
		for _, u := range fixtures.users {
			s.GetUsers().Put(u.ID, *toMemory(u))
		}
		return s
	}

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
		{
			name: "successfully update user",
			fields: fields{
				repo: NewInMemoryRepo(setup()),
			},
			args: args{
				ID: 1,
				e: &Update{
					Biography: "updated biography",
				},
			},
			wantErr: false,
		},
		{
			name: "error on none existing user",
			fields: fields{
				repo: NewInMemoryRepo(setup()),
			},
			args: args{
				ID: 999,
				e: &Update{
					Biography: "updated biography",
				},
			},
			wantErr: true,
		},
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
	fixtures := newTestFixtures()
	setup := func() *memstore.Memstore {
		s := memstore.NewStore()
		for _, u := range fixtures.users {
			s.GetUsers().Put(u.ID, *toMemory(u))
		}
		return s
	}

	now := time.Now()
	testUser := fixtures.users[0]

	type fields struct {
		repo Repository
	}
	type args struct {
		email    string
		password string
		issuedAt time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Authenticated
		wantErr bool
	}{
		{
			name: "Successfully authenticate",
			fields: fields{
				repo: NewInMemoryRepo(setup()),
			},
			args: args{
				email:    testUser.Email,
				password: "abdefABCDEF1234!@#$",
				issuedAt: now,
			},
			want: &Authenticated{
				AccessToken: func() string {
					token, err := CreateAccessToken(testUser.UID, strconv.FormatInt(now.UnixNano(), 10), now, now.Add(time.Minute*60))
					assert.NoError(t, err)
					return token
				}(),
				RefreshToken: func() string {
					token, err := CreateRefreshToken(testUser.UID, strconv.FormatInt(now.UnixNano(), 10), now, now.Add(time.Hour*730))
					assert.NoError(t, err)
					return token
				}(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repo: tt.fields.repo,
			}
			got, err := s.Authenticate(tt.args.email, tt.args.password, now)
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
	fixtures := newTestFixtures()
	setup := func() *memstore.Memstore {
		s := memstore.NewStore()
		for _, u := range fixtures.users {
			s.GetUsers().Put(u.ID, *toMemory(u))
		}
		return s
	}

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
		{
			name: "Successfully validate valid user create",
			fields: fields{
				repo: NewInMemoryRepo(setup()),
			},
			args: args{
				c: &Create{
					UID:       entity.NewUID(),
					Username:  "username",
					Email:     "test@test.nl",
					Password:  "aQAc!@#skk111",
					Biography: "bio",
					CreatedAt: entity.Now(),
					UpdatedAt: entity.Now(),
				},
			},
			wantErr: false,
		},
		{
			name: "Error on invalid user create",
			fields: fields{
				repo: NewInMemoryRepo(setup()),
			},
			args: args{
				c: &Create{
					UID:       entity.NewUID(),
					Username:  "",
					Email:     "test@test.nl",
					Password:  "a",
					Biography: "bio",
					CreatedAt: entity.Now(),
					UpdatedAt: entity.Now(),
				},
			},
			wantErr: true,
		},
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
	fixtures := newTestFixtures()
	setup := func() *memstore.Memstore {
		s := memstore.NewStore()
		for _, u := range fixtures.users {
			s.GetUsers().Put(u.ID, *toMemory(u))
		}
		return s
	}

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
		{
			name: "Successfully validate valid user create",
			fields: fields{
				repo: NewInMemoryRepo(setup()),
			},
			args: args{
				u: &Update{
					Biography: "updated Biography",
				},
			},
			wantErr: false,
		},
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
