package database

import (
	"database/sql"
	"fmt"
	"os"
	"sort"
)

func migrate(db *sql.DB) error {
    stmts := getAllMigrationStmts()
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

func getVersion(tx *sql.Tx) (int, error) {
    var version int
    err := tx.QueryRow("PRAGMA user_version").Scan(&version)
    return version, err
}

func setVersion(tx *sql.Tx, version int) error {
    _, err := tx.Exec(fmt.Sprintf("PRAGMA user_version = %d", version))
    return err
}

func getAllMigrationStmts() []string {
    migrationFiles, err := os.ReadDir("migrations")
    if err != nil {
        panic(err)
    }

    filePaths := make([]string, 0, len(migrationFiles))
    for _, file := range migrationFiles {
        if file.IsDir() {
            continue
        }
        filePaths = append(filePaths, fmt.Sprintf("./migrations/%s", file.Name()))
    }
    sort.Strings(filePaths)

    stmts := make([]string, 0, len(migrationFiles))
    for _, filePath := range filePaths {
        content, err := os.ReadFile(filePath)
        if err != nil {
            panic(err)
        }

        stmts = append(stmts, string(content))
    }

    return stmts
}
