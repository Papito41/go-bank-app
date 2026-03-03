package sqlc

import (
	"database/sql"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(dbConn *sql.DB) *Store {
	return &Store{
		Queries: NewQueries(dbConn),
		db:      dbConn,
	}
}
