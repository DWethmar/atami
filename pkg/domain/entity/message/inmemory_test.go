package message

import (
	"reflect"
	"testing"

	"github.com/dwethmar/atami/pkg/domain/entity"
	"github.com/dwethmar/atami/pkg/memstore"
)

func Test_inMemoryRepo_Get(t *testing.T) {
	deps := newTestFixtures()
	testRepository_Get(
		t,
		deps,
		func() Repository {
			store := memstore.NewStore()
			for _, user := range deps.users {
				store.GetUsers().Put(user.ID, *userToMemoryMap(*user))
			}
			for _, message := range deps.messages {
				store.GetMessages().Put(message.ID, *messageToMemoryMap(message))
			}
			return NewInMemoryRepo(store)
		},
	)
}

func Test_inMemoryRepo_GetByUID(t *testing.T) {
	deps := newTestFixtures()
	testRepository_GetByUID(
		t,
		deps,
		func() Repository {
			store := memstore.NewStore()
			for _, user := range deps.users {
				store.GetUsers().Put(user.ID, *userToMemoryMap(*user))
			}
			for _, message := range deps.messages {
				store.GetMessages().Put(message.ID, *messageToMemoryMap(message))
			}
			return NewInMemoryRepo(store)
		},
	)
}

func Test_inMemoryRepo_List(t *testing.T) {
	deps := newTestFixtures()
	testRepository_List(
		t,
		deps,
		func() Repository {
			store := memstore.NewStore()
			for _, user := range deps.users {
				store.GetUsers().Put(user.ID, *userToMemoryMap(*user))
			}
			for _, message := range deps.messages {
				store.GetMessages().Put(message.ID, *messageToMemoryMap(message))
			}
			return NewInMemoryRepo(store)
		},
	)
}

func Test_inMemoryRepo_Update(t *testing.T) {
	deps := newTestFixtures()
	testRepository_Update(
		t,
		deps,
		func() Repository {
			store := memstore.NewStore()
			for _, user := range deps.users {
				store.GetUsers().Put(user.ID, *userToMemoryMap(*user))
			}
			for _, message := range deps.messages {
				store.GetMessages().Put(message.ID, *messageToMemoryMap(message))
			}
			return NewInMemoryRepo(store)
		},
	)
}

func Test_inMemoryRepo_Create(t *testing.T) {
	deps := newTestFixtures()
	testRepository_Create(
		t,
		deps,
		func() Repository {
			store := memstore.NewStore()
			for _, user := range deps.users {
				store.GetUsers().Put(user.ID, *userToMemoryMap(*user))
			}
			for _, message := range deps.messages {
				store.GetMessages().Put(message.ID, *messageToMemoryMap(message))
			}
			return NewInMemoryRepo(store)
		},
	)
}

func Test_inMemoryRepo_Delete(t *testing.T) {
	deps := newTestFixtures()
	testRepository_Delete(
		t,
		deps,
		func() Repository {
			store := memstore.NewStore()
			for _, user := range deps.users {
				store.GetUsers().Put(user.ID, *userToMemoryMap(*user))
			}
			for _, message := range deps.messages {
				store.GetMessages().Put(message.ID, *messageToMemoryMap(message))
			}
			return NewInMemoryRepo(store)
		},
	)
}

// Move to inmemory_test.go
func Test_findUserInMemstore(t *testing.T) {
	userStore := memstore.NewStore().GetUsers()
	testUser := memstore.NewUserFixture(entity.ID(1))
	userStore.Put(testUser.ID, *testUser)

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
		{
			name: "Successfully find message in memstore",
			args: args{
				store:  userStore,
				userID: entity.ID(1),
			},
			want: &User{
				ID:       testUser.ID,
				UID:      testUser.UID,
				Username: testUser.Username,
			},
			wantErr: false,
		},
		{
			name: "Error on message not found",
			args: args{
				store:  userStore,
				userID: entity.ID(2),
			},
			want:    nil,
			wantErr: true,
		},
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
	testUser := memstore.NewUserFixture(entity.ID(1))
	msg1 := memstore.NewMessageFixture(entity.ID(1), testUser.ID)
	msg2 := memstore.NewMessageFixture(entity.ID(2), testUser.ID)
	msg3 := memstore.NewMessageFixture(entity.ID(3), testUser.ID)

	type args struct {
		list     []memstore.Message
		filterFn func(*Message) bool
	}

	tests := []struct {
		name    string
		args    args
		want    *Message
		wantErr bool
	}{
		{
			name: "Successfully filter messages",
			args: args{
				list: []memstore.Message{
					*msg1, *msg2, *msg3,
				},
				filterFn: func(msg *Message) bool {
					return msg.ID == 2
				},
			},
			want: &Message{
				ID:              msg2.ID,
				UID:             msg2.UID,
				Text:            msg2.Text,
				CreatedByUserID: msg2.CreatedByUserID,
				CreatedAt:       msg2.CreatedAt,
			},
			wantErr: false,
		},
		{
			name: "Error on message not found",
			args: args{
				list: []memstore.Message{
					*msg1, *msg2, *msg3,
				},
				filterFn: func(msg *Message) bool {
					return msg.ID == 4
				},
			},
			want:    nil,
			wantErr: true,
		},
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
