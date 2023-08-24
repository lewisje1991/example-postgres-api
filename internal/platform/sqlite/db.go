package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/libsql/libsql-client-go/libsql"
	_ "modernc.org/sqlite"
)

func Connect(url string) (*sql.DB, error) {
	db, err := sql.Open("libsql", url)
	if err != nil {
		return nil, fmt.Errorf("failed to open db %s: %s", url, err)
	}
	return db, nil
}

func Close(db *sql.DB) error {
	return db.Close()
}

func Ping(db *sql.DB) error {
	return db.Ping()
}
