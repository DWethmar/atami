package message

import (
	"reflect"
	"testing"
	"time"

	"github.com/dwethmar/atami/pkg/domain/entity"
	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/stretchr/testify/assert"
)

func Test_inMemoryRepo_Get(t *testing.T) {
	createdAt := time.Now().UTC()
	testUser := *memstore.NewFixtureUser(entity.ID(1))

	type fields struct {
		memStore *memstore.Memstore
	}
	type args struct {
		ID entity.ID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Message
		wantErr bool
	}{
		{
			name: "Successfully get message by ID",
			fields: fields{
				memStore: memstore.NewStore(),
			},
			args: args{
				ID: 1,
			},
			want: &Message{
				ID:              entity.ID(1),
				UID:             "abc123",
				Text:            "test",
				CreatedByUserID: testUser.ID,
				CreatedAt:       createdAt,
				User: User{
					ID:       testUser.ID,
					UID:      testUser.UID,
					Username: testUser.Username,
				},
			},
			wantErr: false,
		},
		{
			name: "Fail get message by unknown UID",
			fields: fields{
				memStore: memstore.NewStore(),
			},
			args: args{
				ID: 2,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := tt.fields.memStore
			r := NewinMemoryRepoRepository(store)
			store.GetUsers().Put(entity.ID(1), testUser)
			r.Create(Create{
				UID:             "abc123",
				Text:            "test",
				CreatedByUserID: testUser.ID,
				CreatedAt:       createdAt,
			})

			got, err := r.Get(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("inMemoryRepo.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !assert.Equal(t, got, tt.want) {
				t.Errorf("inMemoryRepo.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inMemoryRepo_GetByUID(t *testing.T) {
	createdAt := time.Now().UTC()
	testUser := *memstore.NewFixtureUser(entity.ID(1))

	type fields struct {
		memStore *memstore.Memstore
	}
	type args struct {
		UID entity.UID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Message
		wantErr bool
	}{
		{
			name: "Successfully get message by UID",
			fields: fields{
				memStore: memstore.NewStore(),
			},
			args: args{
				UID: "xyz123",
			},
			want: &Message{
				ID:              entity.ID(1),
				UID:             "xyz123",
				Text:            "test",
				CreatedByUserID: testUser.ID,
				CreatedAt:       createdAt,
				User: User{
					ID:       testUser.ID,
					UID:      testUser.UID,
					Username: testUser.Username,
				},
			},
			wantErr: false,
		},
		{
			name: "Fail get message by unknown UID",
			fields: fields{
				memStore: memstore.NewStore(),
			},
			args: args{
				UID: "abc123",
			},
			want:    nil,
			wantErr: true,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := tt.fields.memStore
			r := NewinMemoryRepoRepository(store)
			store.GetUsers().Put(entity.ID(1), testUser)
			r.Create(Create{
				UID:             "xyz123",
				Text:            "test",
				CreatedByUserID: testUser.ID,
				CreatedAt:       createdAt,
			})

			got, err := r.GetByUID(tt.args.UID)
			if (err != nil) != tt.wantErr {
				t.Errorf("inMemoryRepo.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !assert.Equal(t, got, tt.want) {
				t.Errorf("inMemoryRepo.GetByUID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inMemoryRepo_List(t *testing.T) {
	createdAt := time.Now().UTC()
	testUser := *memstore.NewFixtureUser(entity.ID(1))
	testUser2 := *memstore.NewFixtureUser(entity.ID(2))

	type fields struct {
		memStore *memstore.Memstore
	}
	type args struct {
		limit  uint
		offset uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*Message
		wantErr bool
	}{
		{
			name: "Successfully get messages",
			fields: fields{
				memStore: memstore.NewStore(),
			},
			args: args{
				limit:  10,
				offset: 0,
			},
			want: []*Message{
				{
					ID:              entity.ID(1),
					UID:             "xyz123",
					Text:            "test text 1",
					CreatedByUserID: testUser.ID,
					CreatedAt:       createdAt,
					User: User{
						ID:       testUser.ID,
						UID:      testUser.UID,
						Username: testUser.Username,
					},
				},
				{
					ID:              entity.ID(2),
					UID:             "mnop678",
					Text:            "test text 2",
					CreatedByUserID: testUser2.ID,
					CreatedAt:       createdAt,
					User: User{
						ID:       testUser2.ID,
						UID:      testUser2.UID,
						Username: testUser2.Username,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Successfully get no messages",
			fields: fields{
				memStore: memstore.NewStore(),
			},
			args: args{
				limit:  10,
				offset: 10,
			},
			want:    []*Message{},
			wantErr: false,
		},
		{
			name: "Successfully get paged messages",
			fields: fields{
				memStore: memstore.NewStore(),
			},
			args: args{
				limit:  10,
				offset: 1,
			},
			want: []*Message{
				{
					ID:              entity.ID(2),
					UID:             "mnop678",
					Text:            "test text 2",
					CreatedByUserID: testUser2.ID,
					CreatedAt:       createdAt,
					User: User{
						ID:       testUser2.ID,
						UID:      testUser2.UID,
						Username: testUser2.Username,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := tt.fields.memStore
			r := NewinMemoryRepoRepository(store)
			store.GetUsers().Put(entity.ID(1), testUser)
			store.GetUsers().Put(entity.ID(2), testUser2)

			r.Create(Create{
				UID:             "xyz123",
				Text:            "test text 1",
				CreatedByUserID: testUser.ID,
				CreatedAt:       createdAt,
			})

			r.Create(Create{
				UID:             "mnop678",
				Text:            "test text 2",
				CreatedByUserID: testUser2.ID,
				CreatedAt:       createdAt,
			})

			got, err := r.List(tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("inMemoryRepo.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !assert.Equal(t, tt.want, got) {
				t.Errorf("inMemoryRepo.List() = \n%v, want \n%v", got, tt.want)
				return
			}
		})
	}
}

func Test_inMemoryRepo_Update(t *testing.T) {
	createdAt := time.Now().UTC()
	testUser := *memstore.NewFixtureUser(entity.ID(1))

	type fields struct {
		memStore *memstore.Memstore
	}
	type args struct {
		update Update
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Successfully update message",
			fields: fields{
				memStore: memstore.NewStore(),
			},
			args: args{
				update: Update{
					Text: "updated text",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := tt.fields.memStore
			r := NewinMemoryRepoRepository(store)
			store.GetUsers().Put(entity.ID(1), testUser)

			ID, _ := r.Create(Create{
				UID:             "xyz123",
				Text:            "test text 1",
				CreatedByUserID: testUser.ID,
				CreatedAt:       createdAt,
			})

			expectedMsg, _ := r.Get(ID)
			expectedMsg.Apply(tt.args.update)

			if err := r.Update(ID, tt.args.update); (err != nil) != tt.wantErr {
				t.Errorf("inMemoryRepo.Update() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				updatedMsg, _ := r.Get(ID)
				if !assert.Equal(t, updatedMsg, expectedMsg) {
					t.Errorf("inMemoryRepo.Update() = \n%v, want \n%v", expectedMsg, expectedMsg)
				}
			}
		})
	}
}

func Test_inMemoryRepo_Create(t *testing.T) {
	createdAt := time.Now().UTC()
	testUser := *memstore.NewFixtureUser(entity.ID(1))

	type fields struct {
		memStore *memstore.Memstore
		newID    entity.ID
	}
	type args struct {
		create Create
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.ID
		wantErr bool
	}{
		{
			name: "Successfully create message",
			fields: fields{
				memStore: memstore.NewStore(),
			},
			args: args{
				create: Create{
					UID:             "abc123",
					Text:            "updated text",
					CreatedByUserID: entity.ID(1),
					CreatedAt:       createdAt,
				},
			},
			want:    entity.ID(1),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := tt.fields.memStore
			r := NewinMemoryRepoRepository(store)
			store.GetUsers().Put(entity.ID(1), testUser)

			got, err := r.Create(tt.args.create)
			if (err != nil) != tt.wantErr {
				t.Errorf("inMemoryRepo.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inMemoryRepo.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inMemoryRepo_Delete(t *testing.T) {
	createdAt := time.Now().UTC()
	testUser := *memstore.NewFixtureUser(entity.ID(1))

	type fields struct {
		memStore *memstore.Memstore
	}
	type args struct {
		create Create
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Successfully delete message",
			fields: fields{
				memStore: memstore.NewStore(),
			},
			args: args{
				create: Create{
					UID:             "abc123",
					Text:            "updated text",
					CreatedByUserID: entity.ID(1),
					CreatedAt:       createdAt,
				},
			},
			wantErr: false,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := tt.fields.memStore
			r := NewinMemoryRepoRepository(store)
			store.GetUsers().Put(entity.ID(1), testUser)
			ID, _ := r.Create(tt.args.create)

			if err := r.Delete(ID); (err != nil) != tt.wantErr {
				t.Errorf("inMemoryRepo.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// Move to inmemory_test.go
func Test_findUserInMemstore(t *testing.T) {
	type args struct {
		store  *memstore.UserStore
		userID entity.ID
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
			got, err := findUserInMemstore(tt.args.store, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("findUserInMemstore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findUserInMemstore() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Move to inmemory_test.go
func Test_filterMessagesFromMemory(t *testing.T) {
	type args struct {
		list     []memstore.Message
		filterFn func(Message) bool
	}
	tests := []struct {
		name    string
		args    args
		want    *Message
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := filterMessagesFromMemory(tt.args.list, tt.args.filterFn)
			if (err != nil) != tt.wantErr {
				t.Errorf("filterMessagesFromMemory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterMessagesFromMemory() = %v, want %v", got, tt.want)
			}
		})
	}
}
