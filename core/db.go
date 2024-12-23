package core

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	return db, nil
}
