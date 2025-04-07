package infrastructure

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func NewDatabase(databaseURL string, schema string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}
	if schema != "" {
		_, err := db.Exec("SET search_path TO " + schema)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}
