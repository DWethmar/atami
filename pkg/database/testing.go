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

func NewTestDBName(t *testing.T) string {
	rand.Seed(time.Now().UTC().UnixNano())
	return fmt.Sprintf("%v_%d", strings.ToLower(t.Name()), rand.Int())
}

// CreateDatabase and give privileges to user.
func CreateDatabase(db *sql.DB, name string, user string) error {
	fmt.Printf("Creating database: %v \n", name)

	if _, err := db.Exec("CREATE DATABASE " + name + ";"); err != nil {
		return err
	}

	fmt.Printf("Granting privileges to %v for %v \n", user, name)

	if _, err := db.Exec(fmt.Sprintf("grant all privileges on database %s to %s", name, user)); err != nil {
		return err
	}

	fmt.Printf("Created database: %v \n", name)

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

// Create new testing Db
func NewTestDB(t *testing.T) (*sql.DB, error) {
	c := config.Load()
	if err := c.Valid(); err != nil {
		return nil, err
	}
	dbConfig := &PostgresConnectionConfig{}
	dbConfig.Load(c)

	db, err := CreateTestDB(
		dbConfig, 
		c.DBDriverName, 
		NewTestDBName(t), 
		c.MigrationFiles, 
		c.DBMigrationVersion,
	)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// NewTestDB create new testy db. Returns a cleanup function and error.
func CreateTestDB(
	config *PostgresConnectionConfig, 
	driverName, 
	databaseName,
	migrationFiles string, 
	DBMigrationVersion uint,
) (*sql.DB, error) {
	var db *sql.DB

	dataSource := GetPostgresDataSource(config)

	db, err := Connect(driverName, dataSource)
	if err != nil {
		fmt.Printf("Could not connect to database with driver: %s datasource: %s", driverName, dataSource)
		panic(err)
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	if err := CreateDatabase(db, databaseName, config.DBUser); err != nil {
		return nil, err
	}
	db.Close()

	newDBConfig := *config
	newDBConfig.DBName = databaseName
	dataSource = GetPostgresDataSource(&newDBConfig)

	db, err = Connect(driverName, dataSource)
	if err != nil {
		fmt.Printf("Could not connect to database with driver: %s datasource: %s", driverName, dataSource)
		return nil, err
	}

	if err := RunMigrations(db, config.DBName, migrationFiles, DBMigrationVersion); err != nil {
		fmt.Printf("Error while running migrations")
		return nil, err
	}

	return db, nil
}

// WithTestDB runs test with test DB and remove DB after test.
func WithTestDB(t *testing.T, test func(db *sql.DB) error) error {
	if c := config.Load(); c.Valid() == nil && c.TestWithDB {
		dbConfig := &PostgresConnectionConfig{}
		dbConfig.Load(c)

		if db, err := CreateTestDB(
			dbConfig, 
			c.DBDriverName, 
			NewTestDBName(t),
			c.MigrationFiles, 
			c.DBMigrationVersion,
		); err == nil {
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
