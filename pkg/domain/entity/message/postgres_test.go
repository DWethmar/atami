package message

import (
	"database/sql"
	"fmt"
	"sync"
	"testing"

	"github.com/dwethmar/atami/pkg/database"
	"github.com/dwethmar/atami/pkg/domain/entity"
	"github.com/dwethmar/atami/pkg/domain/seed"
	"github.com/stretchr/testify/assert"
)

func seedDatabase(db *sql.DB, deps repoTestDependencies) error {
	fmt.Println("Seeding")
	for _, user := range deps.users {
		fmt.Println(user)
		if _, err := seed.SeedUser(
			db,
			user.UID,
			user.Username,
			"password",
			user.Username+"@test.nl",
			"biography",
			entity.Now(),
			entity.Now(),
		); err != nil {
			return err
		}
	}
	for _, message := range deps.messages {
		fmt.Println(message)
		if _, err := seed.SeedMessage(
			db,
			message.UID,
			message.Text,
			message.CreatedByUserID,
			message.CreatedAt,
			message.UpdatedAt,
		); err != nil {
			return err
		}
	}
	fmt.Println("stopped Seeding")
	return nil
}

func Test_PostgresRepo_Get(t *testing.T) {
	mux := &sync.Mutex{}
	dbs := []*sql.DB{}
	defer func() {
		for _, db := range dbs {
			db.Close()
		}
	}()
	deps := newRepoTestDependencies()
	testRepositoryGet(
		t,
		deps,
		func() Repository {
			db, err := database.NewTestDB(t)
			if err != nil {
				assert.FailNow(t, err.Error())
			}
			mux.Lock()
			dbs = append(dbs, db)
			mux.Unlock()

			if err := seedDatabase(db, deps); err != nil {
				fmt.Print(err)
				t.FailNow()
			}
			return NewPostgresRepository(db)
		},
	)
}

// func Test_PostgresRepo_GetByUID(t *testing.T) {
// 	deps := newRepoTestDependencies()
// 	testRepositoryGetByUID(
// 		t,
// 		deps,
// 		func() Repository {
// 			store := memstore.NewStore()
// 			for _, user := range deps.users {
// 				store.GetUsers().Put(user.ID, *userToMemoryMap(*user))
// 			}
// 			for _, message := range deps.messages {
// 				store.GetMessages().Put(message.ID, *messageToMemoryMap(*message))
// 			}
// 			return NewInMemoryRepo(store)
// 		},
// 	)
// }

// func Test_PostgresRepo_List(t *testing.T) {
// 	deps := newRepoTestDependencies()
// 	testRepositoryList(
// 		t,
// 		deps,
// 		func() Repository {
// 			store := memstore.NewStore()
// 			for _, user := range deps.users {
// 				store.GetUsers().Put(user.ID, *userToMemoryMap(*user))
// 			}
// 			for _, message := range deps.messages {
// 				store.GetMessages().Put(message.ID, *messageToMemoryMap(*message))
// 			}
// 			return NewInMemoryRepo(store)
// 		},
// 	)
// }

// func Test_PostgresRepo_Update(t *testing.T) {
// 	deps := newRepoTestDependencies()
// 	testRepositoryUpdate(
// 		t,
// 		deps,
// 		func() Repository {
// 			store := memstore.NewStore()
// 			for _, user := range deps.users {
// 				store.GetUsers().Put(user.ID, *userToMemoryMap(*user))
// 			}
// 			for _, message := range deps.messages {
// 				store.GetMessages().Put(message.ID, *messageToMemoryMap(*message))
// 			}
// 			return NewInMemoryRepo(store)
// 		},
// 	)
// }

// func Test_PostgresRepo_Create(t *testing.T) {
// 	deps := newRepoTestDependencies()
// 	testRepositoryCreate(
// 		t,
// 		deps,
// 		func() Repository {
// 			store := memstore.NewStore()
// 			for _, user := range deps.users {
// 				store.GetUsers().Put(user.ID, *userToMemoryMap(*user))
// 			}
// 			for _, message := range deps.messages {
// 				store.GetMessages().Put(message.ID, *messageToMemoryMap(*message))
// 			}
// 			return NewInMemoryRepo(store)
// 		},
// 	)
// }

// func Test_PostgresRepo_Delete(t *testing.T) {
// 	deps := newRepoTestDependencies()
// 	testRepositoryDelete(
// 		t,
// 		deps,
// 		func() Repository {
// 			store := memstore.NewStore()
// 			for _, user := range deps.users {
// 				store.GetUsers().Put(user.ID, *userToMemoryMap(*user))
// 			}
// 			for _, message := range deps.messages {
// 				store.GetMessages().Put(message.ID, *messageToMemoryMap(*message))
// 			}
// 			return NewInMemoryRepo(store)
// 		},
// 	)
// }
