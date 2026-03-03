package main

import (
	"database/sql"
	"log"
	"my-bank-app/api"
	"my-bank-app/db/sqlc"

	_ "github.com/lib/pq"
)

func main() {
	// Connect to Postgres
	conn, err := sql.Open("postgres", "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable")
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := sqlc.NewStore(conn)
	server := api.NewServer(store)

	log.Println("Bank Server starting on port 8080...")
	err = server.Start("0.0.0.0:8080")
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
