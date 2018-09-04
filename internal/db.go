package bkpr

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

const (
	DEFAULT_DB_PATH = "/var/db/bkpr/db.sqlite"
)

func Db_connect() (*sql.DB, error) {
	return sql.Open("sqlite3", Context.Db)
}
