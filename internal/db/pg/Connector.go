package pg

import (
	"database/sql"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
)

var connectorInstance Connector

func CreateConnector(cfg *Config) Connector {
	db, err := sql.Open("postgres", fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		cfg.User(), cfg.Password(), cfg.Database(), cfg.Host(), cfg.Port(),
	))
	if err != nil {
		err = errors.Wrapf(err,
			"Couldn't open connection to postgre database (%s)",
			spew.Sdump(cfg))
		panic(err)
	}

	if err = db.Ping(); err != nil {
		err = errors.Wrapf(err,
			"Couldn't ping postgre database (%s)",
			spew.Sdump(cfg))
		panic(err)
	}

	return databaseImp{db}
}

func ProvideConnector(cfg *Config) Connector {
	if connectorInstance != nil {
		return connectorInstance
	}

	connectorInstance = CreateConnector(cfg)
	return connectorInstance
}
