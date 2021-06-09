package user

import (
	"reflect"
	"testing"

	"github.com/dwethmar/atami/pkg/memstore"
)

func Test_inMemoryRepo_Get(t *testing.T) {
	deps := newRepoTestDependencies()
	testRepositoryGet(
		t,
		deps,
		func() Repository {
			store := memstore.NewStore()
			for _, user := range deps.users {
				store.GetUsers().Put(user.ID, *toMemory(user))
			}
			return NewInMemoryRepo(store)
		},
	)
}

func Test_inMemoryRepo_GetByUID(t *testing.T) {
	deps := newRepoTestDependencies()
	testRepositoryGetByUID(
		t,
		deps,
		func() Repository {
			store := memstore.NewStore()
			for _, user := range deps.users {
				store.GetUsers().Put(user.ID, *toMemory(user))
			}
			return NewInMemoryRepo(store)
		},
	)
}

func Test_inMemoryRepo_GetByUsername(t *testing.T) {
	deps := newRepoTestDependencies()
	testRepositoryGetByUsername(
		t,
		deps,
		func() Repository {
			store := memstore.NewStore()
			for _, user := range deps.users {
				store.GetUsers().Put(user.ID, *toMemory(user))
			}
			return NewInMemoryRepo(store)
		},
	)
}

func Test_inMemoryRepo_GetByEmail(t *testing.T) {
	deps := newRepoTestDependencies()
	testRepositoryGetByEmail(
		t,
		deps,
		func() Repository {
			store := memstore.NewStore()
			for _, user := range deps.users {
				store.GetUsers().Put(user.ID, *toMemory(user))
			}
			return NewInMemoryRepo(store)
		},
	)
}

func Test_inMemoryRepo_GetCredentials(t *testing.T) {
	deps := newRepoTestDependencies()
	testRepositoryGetByUsername(
		t,
		deps,
		func() Repository {
			store := memstore.NewStore()
			for _, user := range deps.users {
				store.GetUsers().Put(user.ID, *toMemory(user))
			}
			return NewInMemoryRepo(store)
		},
	)
}

func Test_inMemoryRepo_List(t *testing.T) {
	deps := newRepoTestDependencies()
	testRepositoryList(
		t,
		deps,
		func() Repository {
			store := memstore.NewStore()
			for _, user := range deps.users {
				store.GetUsers().Put(user.ID, *toMemory(user))
			}
			return NewInMemoryRepo(store)
		},
	)
}

func Test_inMemoryRepo_Update(t *testing.T) {
	deps := newRepoTestDependencies()
	testRepositoryUpdate(
		t,
		deps,
		func() Repository {
			store := memstore.NewStore()
			for _, user := range deps.users {
				store.GetUsers().Put(user.ID, *toMemory(user))
			}
			return NewInMemoryRepo(store)
		},
	)
}

func Test_inMemoryRepo_Create(t *testing.T) {
	deps := newRepoTestDependencies()
	testRepositoryCreate(
		t,
		deps,
		func() Repository {
			store := memstore.NewStore()
			for _, user := range deps.users {
				store.GetUsers().Put(user.ID, *toMemory(user))
			}
			return NewInMemoryRepo(store)
		},
	)
}

func Test_inMemoryRepo_Delete(t *testing.T) {
	deps := newRepoTestDependencies()
	testRepositoryDelete(
		t,
		deps,
		func() Repository {
			store := memstore.NewStore()
			for _, user := range deps.users {
				store.GetUsers().Put(user.ID, *toMemory(user))
			}
			return NewInMemoryRepo(store)
		},
	)
}

func Test_toMemory(t *testing.T) {

}

func Test_fromMemory(t *testing.T) {

}

func Test_filterList(t *testing.T) {

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
			got, err := mapCredentials(tt.args.row)
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
