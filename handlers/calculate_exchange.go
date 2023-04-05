package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CalcExchange(c *gin.Context) {
	amount, from, to, rate := c.Param("amount"), c.Param("from"), c.Param("to"), c.Param("rate")
	println(from)
	original, err := strconv.ParseFloat(amount, 2); if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount parameter", "message": err.Error()}); return
	}
	conversion_rate, err := strconv.ParseFloat(rate, 2); if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount parameter", "message": err.Error()}); return
	}
	simbol, err := CurrencySim(to); if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return
	}
	c.JSON(http.StatusOK, &Reponse{
		Result: original * conversion_rate,
		Simbol: simbol,
	})
}