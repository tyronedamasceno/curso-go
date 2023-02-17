package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // connection driver to mysql
)

// Connect opens db connection
func Connect() (*sql.DB, error) {
	connectionStr := "user:password@/db?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", connectionStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
