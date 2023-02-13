package database

import (
	"database/sql"

	"fmt"

	// Required to initialize postgres connection
	_ "github.com/lib/pq"
)

var client *Client


// Client ...
type Client struct {
	client *sql.DB
}


// MakeDb ...
func MakeDb() (*Client, error) {
	if client != nil {
		return client, nil
	}

	host := "localhost"
	port := "9945"
	databaseName := "rent-cat-db"
	username := "rent-cat-user"
	password := "rent-cat-pass"

	connStr := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", 
		host, port, databaseName, username, password,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("MakeDB Error: %w", err)
	}

	client = &Client{client: db}
	return client, nil
}

