package config

import (
	"fmt"
	"os"
)

const (
	dbHostEnvKey          = "POSTGRES_HOST"
	dbPortEnvKey          = "POSTGRES_PORT"
	dbUserEnvKey          = "POSTGRES_USER"
	dbPasswordEnvKey      = "POSTGRES_PASSWORD"
	dbNameEnvKey          = "POSTGRES_DATABASE"
	dbDriverNameEnvKey    = "DB_DRIVER_NAME"
	migrationFilesEnvKey  = "MIGRATION_FILES"
	accessSecretEnvKey    = "ACCESS_SECRET"
	testWithDBEnvKey      = "TEST_WITH_DB"
	testSeedSQLFileEnvKey = "TEST_SEED_FILE"
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
		return fmt.Errorf("env is missing database connection information: %s", dbDriverNameEnvKey)
	}

	if c.DBHost == "" {
		return fmt.Errorf("env is missing database connection information: %s", dbHostEnvKey)
	}

	if c.DBName == "" {
		return fmt.Errorf("env is missing database connection information: %s", dbNameEnvKey)
	}

	if c.DBPassword == "" {
		return fmt.Errorf("env is missing database connection information: %s", dbPasswordEnvKey)
	}

	if c.DBPort == "" {
		return fmt.Errorf("env is missing database connection information: %s", dbPortEnvKey)
	}

	if c.DBUser == "" {
		return fmt.Errorf("env is missing database connection information: %s", dbUserEnvKey)
	}

	if c.MigrationFiles == "" {
		return fmt.Errorf("env is missing database connection information: %s", migrationFilesEnvKey)
	}

	if c.AccessSecret == "" {
		return fmt.Errorf("env is missing access secret: %s", accessSecretEnvKey)
	}

	return nil
}

// Load collects the necessary env vars and returns them in a struct.
func Load() Config {
	dbHost := os.Getenv(dbHostEnvKey)
	dbPort := os.Getenv(dbPortEnvKey)
	dbUser := os.Getenv(dbUserEnvKey)
	dbPassword := os.Getenv(dbPasswordEnvKey)

	dbName := ""
	if v, success := os.LookupEnv(dbNameEnvKey); success {
		dbName = v
	}

	dbDriverName := os.Getenv(dbDriverNameEnvKey)

	migrationFiles := ""
	if v, success := os.LookupEnv(migrationFilesEnvKey); success {
		migrationFiles = v
	}

	testWithDB := false
	if v, success := os.LookupEnv(testWithDBEnvKey); success {
		testWithDB = success && v == "true"
	}

	testSeedFile := os.Getenv(testSeedSQLFileEnvKey)
	dbMigrationVersion := uint(1)
	accessSecret := os.Getenv(accessSecretEnvKey)

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
