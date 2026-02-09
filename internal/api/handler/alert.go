package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"social-listening-backend-golang/internal/config"
	"social-listening-backend-golang/internal/storage"
)

func AlertHandler(c *gin.Context) {
	db := config.NewDB()
	defer db.Close()

	repo := storage.NewAlertRepository(db)
	alert, err := repo.Latest(c, 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, alert)
}
