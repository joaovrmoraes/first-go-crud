package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Connect() (*sql.DB, error) {
	sqliteConn := "database.db"

	db, err := sql.Open("sqlite3", sqliteConn)

	if err != nil {
		return nil, err
	}

	return db, nil
}
