package user

import (
	"reflect"
	"testing"

	"github.com/dwethmar/atami/pkg/domain/entity"
	"github.com/stretchr/testify/assert"
)

type repoTestDependencies struct {
	users []*User
}

func newRepoTestDependencies() *repoTestDependencies {
	return &repoTestDependencies{
		users: []*User{
			{
				ID:        entity.ID(1),
				UID:       entity.NewUID(),
				Username:  "user1",
				Email:     "user1@user.nl",
				Password:  "abdefABCDEF1234!@#$",
				Biography: "biography text",
				CreatedAt: entity.Now(),
				UpdatedAt: entity.Now(),
			},
			{
				ID:        entity.ID(2),
				UID:       entity.NewUID(),
				Username:  "user2",
				Email:     "user2@user.nl",
				Password:  "abdefABCDEF1234!@#$2",
				Biography: "biography text",
				CreatedAt: entity.Now(),
				UpdatedAt: entity.Now(),
			},
			{
				ID:        entity.ID(3),
				UID:       entity.NewUID(),
				Username:  "user3",
				Email:     "user3@user.nl",
				Password:  "abdefABCDEF1234!@#$2",
				Biography: "biography text",
				CreatedAt: entity.Now(),
				UpdatedAt: entity.Now(),
			},
			{
				ID:        entity.ID(4),
				UID:       entity.NewUID(),
				Username:  "user4",
				Email:     "user4@user.nl",
				Password:  "abdefABCDEF1234!@#$2",
				Biography: "biography text",
				CreatedAt: entity.Now(),
				UpdatedAt: entity.Now(),
			},
			{
				ID:        entity.ID(5),
				UID:       entity.NewUID(),
				Username:  "user5",
				Email:     "user5@user.nl",
				Password:  "abdefABCDEF1234!@#$2",
				Biography: "biography text",
				CreatedAt: entity.Now(),
				UpdatedAt: entity.Now(),
			},
		},
	}
}

type setupRepository = func() Repository

