package database

import (
	"database/sql"
	"fmt"
)

func getVersion(tx *sql.Tx) (int, error) {
    var version int
    err := tx.QueryRow("PRAGMA user_version").Scan(&version)
    return version, err
}

func setVersion(tx *sql.Tx, version int) error {
    _, err := tx.Exec(fmt.Sprintf("PRAGMA user_version = %d", version))
    return err
}

func migrate(db *sql.DB, stmts []string) error {
    tx, err := db.Begin()
    if err != nil {
        panic(err)
    }

    defer tx.Rollback()

    version, err := getVersion(tx)
    if err != nil {
        panic(err)
    }

    for _, stmt := range stmts[version:] {
        _, err := tx.Exec(stmt)
        if err != nil {
            fmt.Println(stmt)
            panic(err)
        }
    }

    err = setVersion(tx, len(stmts))
    if err != nil {
        panic(err)
    }

    return tx.Commit()
}
