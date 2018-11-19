package postgres

import (
	"log"
	"os"

	"github.com/jackc/pgx"
)

func NewConnection() *pgx.Conn {
	config, err := pgx.ParseEnvLibpq()
	if err != nil {
		log.Fatalln("Unable to parse environment:", err)
	}

	conn, err := pgx.Connect(config)
	if err != nil {
		log.Fatalf("Unable to connection to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}