func testRepositoryGet(t *testing.T, dependencies *repoTestDependencies, setup setupRepository) {
	testUser := dependencies.users[0]

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
			name: "successfully get user by ID",
			fields: fields{
				repo: setup(),
			},
			args: args{
				ID: testUser.ID,
			},
			want: &User{
				ID:        testUser.ID,
				UID:       testUser.UID,
				Username:  testUser.Username,
				Email:     testUser.Email,
				Password:  testUser.Password,
				Biography: testUser.Biography,
				CreatedAt: testUser.CreatedAt,
				UpdatedAt: testUser.UpdatedAt,
			},
			wantErr: false,
		},
		{
			name: "fail on user not found",
			fields: fields{
				repo: setup(),
			},
			args: args{
				ID: entity.ID(9999),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.fields.repo

			got, err := r.Get(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("RepositoryGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RepositoryGet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func testRepositoryGetByUID(t *testing.T, dependencies *repoTestDependencies, setup setupRepository) {
	testUser := dependencies.users[0]

	type fields struct {
		repo Repository
	}
	type args struct {
		UID entity.UID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *User
		wantErr bool
	}{
		{
			name: "successfully get user by UID",
			fields: fields{
				repo: setup(),
			},
			args: args{
				UID: testUser.UID,
			},
			want: &User{
				ID:        testUser.ID,
				UID:       testUser.UID,
				Username:  testUser.Username,
				Email:     testUser.Email,
				Password:  "",
				Biography: testUser.Biography,
				CreatedAt: testUser.CreatedAt,
				UpdatedAt: testUser.UpdatedAt,
			},
			wantErr: false,
		},
		{
			name: "fail on user not found",
			fields: fields{
				repo: setup(),
			},
			args: args{
				UID: "abc",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.fields.repo

			got, err := r.GetByUID(tt.args.UID)
			if (err != nil) != tt.wantErr {
				t.Errorf("RepositoryGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RepositoryGet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func testRepositoryList(t *testing.T, dependencies *repoTestDependencies, setup setupRepository) {
	testUsers := dependencies.users

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
			name: "Successfully get users",
			fields: fields{
				repo: setup(),
			},
			args: args{
				limit:  10,
				offset: 0,
			},
			want: []*User{
				testUsers[len(testUsers)-1],
				testUsers[len(testUsers)-2],
				testUsers[len(testUsers)-3],
				testUsers[len(testUsers)-4],
				testUsers[len(testUsers)-5],
			},
			wantErr: false,
		},
		{
			name: "Successfully get no users",
			fields: fields{
				repo: setup(),
			},
			args: args{
				limit:  10,
				offset: 10,
			},
			want:    []*User{},
			wantErr: false,
		},
		{
			name: "Successfully get paged users",
			fields: fields{
				repo: setup(),
			},
			args: args{
				limit:  4,
				offset: 0,
			},
			want: []*User{
				testUsers[len(testUsers)-1],
				testUsers[len(testUsers)-2],
				testUsers[len(testUsers)-3],
				testUsers[len(testUsers)-4],
			},
			wantErr: false,
		},
		{
			name: "Successfully get paged users with offset",
			fields: fields{
				repo: setup(),
			},
			args: args{
				limit:  3,
				offset: 1,
			},
			want: []*User{
				testUsers[len(testUsers)-2],
				testUsers[len(testUsers)-3],
				testUsers[len(testUsers)-4],
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := tt.fields.repo
			got, err := repo.List(tt.args.limit, tt.args.offset)

			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			for i, msg := range got {
				if !assert.Equal(t, tt.want[i], msg) {
					t.Errorf("Repository.List() = \n%v, want \n%v", msg, tt.want[i])
					return
				}
			}
		})
	}
}

func testRepositoryCreate(t *testing.T, dependencies *repoTestDependencies, setup setupRepository) {
	type fields struct {
		repo Repository
	}
	type args struct {
		e *User
	}
	// check duplicate email / username
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.ID
		wantErr bool
	}{
		{
			name: "Successfully create user",
			fields: fields{
				repo: setup(),
			},
			args: args{
				e: &User{
					UID:       "abcdefg12334",
					Username:  "testuser",
					Email:     "testuser@test.nl",
					Password:  "abcdefg12334",
					Biography: "biography",
					CreatedAt: entity.Now(),
					UpdatedAt: entity.Now(),
				},
			},
			want:    dependencies.users[len(dependencies.users)-1].ID + 1,
			wantErr: false,
		},
		{
			name: "Fail on duplicate username",
			fields: fields{
				repo: setup(),
			},
			args: args{
				e: &User{
					UID:       "abcdefg12334",
					Username:  dependencies.users[len(dependencies.users)-1].Username,
					Email:     "testuser@test.nl",
					Password:  "abcdefg12334",
					Biography: "biography",
					CreatedAt: entity.Now(),
					UpdatedAt: entity.Now(),
				},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Fail on duplicate email",
			fields: fields{
				repo: setup(),
			},
			args: args{
				e: &User{
					UID:       "abcdefg12334",
					Username:  "username12",
					Email:     dependencies.users[len(dependencies.users)-1].Email,
					Password:  "abcdefg12334",
					Biography: "biography",
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
			r := tt.fields.repo

			got, err := r.Create(tt.args.e)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func testRepositoryUpdate(t *testing.T, dependencies *repoTestDependencies, setup setupRepository) {
	type fields struct {
		repo Repository
	}
	type args struct {
		e *User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Successfully update user",
			fields: fields{
				repo: setup(),
			},
			args: args{
				e: &User{
					ID:        dependencies.users[len(dependencies.users)-1].ID,
					UID:       "abcdefg12334",
					Username:  "updated_sername12",
					Email:     "uupdated_sername12@test.nl",
					Password:  "updated_password",
					Biography: "updated biography",
					CreatedAt: entity.Now(),
					UpdatedAt: entity.Now(),
				},
			},
			wantErr: false,
		},
		{
			name: "Fail on duplicate username",
			fields: fields{
				repo: setup(),
			},
			args: args{
				e: &User{
					UID:       "abcdefg12334",
					Username:  dependencies.users[len(dependencies.users)-1].Username,
					Email:     "testuser@test.nl",
					Password:  "abcdefg12334",
					Biography: "biography",
					CreatedAt: entity.Now(),
					UpdatedAt: entity.Now(),
				},
			},
			wantErr: true,
		},
		{
			name: "Fail on duplicate email",
			fields: fields{
				repo: setup(),
			},
			args: args{
				e: &User{
					UID:       "abcdefg12334",
					Username:  dependencies.users[len(dependencies.users)-1].Email,
					Email:     "testuser@test.nl",
					Password:  "abcdefg12334",
					Biography: "biography",
					CreatedAt: entity.Now(),
					UpdatedAt: entity.Now(),
				},
			},
			wantErr: true,
		},
		{
			name: "Fail on user not found",
			fields: fields{
				repo: setup(),
			},
			args: args{
				e: &User{
					ID:        entity.ID(999),
					UID:       "abcdefg12334",
					Username:  "updated_sername12",
					Email:     "updated_sername12@test.nl",
					Password:  "updated_password",
					Biography: "updated biography",
					CreatedAt: entity.Now(),
					UpdatedAt: entity.Now(),
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.fields.repo

			if err := r.Update(tt.args.e); (err != nil) != tt.wantErr {
				t.Errorf("Repository.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func testRepositoryDelete(t *testing.T, dependencies *repoTestDependencies, setup setupRepository) {
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.fields.repo

			if err := r.Delete(tt.args.ID); (err != nil) != tt.wantErr {
				t.Errorf("Repository.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
