package worker

import (
	"context"
	"log"
	"time"

	"social-listening-backend-golang/internal/config"
	"social-listening-backend-golang/internal/domain"
	"social-listening-backend-golang/internal/ingestion/collector"
	"social-listening-backend-golang/internal/processing"
	"social-listening-backend-golang/internal/storage"
)

func main() {
	log.Println("worked start")

	db := config.NewDB()
	defer db.Close()

	commentRepo := storage.NewCommentRepository(db)
	analysisRepo := storage.NewAnalysisRepository(db)

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

		log.Println("worker sleep 30s")
		time.Sleep(30 * time.Second)
	}
}
