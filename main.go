package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:12345678@tcp(localhost:3306)/genesis_logs")
	if err != nil {
		panic(err)
	}
	err = db.Ping(); if err != nil {
		panic(err)
	}
	fmt.Println("Okay")
	defer db.Close()
}