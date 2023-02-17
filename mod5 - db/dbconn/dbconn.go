package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	connectionStr := "user:password@/db?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", connectionStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// fmt.Println(db)

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("connection is open")

	lines, err := db.Query("select * from users")

	if err != nil {
		log.Fatal(err)
	}
	defer lines.Close()

	fmt.Println(lines)
}
