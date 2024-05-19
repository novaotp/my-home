package database

import (
	"database/sql"
	"os"
)

// Connects to the SQLite database and returns the connector.
//
// If the database doesn't exist, it is created.
func Connect() (*sql.DB, error) {
	const db_file_path string = "./src/database/data.db"
	const db_setup_path string = "./src/database/setup.sql"

	if _, err := os.Stat(db_file_path); err != nil {
		os.Create(db_file_path)

		b, err := os.ReadFile(db_setup_path)
		if err != nil {
			return nil, err
		}

		setup := string(b)

		db, err := sql.Open("sqlite", db_file_path)
		if err != nil {
			return nil, err
		}

		_, err = db.Exec(setup)
		if err != nil {
			return nil, err
		}

		return db, nil
	}

	db, err := sql.Open("sqlite", db_file_path)
	if err != nil {
		return nil, err
	}

	return db, nil
}
