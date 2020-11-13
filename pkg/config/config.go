package config

import (
	"errors"
	"os"
)

// Config collection of the config variables.
type Config struct {
	DBHost             string
	DBPort             string
	DBUser             string
	DBPassword         string
	DBName             string
	DBDriverName       string
	MigrationFiles     string
	DBMigrationVersion uint
	AccessSecret       string
	TestWithDB         bool
	TestSeedSQLFile    string
}

// Valid checks if the config has valid values.
func (c Config) Valid() error {
	if c.DBDriverName == "" {
		return errors.New("Config is missing database connection information: DBDriverName")
	}

	if c.DBHost == "" {
		return errors.New("Config is missing database connection information: DBHost")
	}

	if c.DBName == "" {
		return errors.New("Config is missing database connection information: DBName")
	}

	if c.DBPassword == "" {
		return errors.New("Config is missing database connection information: DBPassword")
	}

	if c.DBPort == "" {
		return errors.New("Config is missing database connection information: DBPort")
	}

	if c.DBUser == "" {
		return errors.New("Config is missing database connection information: DBUser")
	}

	if c.MigrationFiles == "" {
		return errors.New("Config is missing database connection information: MigrationFiles")
	}

	if c.DBMigrationVersion == 0 {
		return errors.New("Config is missing database connection information: DBMigrationVersion")
	}

	return nil
}

// Load collects the necessary env vars and returns them in a struct.
func Load() Config {
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")

	dbName := "atami"
	if v, success := os.LookupEnv("POSTGRES_DATABASE"); success {
		dbName = v
	}

	dbDriverName := os.Getenv("DB_DRIVER_NAME")

	migrationFiles := "/app/migrations"
	if v, success := os.LookupEnv("MIGRATION_FILES"); success {
		migrationFiles = v
	}

	testWithDB := false
	if v, success := os.LookupEnv("TEST_WITH_DB"); success {
		testWithDB = success && v == "true"
	}

	testSeedFile := os.Getenv("TEST_SEED_FILE")
	dbMigrationVersion := uint(1)
	accessSecret := os.Getenv("ACCESS_SECRET")

	return Config{
		DBHost:             dbHost,
		DBPort:             dbPort,
		DBUser:             dbUser,
		DBPassword:         dbPassword,
		DBName:             dbName,
		DBDriverName:       dbDriverName,
		MigrationFiles:     migrationFiles,
		DBMigrationVersion: dbMigrationVersion,
		AccessSecret:       accessSecret,
		TestWithDB:         testWithDB,
		TestSeedSQLFile:    testSeedFile,
	}
}
