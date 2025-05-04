package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/sahilchouksey/go-init-setup/config"
	"github.com/sahilchouksey/go-init-setup/model"
)

type Storage interface {
	Init() error
	Close()
	// TODO: Add methods for interacting with the database
	GetTodos() ([]model.Todo, error)
}

type PostgreSQLStore struct {
	db *sql.DB
}

func Start() (*PostgreSQLStore, error) {
	getEnv, err := config.Get()

	if err != nil {
		return nil, err
	}

	// connectStr := fmt.Sprintf("user=postgres password=lol dbname=postgres sslmode=disable", )
	connectStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", getEnv.DB_USER_NAME, getEnv.DB_PASSWORD, getEnv.DB_NAME, getEnv.DB_SSL_MODE)

	db, err := sql.Open("postgres", connectStr)
	if err != nil {
		fmt.Println("Unable to Start PostgresSQL Databse.")
		return nil, err
	}

	log.Println("Successfully connected to PostgresSQL Database.")
	return &PostgreSQLStore{
		db: db,
	}, nil
}

func (s *PostgreSQLStore) Init() error {
	log.Println("Initializing PostgresSQL Database.")
	err := s.Initialize()
	return err
}

func (s *PostgreSQLStore) Close() {
	log.Println("Closing PostgresSQL Database.")
	s.db.Close()
}
