package handlers

import (
	"net/http"
	"strconv"

	"github.com/Obdurat/genesis/repository"
	"github.com/gin-gonic/gin"
)

func ListExchanges(c *gin.Context) {
	offset, err := strconv.ParseInt(c.Param("page"), 0, 8); if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Page parameter must be a number "+err.Error()}); return
	}
	if offset < 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Page parameter must greater than 1"}); return
	}
	results, err := repository.Repo.List(offset); if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return
	}
	c.JSON(http.StatusOK, results); return
}