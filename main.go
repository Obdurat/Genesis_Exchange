package main

import (
	"log"

	"github.com/Obdurat/genesis/handlers"
	"github.com/Obdurat/genesis/infra"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := infra.DB
	if err := db.Connect("mysql", "root:12345678@tcp(localhost:3306)/genesis_logs"); err != nil {
		panic(err)
	}
	r := gin.Default()
	r.GET("/exchange/:amount/:from/:to/:rate", handlers.CalcExchange)
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}