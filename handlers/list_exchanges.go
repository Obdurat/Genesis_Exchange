package handlers

import (
	"net/http"

	"github.com/Obdurat/genesis/repository"
	"github.com/gin-gonic/gin"
)

func ListExchanges(c *gin.Context) {
	results, err := repository.Repo.List(); if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return
	}
	c.JSON(http.StatusOK, results); return
}