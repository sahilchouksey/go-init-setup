package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/TVM-LLC/docapp_backend/model"
	_ "github.com/lib/pq"
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
	connectStr := "user=postgres password=lol dbname=postgres sslmode=disable"

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
