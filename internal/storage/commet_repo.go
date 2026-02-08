package storage

import (
	"context"

	"social-listening-backend-golang/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CommentRepository struct {
	db *pgxpool.Pool
}

func NewCommentRepository(db *pgxpool.Pool) *CommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) Save(ctx context.Context, c domain.Comment) error {
	_, err := r.db.Exec(ctx,
		`INSERT INTO comments (id, post_id, message, like_count, create_at)
		 VALUES ($1, $2, $3, $4, $5)
		 ON CONFLICT (id) DO RETURNING`,
		c.ID, c.PostID, c.Message, c.LikeCount, c.CreateAt,
	)
	return err
}
