package storage

import (
	"context"

	"social-listening-backend-golang/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AnalysisRepository struct {
	db *pgxpool.Pool
}

func NewAnalysisRepository(db *pgxpool.Pool) *AnalysisRepository {
	return &AnalysisRepository{db: db}
}

func (r *AnalysisRepository) Save(ctx context.Context, a domain.CommentAnalysis) error {
	_, err := r.db.Exec(ctx,
		`INSERT INTO comment_analysis (comment_id, sentiment, intent)
		 VALUES ($1,$2,$3)
		 ON CONFLICT (comment_id) DO NOTHING`,
		a.CommentID, a.Sentiment, a.Intent,
	)
	return err
}
