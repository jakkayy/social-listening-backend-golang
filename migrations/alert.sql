CREATE TABLE IF NOT EXISTS alerts (
    id SERIAL PRIMARY KEY,
    alert_type TEXT NOT NULL,
    message TEXT NOT NULL,
    metric_value NUMERIC,
    created_at TIMESTAMP DEFAULT NOW()
);
