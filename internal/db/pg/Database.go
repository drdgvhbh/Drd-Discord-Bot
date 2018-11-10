package pg

import (
	"database/sql"
)

type Connector interface {
	QueryRow(query string, args ...interface{}) Row
}

type connectorImp struct {
	db *sql.DB
}

func (connector connectorImp) QueryRow(query string, args ...interface{}) Row {
	return connector.db.QueryRow(query, args...)
}
