package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Alert struct {
	Type        string  `json:"type"`
	Message     string  `json:"message"`
	MetricValue float64 `json:"metric_value"`
}

type AlertRepository struct {
	db *pgxpool.Pool
}

func NewAlertRepository(db *pgxpool.Pool) *AlertRepository {
	return &AlertRepository{db: db}
}

func (r *AlertRepository) Save(ctx context.Context, a Alert) error {
	_, err := r.db.Exec(ctx,
		`INSERT INTO alerts (alert_type, message, metric_value)
		 VALUES ($1,$2,$3)`,
		a.Type, a.Message, a.MetricValue,
	)
	return err
}

func (r *AlertRepository) Latest(ctx context.Context, limit int) ([]Alert, error) {
	rows, err := r.db.Query(ctx,
		`SELECT alert_type, message, metric_value
		 FROM alerts
		 ORDER BY created_at DESC
		 LIMIT $1`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []Alert
	for rows.Next() {
		var a Alert
		if err := rows.Scan(&a.Type, &a.Message, &a.MetricValue); err != nil {
			return nil, err
		}
		out = append(out, a)
	}
	return out, nil
}

func (r *AlertRepository) ExistsRecent(
	ctx context.Context,
	alertType string,
	windowMinutes int,
) (bool, error) {

	var exists bool
	err := r.db.QueryRow(ctx, `
		SELECT EXISTS (
			SELECT 1
			FROM alerts
			WHERE alert_type = $1
			  AND created_at >= NOW() - make_interval(mins => $2)
		)
	`, alertType, windowMinutes).Scan(&exists)

	return exists, err
}
