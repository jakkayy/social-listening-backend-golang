package handler

import (
	"social-listening-backend-golang/internal/config"
	"social-listening-backend-golang/internal/domain"
	"social-listening-backend-golang/internal/ingestion/collector"
	"social-listening-backend-golang/internal/processing"
	"social-listening-backend-golang/internal/storage"

	"github.com/gin-gonic/gin"
)

func OverviewHandler(c *gin.Context) {
	db := config.NewDB()
	defer db.Close()

	commentRepo := storage.NewCommentRepository(db)
	analysisRepo := storage.NewAnalysisRepository(db)

	comments := collector.CollectorMockComments()

	for _, comment := range comments {
		_ = commentRepo.Save(c, comment)

		analysis := domain.CommentAnalysis{
			CommentID: comment.ID,
			Sentiment: domain.Sentiment(processing.AnalizeSentiment(comment.Message)),
			Intent:    domain.Intent(processing.DetectIntent(comment.Message)),
		}

		_ = analysisRepo.Save(c, analysis)
	}
}
