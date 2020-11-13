package database

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/dwethmar/atami/pkg/config"
)

// DB model
type DB struct {
	*sql.DB
}

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

// GetPostgresConnectionString return the connection info
func GetPostgresConnectionString(env config.Config) string {
	cParts := []string{
		fmt.Sprintf("host=%s", env.DBHost),
		fmt.Sprintf("port=%s", env.DBPort),
		"sslmode=disable",
	}

	if env.DBUser != "" {
		cParts = append(cParts, fmt.Sprintf("user=%s", env.DBUser))
	}

	if env.DBPassword != "" {
		cParts = append(cParts, fmt.Sprintf("password=%s", env.DBPassword))
	}

	if env.DBName != "" {
		cParts = append(cParts, fmt.Sprintf("dbname=%s", env.DBName))
	}

	return strings.Join(cParts, " ")
}
