package infrastructure

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/ryo-kagawa/Vercel/environment"
)

func NewDatabase(environment environment.EnvironmentDatabase, schema string) (*sql.DB, error) {
	db, err := sql.Open("postgres", environment.DATABASE_URL)
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
