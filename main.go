package main

import (
	"fmt"

	"github.com/Obdurat/genesis/infra"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := infra.New()
	if err := db.Connect("mysql", "root:12345678@tcp(localhost:3306)/genesis_logs"); err != nil {
		panic(err)
	}
	fmt.Println("Okay")
}