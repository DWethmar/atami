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

// TestDBConfig config to create new testing db
type TestDBConfig struct {
	DBHost       string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string
	DBDriverName string
	MigrationFiles string
	DBMigrationVersion uint
}

// NewTestDB create new testy db. Returns a cleanup function and error.
func NewTestDB(testDBConfig *TestDBConfig) (*sql.DB, error) {
	var db *sql.DB

	dc := postgresConnectionConfig{
		DBHost:     testDBConfig.DBHost,
		DBPort:     testDBConfig.DBPort,
		DBUser:     testDBConfig.DBUser,
		DBPassword: testDBConfig.DBPassword,
		DBName:     "postgres",
	}

	dataSource := GetPostgresDataSource(dc)

	db, err := Connect(testDBConfig.DBDriverName, dataSource)
	if err != nil {
		fmt.Printf("Could not connect to database with %s %s", testDBConfig.DBDriverName, dataSource)
		panic(err)
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	if err := CreateDatabase(db, testDBConfig.DBName, testDBConfig.DBUser); err != nil {
		return nil, err
	}
	db.Close()

	dataSource = GetPostgresDataSource(dc)
	db, err = Connect(testDBConfig.DBDriverName, dataSource)
	if err != nil {
		fmt.Printf("Could not connect to database with %v %v", testDBConfig.DBDriverName, dataSource)
		return nil, err
	}

	if err := RunMigrations(db, testDBConfig.DBName, testDBConfig.MigrationFiles, testDBConfig.DBMigrationVersion); err != nil {
		fmt.Printf("Error while running migrations")
		return nil, err
	}

	return db, nil
}

// WithTestDB runs test with test DB and remove DB after test.
func WithTestDB(t *testing.T, test func(db *sql.DB) error) error {
	if c := config.Load(); c.Valid() == nil && c.TestWithDB {
		rand.Seed(time.Now().UTC().UnixNano())
		if db, err := NewTestDB(&TestDBConfig{
			DBHost:     c.DBHost,
			DBPort:     c.DBPort,
			DBUser:     c.DBUser,
			DBPassword: c.DBPassword,
			DBName:     fmt.Sprintf("%v_%v_%d", c.DBName, strings.ToLower(t.Name()), rand.Int()),
			MigrationFiles: c.MigrationFiles,
			DBMigrationVersion: c.DBMigrationVersion,
		}); err == nil {
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
