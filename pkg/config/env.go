package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

var defaultEnvPath = ".env"

// LoadEnvFile by file path
func LoadEnvFile(path ...string) *Config {
	if path == nil {
		path = []string{defaultEnvPath}
	}
	err := godotenv.Load(path...)
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	return Load()
}
