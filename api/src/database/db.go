package database

import (
	"database/sql"
	"fmt"
	"os"
)

// The SQLite database connection.
var Connection *sql.DB

func HandleDatabaseError(err error) {
	fmt.Println(err)
	os.Exit(1)
}

// Connects to the SQLite database and sets the global connector.
//
// If the database file doesn't exist, it is created.
func Connect() {
	const db_file_path string = "./src/database/data.db"
	const db_setup_path string = "./src/database/setup.sql"

	if _, err := os.Stat(db_file_path); err != nil {
		fmt.Println("> Setting up SQLite database")
		os.Create(db_file_path)

		b, err := os.ReadFile(db_setup_path)
		if err != nil {
			HandleDatabaseError(err)
			return
		}

		setup := string(b)

		Connection, err = sql.Open("sqlite", db_file_path)
		if err != nil {
			HandleDatabaseError(err)
			return
		}

		_, err = Connection.Exec(setup)
		if err != nil {
			HandleDatabaseError(err)
			return
		}

		return
	}

	fmt.Println("> SQLite database found")
	var err error
	Connection, err = sql.Open("sqlite", db_file_path)
	if err != nil {
		HandleDatabaseError(err)
		return
	}
}
