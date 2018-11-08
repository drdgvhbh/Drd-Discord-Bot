package pg

import (
	"database/sql"
)

type Connector interface {
	QueryRow(query string, args ...interface{}) Row
}

type databaseImp struct {
	db *sql.DB
}

func (db databaseImp) QueryRow(query string, args ...interface{}) Row {
	return db.QueryRow(query, args)
}
