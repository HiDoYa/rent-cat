package database

import (
	"database/sql"

	"fmt"

	// Required to initialize postgres connection
	_ "github.com/lib/pq"
)

var client *Client


// Client holds a connection to the database
type Client struct {
	client *sql.DB
}


// MakeDb is a singleton that initializes a client object to database
func MakeDb() (*Client, error) {
	if client != nil {
		return client, nil
	}

	// TODO Populate with prod credentials
	host := "localhost"
	port := "9945"
	database := "rent-cat-db"
	username := "rent-cat-user"
	password := "rent-cat-pass"

	connStr := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", 
		host, port, database, username, password,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("MakeDB Error: %w", err)
	}

	client = &Client{client: db}
	return client, nil
}

