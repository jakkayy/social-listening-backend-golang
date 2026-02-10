package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AnalysisTrendRepository struct {
	db *pgxpool.Pool
}

func NewAnalysisTrendRepository(db *pgxpool.Pool) *AnalysisTrendRepository {
	return &AnalysisTrendRepository{db: db}
}

func (r *AnalysisTrendRepository) CountNegativeLastMinutes(
	ctx context.Context,
	minutes int,
) (int, error) {

	var count int
	err := r.db.QueryRow(ctx, `
		SELECT COUNT(*)
		FROM comment_analysis
		WHERE sentiment = 'negative'
		  AND analyzed_at >= NOW() - make_interval(mins => $1)
	`, minutes).Scan(&count)

	return count, err
}

func (r *AnalysisTrendRepository) CountNegativePrevWindow(
	ctx context.Context,
	startMinutesAgo int,
	endMinutesAgo int,
) (int, error) {

	var count int
	err := r.db.QueryRow(ctx, `
		SELECT COUNT(*)
		FROM comment_analysis
		WHERE sentiment = 'negative'
		  AND analyzed_at BETWEEN
		      NOW() - make_interval(mins => $1)
		  AND NOW() - make_interval(mins => $2)
	`, startMinutesAgo, endMinutesAgo).Scan(&count)

	return count, err
}
