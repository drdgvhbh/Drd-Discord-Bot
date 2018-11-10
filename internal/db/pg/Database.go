package pg

import (
	"database/sql"
)

type SQLDb interface {
	QueryRow(query string, args ...interface{}) *sql.Row
}

type Connector interface {
	QueryRow(query string, args ...interface{}) Row
}

type connectorImp struct {
	db SQLDb
}

func (connector connectorImp) QueryRow(query string, args ...interface{}) Row {
	return connector.db.QueryRow(query, args...)
}
