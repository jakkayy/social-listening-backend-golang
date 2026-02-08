CREATE TABLE comments (
    id TEXT PRIMARY KEY,
    post_id TEXT,
    message TEXT,
    like_count INT,
    created_at TIMESTAMP
);

CREATE TABLE comment_analysis (
    comment_id TEXT PRIMARY KEY,
    sentiment TEXT,
    intent TEXT
);
