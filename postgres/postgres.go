package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var Client Postgres

type Postgres struct {
	Connection *sql.DB
}

func ConnectToPostgres() {
	CheckEnv()
	psqlInfo := fmt.Sprintf("host= %s port=%s user=%s password=%s dbname=%s sslmode=disable",
		POSTGRES_HOST, POSTGRES_PORT, POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DBNAME)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	Client.Connection = db
}

func (c *Postgres) ExecQuery(query string) (*sql.Rows, error) {
	result, err := c.Connection.Query(query)
	if err != nil {
		return nil, err
	}

	return result, nil
}
