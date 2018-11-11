package pg

import (
	"database/sql"
)

type SQLDb interface {
	QueryRow(query string, args ...interface{}) *sql.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
}

type Connector interface {
	QueryRow(query string, args ...interface{}) Row
	Exec(query string, args ...interface{}) (sql.Result, error)
}

type connectorImp struct {
	db SQLDb
}

func (
	connector connectorImp,
) QueryRow(
	query string, args ...interface{},
) Row {
	return connector.db.QueryRow(query, args...)
}

func (
	connector connectorImp,
) Exec(
	query string, args ...interface{},
) (sql.Result, error) {
	return connector.db.Exec(query, args...)
}
