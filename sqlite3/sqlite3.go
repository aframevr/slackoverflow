package sqlite3

import (
	"database/sql"

	// Importing database driver
	_ "github.com/mattn/go-sqlite3"
)

// Client resource
type Client struct {
	path    string
	DB      *sql.DB
	Schemas SlackOverflowSchemas
}

// Load SQLite3 client
func Load(path string) (*Client, error) {
	sqlite := &Client{
		path: path,
	}

	// Open the database
	var err error
	sqlite.DB, err = sql.Open("sqlite3", sqlite.path)
	if err != nil {
		return nil, err
	}

	// Load schemas
	sqlite.Schemas.loadSchemas()

	return sqlite, nil
}

// VerifyTable creates new table if one does not exist
func (sqlite3 *Client) VerifyTable(table string) error {
	switch table {

	// Create StackQuestion table if needed
	case "StackQuestion":
		_, err := sqlite3.DB.Exec(sqlite3.Schemas.StackQuestion)
		if err != nil {
			return err
		}
		break
	// Create StackQuestionLink table if needed
	case "StackQuestionLink":
		_, err := sqlite3.DB.Exec(sqlite3.Schemas.StackQuestionLink)
		if err != nil {
			return err
		}
		break
	// Create StackUser tablr if needed
	case "StackUser":
		_, err := sqlite3.DB.Exec(sqlite3.Schemas.StackUser)
		if err != nil {
			return err
		}
		break
	}
	return nil
}
