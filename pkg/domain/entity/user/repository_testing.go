package user

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/dwethmar/atami/pkg/database"
	"github.com/dwethmar/atami/pkg/domain/entity"
)


func TestNewPostgresRepository(t *testing.T) {
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name string
		args args
		want Repository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPostgresRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPostgresRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func testRepositoryGetByUID(t *testing.T) {
	type fields struct {
		db database.Transaction
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &postgresRepo{
				db: tt.fields.db,
			}
			got, err := r.GetByUID(tt.args.UID)
			if (err != nil) != tt.wantErr {
				t.Errorf("postgresRepo.GetByUID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postgresRepo.GetByUID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func testRepositoryGet(t *testing.T) {
	type fields struct {
		db database.Transaction
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
			r := &postgresRepo{
				db: tt.fields.db,
			}
			got, err := r.Get(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("postgresRepo.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postgresRepo.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func testRepositoryList(t *testing.T) {
	type fields struct {
		db database.Transaction
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
			r := &postgresRepo{
				db: tt.fields.db,
			}
			got, err := r.List(tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("postgresRepo.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postgresRepo.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func testRepositoryCreate(t *testing.T) {
	type fields struct {
		db database.Transaction
	}
	type args struct {
		e *User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.ID
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &postgresRepo{
				db: tt.fields.db,
			}
			got, err := r.Create(tt.args.e)
			if (err != nil) != tt.wantErr {
				t.Errorf("postgresRepo.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postgresRepo.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func testRepositoryUpdate(t *testing.T) {
	type fields struct {
		db database.Transaction
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &postgresRepo{
				db: tt.fields.db,
			}
			if err := r.Update(tt.args.e); (err != nil) != tt.wantErr {
				t.Errorf("postgresRepo.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func testRepositoryDelete(t *testing.T) {
	type fields struct {
		db database.Transaction
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
			r := &postgresRepo{
				db: tt.fields.db,
			}
			if err := r.Delete(tt.args.ID); (err != nil) != tt.wantErr {
				t.Errorf("postgresRepo.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_defaultMap(t *testing.T) {
	type args struct {
		row Row
	}
	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := defaultMap(tt.args.row)
			if (err != nil) != tt.wantErr {
				t.Errorf("defaultMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("defaultMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mapIsUniqueCheck(t *testing.T) {
	type args struct {
		row Row
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mapIsUniqueCheck(tt.args.row)
			if (err != nil) != tt.wantErr {
				t.Errorf("mapIsUniqueCheck() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("mapIsUniqueCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mapWithPassword(t *testing.T) {
	type args struct {
		row Row
	}
	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mapWithPassword(tt.args.row)
			if (err != nil) != tt.wantErr {
				t.Errorf("mapWithPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapWithPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
