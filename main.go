package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Obdurat/genesis/handlers"
	"github.com/Obdurat/genesis/infra"
	"github.com/gin-gonic/gin"
)

func main() {
	err := infra.DB.Connect("mysql", fmt.Sprintf("root:%s@tcp(db:3306)/genesis", os.Getenv("MYSQL_ROOT_PASSWORD")))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer infra.DB.Disconnect()
	r := gin.Default()
	r.POST("/exchange/:amount/:from/:to/:rate", handlers.CalcExchange)
	r.GET("/exchange/list/:page", handlers.ListExchanges)
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}