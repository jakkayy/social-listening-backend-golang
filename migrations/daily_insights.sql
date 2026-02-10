CREATE TABLE IF NOT EXISTS daily_insights (
    insight_date DATE PRIMARY KEY,

    total_comments INT NOT NULL,
    positive_count INT NOT NULL,
    neutral_count  INT NOT NULL,
    negative_count INT NOT NULL,
    top_keywords TEXT[], 
    alert_count  INT NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
