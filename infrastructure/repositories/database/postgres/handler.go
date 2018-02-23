package postgres

import (
	"log"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattes/migrate/source/file"

	"os"
)

type PostgresStore struct {
	conn   *sqlx.DB
}

func NewPostgresStore() *PostgresStore {
	conn, err := sqlx.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	err = conn.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}


	postgresHandler := &PostgresStore{conn}
	return postgresHandler
}



func (s *PostgresStore) CloseConn() {
	s.conn.Close()
}
