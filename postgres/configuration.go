package postgres

import "os"

var (
	POSTGRES_HOST     = "heimserver"
	POSTGRES_PORT     = "5432"
	POSTGRES_USER     = "postgres"
	POSTGRES_PASSWORD = "postgres"
	POSTGRES_DBNAME   = "postgres"
)

func CheckEnv() {
	if os.Getenv("POSTGRES_HOST") != "" {
		POSTGRES_HOST = os.Getenv("POSTGRES_HOST")
	}
	if os.Getenv("POSTGRES_PORT") != "" {
		POSTGRES_PORT = os.Getenv("POSTGRES_PORT")
	}
	if os.Getenv("POSTGRES_USER") != "" {
		POSTGRES_USER = os.Getenv("POSTGRES_USER")
	}
	if os.Getenv("POSTGRES_PASSWORD") != "" {
		POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	}
	if os.Getenv("POSTGRES_DBNAME") != "" {
		POSTGRES_DBNAME = os.Getenv("POSTGRES_DBNAME")
	}
}
