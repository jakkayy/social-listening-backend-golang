package main

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
	log.Println("worker start")

	db := config.NewDB()
	defer db.Close()

	commentRepo := storage.NewCommentRepository(db)
	analysisRepo := storage.NewAnalysisRepository(db)
	alertRepo := storage.NewAlertRepository(db)
	trendRepo := storage.NewAnalysisTrendRepository(db)
	keywordRepo := storage.NewKeywordTrendRepository(db)
	dailyRepo := storage.NewDailyInsightRepository(db)

	ctx := context.Background()

	for {
		log.Println("collecting comments...")

		comments := collector.CollectorMockComments()

		for _, comment := range comments {
			_ = commentRepo.Save(ctx, comment)

			analysis := domain.CommentAnalysis{
				CommentID: comment.ID,
				Sentiment: domain.Sentiment(processing.AnalizeSentiment(comment.Message)),
				Intent:    domain.Intent(processing.DetectIntent(comment.Message)),
			}
			_ = analysisRepo.Save(ctx, analysis)
		}

		currNeg, _ := trendRepo.CountNegativeLastMinutes(ctx, 10)
		prevNeg, _ := trendRepo.CountNegativePrevWindow(ctx, 20, 10)

		change := insight.PercentChange(prevNeg, currNeg)
		log.Printf(
			"negative trend window: prev=%d curr=%d change=%.2f%%",
			prevNeg, currNeg, change,
		)

		if change > 30 {
			exists, _ := alertRepo.ExistsRecent(ctx, "negative_spike", 30)
			if !exists {
				_ = alertRepo.Save(ctx, storage.Alert{
					Type:        "negative_spike",
					Message:     "Negative sentiment spike detected (10m window)",
					MetricValue: change,
				})
				log.Println("alert saved: negative_spike")
			}
		}

		for _, kw := range insight.Keywords {

			curr, _ := keywordRepo.CountKeywordLastMinutes(ctx, kw, 10)
			prev, _ := keywordRepo.CountKeywordPrevWindow(ctx, kw, 20, 10)

			change := insight.PercentChange(prev, curr)

			log.Printf(
				"keyword trend [%s]: prev=%d curr=%d change=%.2f%%",
				kw, prev, curr, change,
			)

			if change > 50 {
				alertType := "keyword_spike:" + kw

				exists, _ := alertRepo.ExistsRecent(ctx, alertType, 30)
				if !exists {
					_ = alertRepo.Save(ctx, storage.Alert{
						Type:        alertType,
						Message:     "Keyword spike detected: " + kw,
						MetricValue: change,
					})
					log.Println("alert saved:", alertType)
				}
			}
		}

		today := time.Now().Truncate(24 * time.Hour)

		total, _ := storage.CountCommentsToday(ctx, db)
		pos, _ := storage.CountSentimentToday(ctx, db, "positive")
		neu, _ := storage.CountSentimentToday(ctx, db, "neutral")
		neg, _ := storage.CountSentimentToday(ctx, db, "negative")
		alertsToday, _ := storage.CountAlertsToday(ctx, db)
		keywords, _ := storage.TopKeywordsToday(ctx, db, 5)

		_ = dailyRepo.Upsert(ctx, storage.DailyInsight{
			InsightDate:   today,
			TotalComments: total,
			PositiveCount: pos,
			NeutralCount:  neu,
			NegativeCount: neg,
			TopKeywords:   keywords,
			AlertCount:    alertsToday,
		})

		log.Println("daily insight snapshot upserted")

		log.Println("worker sleep 30s")
		time.Sleep(30 * time.Second)
	}
}
