package db

import (
	"database/sql"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(dbConn *sql.DB) *Store {
	return &Store{
		Queries: New(dbConn),
		db:      dbConn,
	}
}
