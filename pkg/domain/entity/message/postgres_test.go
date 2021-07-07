package message

import (
	"database/sql"
	"fmt"
	"sync"
	"testing"

	"github.com/dwethmar/atami/pkg/config"
	"github.com/dwethmar/atami/pkg/database"
	"github.com/dwethmar/atami/pkg/domain/entity"
	"github.com/dwethmar/atami/pkg/domain/seed"
	"github.com/stretchr/testify/assert"
)

func seedDatabase(db *sql.DB, deps *repoTestDependencies) error {
	return database.WithTransaction(db, func(t database.Transaction) error {
		for _, user := range deps.users {
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
		return nil
	})
}

func Test_PostgresRepo_Get(t *testing.T) {
	if c := config.Load(); !c.TestWithDB {
		t.Skip("Skip test")
	}
	mux := &sync.Mutex{}
	dbs := []*sql.DB{}
	defer func() {
		for _, db := range dbs {
			db.Close()
		}
	}()
	deps := newRepoTestDependencies()
	testRepository_Get(
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

func Test_PostgresRepo_GetByUID(t *testing.T) {
	if c := config.Load(); !c.TestWithDB {
		t.Skip("Skip test")
	}
	mux := &sync.Mutex{}
	dbs := []*sql.DB{}
	defer func() {
		for _, db := range dbs {
			db.Close()
		}
	}()
	deps := newRepoTestDependencies()
	testRepository_GetByUID(
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

func Test_PostgresRepo_List(t *testing.T) {
	if c := config.Load(); !c.TestWithDB {
		t.Skip("Skip test")
	}
	mux := &sync.Mutex{}
	dbs := []*sql.DB{}
	defer func() {
		for _, db := range dbs {
			db.Close()
		}
	}()
	deps := newRepoTestDependencies()
	testRepository_List(
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

func Test_PostgresRepo_Update(t *testing.T) {
	if c := config.Load(); !c.TestWithDB {
		t.Skip("Skip test")
	}
	mux := &sync.Mutex{}
	dbs := []*sql.DB{}
	defer func() {
		for _, db := range dbs {
			db.Close()
		}
	}()
	deps := newRepoTestDependencies()
	testRepository_Update(
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

func Test_PostgresRepo_Create(t *testing.T) {
	if c := config.Load(); !c.TestWithDB {
		t.Skip("Skip test")
	}
	mux := &sync.Mutex{}
	dbs := []*sql.DB{}
	defer func() {
		for _, db := range dbs {
			db.Close()
		}
	}()
	deps := newRepoTestDependencies()
	testRepository_Create(
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

func Test_PostgresRepo_Delete(t *testing.T) {
	if c := config.Load(); !c.TestWithDB {
		t.Skip("Skip test")
	}
	mux := &sync.Mutex{}
	dbs := []*sql.DB{}
	defer func() {
		for _, db := range dbs {
			db.Close()
		}
	}()
	deps := newRepoTestDependencies()
	testRepository_Delete(
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
