package storage

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Queryer interface {
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
}
