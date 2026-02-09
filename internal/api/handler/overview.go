package handler

import (
	"net/http"
	"social-listening-backend-golang/internal/config"
	"social-listening-backend-golang/internal/storage"

	"github.com/gin-gonic/gin"
)

func OverviewHandler(c *gin.Context) {
	db := config.NewDB()
	defer db.Close()

	overviewRopo := storage.NewOverviewRepository(db)

	result, err := overviewRopo.GetOverview(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)

}
