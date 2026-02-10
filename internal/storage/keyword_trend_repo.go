package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type KeywordTrendRepository struct {
	db *pgxpool.Pool
}

func NewKeywordTrendRepository(db *pgxpool.Pool) *KeywordTrendRepository {
	return &KeywordTrendRepository{db: db}
}

func (r *KeywordTrendRepository) CountKeywordLastMinutes(
	ctx context.Context,
	keyword string,
	minutes int,
) (int, error) {

	var count int
	err := r.db.QueryRow(ctx, `
		SELECT COUNT(*)
		FROM comments
		WHERE message ILIKE '%' || $1 || '%'
			AND createed_at >= NOW() - make_interval(mins => $2)	
	`, keyword, minutes).Scan(&count)

	return count, err
}

func (r *KeywordTrendRepository) CountKeywordPrevWindow(
	ctx context.Context,
	keyword string,
	startMinutesAgo int,
	endMinutesAgo int,
) (int, error) {

	var count int
	err := r.db.QueryRow(ctx, `
		SELECT COUNT(*)
		FROM comments
		WHERE message ILIKE '%' || $1 || '%'
			AND created_at BETWEEN
				NOW() - make_interval(mins => $2)
			AND NOW() - make_interval(mins => $3)
	`, keyword, startMinutesAgo, endMinutesAgo).Scan(&count)

	return count, err
}
