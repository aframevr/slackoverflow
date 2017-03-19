package sqlite3

import (
	"database/sql"
	"fmt"

	// Importing database driver
	_ "github.com/mattn/go-sqlite3"
)

// Client resource
type Client struct {
	path string
	DB   *sql.DB
}

var sqlite *Client

// Load SQLite3 client
func Load(path string) (*Client, error) {
	sqlite = &Client{
		path: path,
	}

	// Open the database
	var err error
	sqlite.DB, err = sql.Open("sqlite3", sqlite.path)
	if err != nil {
		return nil, err
	}

	return sqlite, nil
}

// VerifyTable creates new table if one does not exist
func (sqlite3 *Client) VerifyTable(table string) error {
	switch table {
	// Create SlackQuestion table if needed
	case "SlackQuestion":
		_, err := sqlite3.DB.Exec(SlackQuestion{}.GetSchema())
		if err != nil {
			return err
		}
		break

	// Create StackExchangeQuestion table if needed
	case "StackExchangeQuestion":
		_, err := sqlite3.DB.Exec(StackExchangeQuestion{}.GetSchema())
		if err != nil {
			return err
		}
		break

	// Create StackExchangeUser tablr if needed
	case "StackExchangeUser":
		_, err := sqlite3.DB.Exec(StackExchangeUser{}.GetSchema())
		if err != nil {
			return err
		}
		break
	default:
		return fmt.Errorf("Unknown table %s", table)
	}
	return nil
}
