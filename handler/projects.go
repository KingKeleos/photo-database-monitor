package handler

import "github.com/kingkeleos/database-monitor/postgres"

func GetProjectCount() (float64, error) {
	rows, err := postgres.Client.ExecQuery("SELECT Count(*) FROM projects;")
	if err != nil {
		return 0, err
	}

	var count float64
	for rows.Next() {
		rows.Scan(&count)
	}

	return count, err
}
