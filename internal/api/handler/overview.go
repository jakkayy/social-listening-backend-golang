package handler

import (
	"net/http"

	"social-listening-backend-golang/internal/domain"
	"social-listening-backend-golang/internal/ingestion/collector"
	"social-listening-backend-golang/internal/insight"
	"social-listening-backend-golang/internal/processing"

	"github.com/gin-gonic/gin"
)

func OverviewHandler(c *gin.Context) {
	comments := collector.CollectorMockComments()

	var analyses []domain.CommentAnalysis
	for _, comment := range comments {
		analyses = append(analyses, domain.CommentAnalysis{
			CommentID: comment.ID,
			Sentiment: domain.Sentiment(processing.AnalizeSentiment(comment.Message)),
			Intent:    domain.Intent(processing.DetectIntent(comment.Message)),
		})
	}

	result := insight.GenerateOverview(analyses)

	c.JSON(http.StatusOK, result)
}
