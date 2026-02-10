package storage

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DailyInsight struct {
	InsightDate   time.Time
	TotalComments int
	PositiveCount int
	NeutralCount  int
	NegativeCount int
	TopKeywords   []string
	AlertCount    int
}

type DailyInsightRepository struct {
	db *pgxpool.Pool
}

func NewDailyInsightRepository(db *pgxpool.Pool) *DailyInsightRepository {
	return &DailyInsightRepository{db: db}
}

func (r *DailyInsightRepository) Upsert(
	ctx context.Context,
	d DailyInsight,
) error {

	_, err := r.db.Exec(ctx, `
		INSERT INTO daily_insights (
			insight_date,
			total_comments,
			positive_count,
			neutral_count,
			negative_count,
			top_keywords,
			alert_count
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7)
		ON CONFLICT (insight_date)
		DO UPDATE SET
			total_comments = EXCLUDED.total_comments,
			positive_count = EXCLUDED.positive_count,
			neutral_count  = EXCLUDED.neutral_count,
			negative_count = EXCLUDED.negative_count,
			top_keywords   = EXCLUDED.top_keywords,
			alert_count    = EXCLUDED.alert_count
	`, d.InsightDate, d.TotalComments, d.PositiveCount, d.NeutralCount, d.NegativeCount, d.TopKeywords, d.AlertCount)

	return err
}
