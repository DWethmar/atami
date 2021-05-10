package database

import (
	"database/sql"
	"fmt"
	"strings"
)

// Connect creates new DB
func Connect(driverName string, dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	fmt.Printf("Successfully connected with %s db! \n", driverName)
	return db, nil
}

type postgresConnectionConfig struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

// GetPostgresDataSource return the connection info
func GetPostgresDataSource(config postgresConnectionConfig) string {
	cParts := []string{
		fmt.Sprintf("host=%s", config.DBHost),
		fmt.Sprintf("port=%s", config.DBPort),
		"sslmode=disable",
	}

	if config.DBUser != "" {
		cParts = append(cParts, fmt.Sprintf("user=%s", config.DBUser))
	}

	if config.DBPassword != "" {
		cParts = append(cParts, fmt.Sprintf("password=%s", config.DBPassword))
	}

	if config.DBName != "" {
		cParts = append(cParts, fmt.Sprintf("dbname=%s", config.DBName))
	}

	return strings.Join(cParts, " ")
}
