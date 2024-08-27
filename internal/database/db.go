package database

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const enforceForeignKeys = `
    PRAGMA foreign_keys = ON;
`

func InitDb() (*sql.DB, error) {
	dsn := os.Getenv("DB")
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(enforceForeignKeys)
    if err != nil {
        panic(err)
    }

	err = migrate(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}
