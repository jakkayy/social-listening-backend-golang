package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type OverviewCount struct {
	Positive int `json:"positive"`
	Neutral  int `json:"neutral"`
	Negative int `json:"negative"`
}

type OverviewRepository struct {
	db *pgxpool.Pool
}

func NewOverviewRepository(db *pgxpool.Pool) *OverviewRepository {
	return &OverviewRepository{db: db}
}

func (r *OverviewRepository) GetOverview(ctx context.Context) (OverviewCount, error) {
	var result OverviewCount

	rows, err := r.db.Query(ctx, `
		SELECT sentiment, COUNT(*)
		FROM comment_analysis
		GROUP BY sentiment
	`)

	if err != nil {
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		var sentiment string
		var count int
		if err := rows.Scan(&sentiment, &count); err != nil {
			return result, err
		}

		switch sentiment {
		case "positive":
			result.Positive = count
		case "neutral":
			result.Neutral = count
		case "negative":
			result.Negative = count
		}
	}

	return result, err
}
