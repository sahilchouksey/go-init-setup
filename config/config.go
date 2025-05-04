package config

import (
	"os"
	"strconv"

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

type EnviornmentVariable struct {
	// All variables
	GO_ENV       string
	DB_USER_NAME string
	DB_PASSWORD  string
	DB_NAME      string
	DB_SSL_MODE  string
	PORT         int
}

func Get() (*EnviornmentVariable, error) {

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8080
	}

	envVariables := &EnviornmentVariable{
		GO_ENV:       os.Getenv("GO_ENV"),
		DB_USER_NAME: os.Getenv("DB_USER_NAME"),
		DB_PASSWORD:  os.Getenv("DB_PASSWORD"),
		DB_NAME:      os.Getenv("DB_NAME"),
		DB_SSL_MODE:  os.Getenv("DB_SSL_MODE"),
		PORT:         port,
	}

	return envVariables, nil
}
