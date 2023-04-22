package db

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // DB Driver
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.ConnStr)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
