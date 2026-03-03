package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func ConnectTest() {
	// 1. The Connection String (The Address of the Vault)
	connStr := "postgres://username:password@localhost:5432/my_bank"

	// 2. Connect to the Database
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	// 3. The "Ping" (Testing the connection)
	err = conn.Ping(context.Background())
	if err != nil {
		fmt.Println("Database is sleeping... couldn't ping!")
	} else {
		fmt.Println("Oshey! Connection successful!")
	}
}
