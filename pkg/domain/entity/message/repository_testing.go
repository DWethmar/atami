package message

import (
	"reflect"
	"testing"
	"time"

	"github.com/dwethmar/atami/pkg/domain/entity"
	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/stretchr/testify/assert"
)

type repoTestDependencies struct {
	users    []*User
	messages []*Message
}

func newRepoTestDependencies() *repoTestDependencies {
	users := []*User{
		mapUserFromMemory(memstore.User{
			ID:        1,
			UID:       entity.NewUID(),
			Username:  "user1",
			Email:     "user1@test.nl",
			Password:  "hashedpassword",
			Biography: "bio",
			CreatedAt: entity.Now(),
			UpdatedAt: entity.Now(),
		}),
		mapUserFromMemory(memstore.User{
			ID:        2,
			UID:       entity.NewUID(),
			Username:  "user2",
			Email:     "user2@test.nl",
			Password:  "hashedpassword",
			Biography: "bio",
			CreatedAt: entity.Now(),
			UpdatedAt: entity.Now(),
		}),	
	}
	messages := []*Message{
		{
			ID:              entity.ID(1),
			UID:             entity.NewUID(),
			Text:            "message text 1",
			CreatedByUserID: users[0].ID,
			CreatedBy:       *users[0],
			CreatedAt:       entity.Now().Add(time.Duration(1000)),
			UpdatedAt:       entity.Now().Add(time.Duration(1000)),
		},
		{
			ID:              entity.ID(2),
			UID:             entity.NewUID(),
			Text:            "message text 2",
			CreatedByUserID: users[1].ID,
			CreatedBy:       *users[1],
			CreatedAt:       entity.Now().Add(time.Duration(2000)),
			UpdatedAt:       entity.Now().Add(time.Duration(2000)),
		},
		{
			ID:              entity.ID(3),
			UID:             entity.NewUID(),
			Text:            "message text 3",
			CreatedByUserID: users[1].ID,
			CreatedBy:       *users[1],
			CreatedAt:       entity.Now().Add(time.Duration(3000)),
			UpdatedAt:       entity.Now().Add(time.Duration(3000)),
		},
		{
			ID:              entity.ID(4),
			UID:             entity.NewUID(),
			Text:            "message text 4",
			CreatedByUserID: users[1].ID,
			CreatedBy:       *users[1],
			CreatedAt:       entity.Now().Add(time.Duration(4000)),
			UpdatedAt:       entity.Now().Add(time.Duration(4000)),
		},
		{
			ID:              entity.ID(5),
			UID:             entity.NewUID(),
			Text:            "message text 5",
			CreatedByUserID: users[1].ID,
			CreatedBy:       *users[1],
			CreatedAt:       entity.Now().Add(time.Duration(5000)),
			UpdatedAt:       entity.Now().Add(time.Duration(5000)),
		},
	}

	return &repoTestDependencies{
		users:    users,
		messages: messages,
	}
}

type setupRepository = func() Repository

