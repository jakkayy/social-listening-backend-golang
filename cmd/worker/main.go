package worker

import (
	"context"
	"log"
	"time"

	"social-listening-backend-golang/internal/config"
	"social-listening-backend-golang/internal/domain"
	"social-listening-backend-golang/internal/ingestion/collector"
	"social-listening-backend-golang/internal/insight"
	"social-listening-backend-golang/internal/processing"
	"social-listening-backend-golang/internal/storage"
)

func main() {
	log.Println("worked start")

	db := config.NewDB()
	defer db.Close()

	commentRepo := storage.NewCommentRepository(db)
	analysisRepo := storage.NewAnalysisRepository(db)
	alertRepo := storage.NewAlertRepository(db)

	prevNegative := 0

	ctx := context.Background()

	for {
		log.Println("collecting comments ...")

		comments := collector.CollectorMockComments()

		for _, comment := range comments {
			if err := commentRepo.Save(ctx, comment); err != nil {
				log.Println("save comment eerror: ", err)
				continue
			}

			analysis := domain.CommentAnalysis{
				CommentID: comment.ID,
				Sentiment: domain.Sentiment(processing.AnalizeSentiment(comment.Message)),
				Intent:    domain.Intent(processing.DetectIntent(comment.Message)),
			}

			if err := analysisRepo.Save(ctx, analysis); err != nil {
				log.Println("save analysis error: ", err)
				continue
			}
		}

		var currNegative int
		row := db.QueryRow(ctx, `
			SELECT COUNT(*) FROM comment_analysis WHERE sentiment='negative'
		`)
		_ = row.Scan(&currNegative)

		change := insight.PercentChange(prevNegative, currNegative)

		if change > 30 {
			_ = alertRepo.Save(ctx, storage.Alert{
				Type:        "negative_spike",
				Message:     "Negative sentiment increased significantly",
				MetricValue: change,
			})
		}

		log.Println("worker sleep 30s")
		time.Sleep(30 * time.Second)
	}
}
