package main

import (
	"social-listening-backend-golang/internal/api/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/api/overview", handler.OverviewHandler)

	r.GET("/api/alerts", handler.AlertHandler)

	r.Run(":8080")
}
