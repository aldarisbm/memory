package sqlite

import (
	"database/sql"
	"fmt"
	"os"
)

const CheckSchemaExistsQuery = `
		SELECT 1
		FROM sqlite_master
		WHERE type = 'table' AND name = 'schema_info';
`

const MarkSchemaCreatedQuery = `
		CREATE TABLE schema_info (
			id INTEGER PRIMARY KEY
		);
		INSERT INTO schema_info (id) VALUES (1);
	`

func createTable(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("opening database: %w", err)
	}

	schemaExists, err := checkSchemaExists(db)
	if err != nil {
		fmt.Println("Error checking schema existence:", err)
	}
	if !schemaExists {
		err = createSchema(db)
		if err != nil {
			return nil, fmt.Errorf("creating schema: %w", err)
		}
		fmt.Println("local store has been created")
	} else {
		fmt.Println("local store has already been created.")
	}
	return db, nil
}

func createSchema(db *sql.DB) error {
	schema, err := os.ReadFile("schema.sql")
	if err != nil {
		return fmt.Errorf("reading schema file: %w", err)
	}
	_, err = db.Exec(string(schema))
	if err != nil {
		return fmt.Errorf("executing SQL statement: %w", err)
	}

	err = markSchemaCreated(db)
	if err != nil {
		return fmt.Errorf("marking schema as created: %w", err)
	}
	return nil
}

func checkSchemaExists(db *sql.DB) (bool, error) {
	query := CheckSchemaExistsQuery
	var exists int
	err := db.QueryRow(query).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return exists == 1, nil
}

func markSchemaCreated(db *sql.DB) error {
	query := MarkSchemaCreatedQuery
	_, err := db.Exec(query)
	return err
}
