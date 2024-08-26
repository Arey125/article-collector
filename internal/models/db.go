package models

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func InitDb() (*sql.DB, error) {
    dsn := os.Getenv("DB")
    db, err := sql.Open("sqlite3", dsn)

    if err != nil {
        return nil, err
    }

    createArticleTableStmt := `
        CREATE TABLE IF NOT EXISTS articles (
            link TEXT NOT NULL PRIMARY KEY,
            name TEXT NOT NULL,
            source_id TEXT NOT NULL
        );
    `

    _, err = db.Exec(createArticleTableStmt)
    if err != nil {
        return nil, err
    }

    return db, nil
}

