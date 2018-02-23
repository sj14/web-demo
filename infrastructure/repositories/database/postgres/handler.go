package postgres

import (
	"log"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

type PostgresStore struct {
	conn   *sqlx.DB
}

func NewPostgresStore(dbfileName string) *PostgresStore {
	db, err := sqlx.Open("postgres", "user=postgres password=example dbname=demo sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	postgresHandler := &PostgresStore{db}
	return postgresHandler
}

func (s *PostgresStore) CloseConn() {
	s.conn.Close()
}
