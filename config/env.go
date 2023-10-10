package config

import (
	"os"

	"github.com/joho/godotenv"
)

// This function will Load the ENVIORNMENT VARIABLES from .env if GO_ENV variable is not set
func LoadENV() error {
	goEnv := os.Getenv("GO_ENV")

	if goEnv == "" || goEnv == "development" {
		err := godotenv.Load()
		if err != nil {
			return err
		}
	}

	return nil
}