func testRepository_Get(t *testing.T, dependencies *repoTestDependencies, setup setupRepository) {
	testMessage := dependencies.messages[0]

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
		want    *Message
		wantErr bool
	}{
		{
			name: "Successfully get message by ID",
			fields: fields{
				repo: setup(),
			},
			args: args{
				ID: testMessage.ID,
			},
			want:    testMessage,
			wantErr: false,
		},
		{
			name: "Fail get message by unknown UID",
			fields: fields{
				repo: setup(),
			},
			args: args{
				ID: 999,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := tt.fields.repo
			got, err := repo.Get(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !assert.Equal(t, tt.want, got) {
				t.Errorf("Repository.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func testRepository_GetByUID(t *testing.T, dependencies *repoTestDependencies, setup setupRepository) {
	testMessage := dependencies.messages[0]

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
		want    *Message
		wantErr bool
	}{
		{
			name: "Successfully get message by UID",
			fields: fields{
				repo: setup(),
			},
			args: args{
				UID: testMessage.UID,
			},
			want:    testMessage,
			wantErr: false,
		},
		{
			name: "Fail get message by unknown UID",
			fields: fields{
				repo: setup(),
			},
			args: args{
				UID: "abc123",
			},
			want:    nil,
			wantErr: true,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := tt.fields.repo
			got, err := repo.GetByUID(tt.args.UID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !assert.Equal(t, tt.want, got) {
				t.Errorf("Repository.GetByUID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func testRepository_List(t *testing.T, dependencies *repoTestDependencies, setup setupRepository) {
	testMessages := dependencies.messages

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
		want    []*Message
		wantErr bool
	}{
		{
			name: "Successfully get messages",
			fields: fields{
				repo: setup(),
			},
			args: args{
				limit:  10,
				offset: 0,
			},
			want: []*Message{
				testMessages[len(testMessages)-1],
				testMessages[len(testMessages)-2],
				testMessages[len(testMessages)-3],
				testMessages[len(testMessages)-4],
				testMessages[len(testMessages)-5],
			},
			wantErr: false,
		},
		{
			name: "Successfully get no messages",
			fields: fields{
				repo: setup(),
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
				repo: setup(),
			},
			args: args{
				limit:  4,
				offset: 0,
			},
			want: []*Message{
				testMessages[len(testMessages)-1],
				testMessages[len(testMessages)-2],
				testMessages[len(testMessages)-3],
				testMessages[len(testMessages)-4],
			},
			wantErr: false,
		},
		{
			name: "Successfully get paged messages with offset",
			fields: fields{
				repo: setup(),
			},
			args: args{
				limit:  3,
				offset: 1,
			},
			want: []*Message{
				testMessages[len(testMessages)-2],
				testMessages[len(testMessages)-3],
				testMessages[len(testMessages)-4],
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

func testRepository_Update(t *testing.T, dependencies *repoTestDependencies, setup setupRepository) {
	testMessage := dependencies.messages[0]

	type fields struct {
		repo Repository
	}
	type args struct {
		ID     entity.ID
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
				repo: setup(),
			},
			args: args{
				ID: testMessage.ID,
				update: Update{
					Text:      "updated text",
					UpdatedAt: entity.Now(),
				},
			},
			wantErr: false,
		},
		{
			name: "fail on nonexisting message",
			fields: fields{
				repo: setup(),
			},
			args: args{
				ID: entity.ID(9999999),
				update: Update{
					Text:      "updated text",
					UpdatedAt: entity.Now(),
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := tt.fields.repo

			expectedMsg := *testMessage
			expectedMsg.ID = tt.args.ID
			expectedMsg.Apply(tt.args.update)

			err := repo.Update(&expectedMsg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Check if the update really was successful
			if !tt.wantErr {
				updatedMsg, err := repo.Get(expectedMsg.ID)
				if assert.NoError(t, err) {
					return
				}
				if !reflect.DeepEqual(updatedMsg, expectedMsg) {
					t.Errorf("Repository.Create() = %v, want %v", updatedMsg, expectedMsg)
				}
			}
		})
	}
}

func testRepository_Create(t *testing.T, dependencies *repoTestDependencies, setup setupRepository) {
	createdAt := time.Now().UTC()
	testUser := dependencies.users[0]

	type fields struct {
		repo Repository
	}
	type args struct {
		message *Message
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
				repo: setup(),
			},
			args: args{
				message: &Message{
					UID:             "abc123",
					Text:            "updated text",
					CreatedByUserID: testUser.ID,
					CreatedBy: User{
						ID:       testUser.ID,
						UID:      testUser.UID,
						Username: testUser.Username,
					},
					CreatedAt: createdAt,
				},
			},
			want:    dependencies.messages[len(dependencies.messages)-1].ID + 1,
			wantErr: false,
		},
		{
			name: "Fail on unknown created by user",
			fields: fields{
				repo: setup(),
			},
			args: args{
				message: &Message{
					UID:             "abc123",
					Text:            "updated text",
					CreatedByUserID: entity.ID(9999),
					CreatedBy: User{
						ID:       entity.ID(9999),
						UID:      testUser.UID,
						Username: testUser.Username,
					},
					CreatedAt: createdAt,
				},
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := tt.fields.repo
			ID, err := repo.Create(tt.args.message)

			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(ID, tt.want) {
				t.Errorf("Repository.Create() = %v, want %v", ID, tt.want)
			}
		})
	}
}

func testRepository_Delete(t *testing.T, dependencies *repoTestDependencies, setup setupRepository) {
	testMessage := dependencies.messages[0]

	type fields struct {
		repo Repository
	}
	type args struct {
		messageID entity.ID
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
				repo: setup(),
			},
			args: args{
				messageID: testMessage.ID,
			},
			wantErr: false,
		},
		{
			name: "Fail on message not found",
			fields: fields{
				repo: setup(),
			},
			args: args{
				messageID: entity.ID(999),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := tt.fields.repo
			err := repo.Delete(tt.args.messageID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
