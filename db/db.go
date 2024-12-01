package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
)

func getDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=localhost port=5432 user=admin dbname=jwt_auth password=%s sslmode=disable", os.Getenv("DB_PASSWORD")))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
