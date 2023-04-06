package main

import (
	"log"

	"github.com/Obdurat/genesis/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/exchange/:amount/:from/:to/:rate", handlers.CalcExchange)
	r.GET("/exchange/list", handlers.ListExchanges)
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}