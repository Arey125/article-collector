package database

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const enforceForeignKeys = `
    PRAGMA foreign_keys = ON;
`

const createArticleTable = `
    CREATE TABLE IF NOT EXISTS articles (
        link TEXT NOT NULL PRIMARY KEY,
        name TEXT NOT NULL,
        source_id TEXT NOT NULL
    );
`

const addArticleStatus = `
    CREATE TABLE IF NOT EXISTS statuses (
        id TEXT NOT NULL PRIMARY KEY,
        name TEXT NOT NULL
    );

    INSERT INTO statuses (id, name) VALUES ('unread', 'unread');
    INSERT INTO statuses (id, name) VALUES ('in_progress', 'in progress');
    INSERT INTO statuses (id, name) VALUES ('read', 'read');
        
    ALTER TABLE articles ADD COLUMN status_id TEXT NOT NULL REFERENCES statuses(id) DEFAULT 'unread';
`

var stmts []string = []string{
	enforceForeignKeys,
	createArticleTable,
	addArticleStatus,
}

func InitDb() (*sql.DB, error) {
	dsn := os.Getenv("DB")
	db, err := sql.Open("sqlite3", dsn)

	if err != nil {
		return nil, err
	}

	err = migrate(db, stmts)
	if err != nil {
		return nil, err
	}

    return db, nil
}
