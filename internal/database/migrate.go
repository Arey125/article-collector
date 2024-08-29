package database

import (
	"database/sql"
	"fmt"
	"os"
	"sort"
)

type migration struct {
    file string
    stmt string
}

func migrate(db *sql.DB) error {
    migrations := getAllMigrations()
    tx, err := db.Begin()
    if err != nil {
        panic(err)
    }

    defer tx.Rollback()

    version, err := getVersion(tx)
    if err != nil {
        panic(err)
    }

    for _, migration := range migrations[version:] {
        fmt.Printf("Migration %s started\n", migration.file)
        _, err := tx.Exec(migration.stmt)
        if err != nil {
            fmt.Printf("Error in migration %s",migration.file)
            panic(err)
        }
        fmt.Println("Done")
    }

    err = setVersion(tx, len(migrations))
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

func getAllMigrations() []migration {
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

    migrations := make([]migration, 0, len(migrationFiles))
    for _, filePath := range filePaths {
        content, err := os.ReadFile(filePath)
        if err != nil {
            panic(err)
        }
        migration := migration{file: filePath, stmt: string(content)}
        migrations = append(migrations, migration)
    }

    return migrations
}
