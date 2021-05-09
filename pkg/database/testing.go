package database

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/dwethmar/atami/pkg/config"
	"github.com/stretchr/testify/assert"
)

// CreateDatabase and give privileges to user.
func CreateDatabase(db *sql.DB, database string, user string) error {
	fmt.Printf("Creating database: %v \n", database)

	if _, err := db.Exec("CREATE DATABASE " + database + ";"); err != nil {
		return err
	}

	fmt.Printf("Granting privileges to %v for %v \n", user, database)

	if _, err := db.Exec(fmt.Sprintf("grant all privileges on database %s to %s", database, user)); err != nil {
		return err
	}

	fmt.Printf("Created database: %v \n", database)

	return nil
}

// DropDatabase drops database
func DropDatabase(db *sql.DB, database string) error {
	fmt.Printf("Dropping database: %v \n", database)
	if _, err := db.Exec("DROP DATABASE " + database); err != nil {
		return err
	}
	return nil
}

// ExecSQLFile runs a sql file
func ExecSQLFile(db *sql.DB, sqlFile string) error {
	fmt.Printf("Reading SQL file: %v \n", sqlFile)

	b, err := ioutil.ReadFile(sqlFile)
	if err != nil {
		log.Fatal(err)
	}
	return ExecSQL(db, b)
}

// ExecSQL runs a sql query
func ExecSQL(db *sql.DB, sql []byte) error {
	fmt.Printf("Excecuting SQL file: %v \n", string(sql))
	if _, err := db.Exec(string(sql)); err != nil {
		return err
	}
	return nil
}

// NewTestDB create new testy db. Returns a cleanup function and error.
func NewTestDB(c config.Config) (*sql.DB, error) {
	var db *sql.DB

	if err := c.Valid(); err != nil {
		return nil, err
	}

	dataSource := GetPostgresConnectionString(config.Config{
		DBHost:       c.DBHost,
		DBPort:       c.DBPort,
		DBUser:       c.DBUser,
		DBPassword:   c.DBPassword,
		DBName:       "postgres",
		DBDriverName: c.DBDriverName,
	})

	db, err := Connect(c.DBDriverName, dataSource)
	if err != nil {
		fmt.Printf("Could not connect to database with %v %v", c.DBDriverName, dataSource)
		panic(err)
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	if err := CreateDatabase(db, c.DBName, c.DBUser); err != nil {
		return nil, err
	}
	db.Close()

	dataSource = GetPostgresConnectionString(c)
	db, err = Connect(c.DBDriverName, dataSource)
	if err != nil {
		fmt.Printf("Could not connect to database with %v %v", c.DBDriverName, dataSource)
		return nil, err
	}

	if err := RunMigrations(db, c.DBName, c.MigrationFiles, c.DBMigrationVersion); err != nil {
		fmt.Printf("Error while running migrations")
		return nil, err
	}

	// if err := ExecSQLFile(db, c.TestSeedSQLFile); err != nil {
	// 	fmt.Printf("Error while running migrations")
	// 	return nil, err
	// }

	return db, nil
}

// WithTestDB runs test with test DB and remove DB after test.
func WithTestDB(t *testing.T, test func(db *sql.DB) error) error {
	if c := config.Load(); c.Valid() == nil && c.TestWithDB {
		rand.Seed(time.Now().UTC().UnixNano())
		c.DBName = fmt.Sprintf("%v_%v_%d", c.DBName, strings.ToLower(t.Name()), rand.Int())

		if db, err := NewTestDB(c); err == nil {
			defer db.Close()
			if assert.NoError(t, err) {
				return test(db)
			}
		} else {
			return err
		}
	} else {
		t.Skip("Skip test")
		return nil
	}

	return nil
}
