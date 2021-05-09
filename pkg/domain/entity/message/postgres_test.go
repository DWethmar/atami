package message

import (
	"database/sql"
	"testing"
	"time"

	"github.com/dwethmar/atami/pkg/database"
	"github.com/dwethmar/atami/pkg/domain/seed"
	"github.com/dwethmar/atami/pkg/memstore"
)

func seedPostgressRepo(db *sql.DB, deps repoTestDependencies) {
	for _, user := range deps.users {
		seed.SeedUser(db, user.UID, user.Username, user.Username+"@test.nl", "abc", time.Now(), time.Now())
	}
	for _, message := range deps.messages {
		seed.SeedMessage(db, message.UID, message.Text, message.CreatedByUserID, message.CreatedAt, message.UpdatedAt)
	}
}

func Test_PostgresRepo_Get(t *testing.T) {
	database.WithTestDB(t, func(db *sql.DB) error {
		deps := newRepoTestDependencies()
		testRepositoryGet(
			t,
			deps,
			func() Repository {
				seedPostgressRepo(db, deps)
				return NewPostgresRepository(db)
			},
		)
		return nil
	})
}

func Test_PostgresRepo_GetByUID(t *testing.T) {
	deps := newRepoTestDependencies()
	testRepositoryGetByUID(
		t,
		deps,
		func() Repository {
			store := memstore.NewStore()
			for _, user := range deps.users {
				store.GetUsers().Put(user.ID, *userToMemoryMap(*user))
			}
			for _, message := range deps.messages {
				store.GetMessages().Put(message.ID, *messageToMemoryMap(*message))
			}
			return NewInMemoryRepo(store)
		},
	)
}

func Test_PostgresRepo_List(t *testing.T) {
	deps := newRepoTestDependencies()
	testRepositoryList(
		t,
		deps,
		func() Repository {
			store := memstore.NewStore()
			for _, user := range deps.users {
				store.GetUsers().Put(user.ID, *userToMemoryMap(*user))
			}
			for _, message := range deps.messages {
				store.GetMessages().Put(message.ID, *messageToMemoryMap(*message))
			}
			return NewInMemoryRepo(store)
		},
	)
}

func Test_PostgresRepo_Update(t *testing.T) {
	deps := newRepoTestDependencies()
	testRepositoryUpdate(
		t,
		deps,
		func() Repository {
			store := memstore.NewStore()
			for _, user := range deps.users {
				store.GetUsers().Put(user.ID, *userToMemoryMap(*user))
			}
			for _, message := range deps.messages {
				store.GetMessages().Put(message.ID, *messageToMemoryMap(*message))
			}
			return NewInMemoryRepo(store)
		},
	)
}

func Test_PostgresRepo_Create(t *testing.T) {
	deps := newRepoTestDependencies()
	testRepositoryCreate(
		t,
		deps,
		func() Repository {
			store := memstore.NewStore()
			for _, user := range deps.users {
				store.GetUsers().Put(user.ID, *userToMemoryMap(*user))
			}
			for _, message := range deps.messages {
				store.GetMessages().Put(message.ID, *messageToMemoryMap(*message))
			}
			return NewInMemoryRepo(store)
		},
	)
}

func Test_PostgresRepo_Delete(t *testing.T) {
	deps := newRepoTestDependencies()
	testRepositoryDelete(
		t,
		deps,
		func() Repository {
			store := memstore.NewStore()
			for _, user := range deps.users {
				store.GetUsers().Put(user.ID, *userToMemoryMap(*user))
			}
			for _, message := range deps.messages {
				store.GetMessages().Put(message.ID, *messageToMemoryMap(*message))
			}
			return NewInMemoryRepo(store)
		},
	)
}
