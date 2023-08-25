package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/libsql/libsql-client-go/libsql"
	_ "modernc.org/sqlite"
)

type DbConfig struct {
	DBURL   string
	DBToken string
}

func Connect(url string) (*sql.DB, error) {
	db, err := sql.Open("libsql", url)
	if err != nil {
		return nil, fmt.Errorf("failed to open db %s: %s", url, err)
	}
	return db, nil
}

func BuildURL(cfg DbConfig) string {
	var dbURL string
	if cfg.DBToken != "" {
		dbURL = fmt.Sprintf("%s?authToken=%s", cfg.DBURL, cfg.DBToken)
	} else {
		dbURL = cfg.DBURL
	}

	return dbURL
}

func Close(db *sql.DB) error {
	return db.Close()
}

func Ping(db *sql.DB) error {
	return db.Ping()
}
