package handlers

import (
	"net/http"
	"strconv"

	"github.com/Obdurat/genesis/repository"
	"github.com/gin-gonic/gin"
)

func CalcExchange(c *gin.Context) {
	amount, from, to, rate := c.Param("amount"), c.Param("from"), c.Param("to"), c.Param("rate")
	original, err := strconv.ParseFloat(amount, 64); if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount parameter", "message": err.Error()}); return
	}
	conversion_rate, err := strconv.ParseFloat(rate, 64); if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid rate parameter", "message": err.Error()}); return
	}
	_, err = CurrencySim(from); if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return
	}
	simbol, err := CurrencySim(to); if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return
	}

	result := original * conversion_rate

	errCn := make(chan error)

	go func() {
		errCn <- repository.Repo.Save(original, from, to ,conversion_rate, result);
	}()

	select {
	case err := <-errCn:
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return
		}
	default:
	}
	c.JSON(http.StatusOK, &Reponse{
		Result: result,
		Simbol: simbol,
	}); return
}