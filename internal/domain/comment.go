package domain

import "time"

type Comment struct {
	ID        string
	PostID    string
	Message   string
	LikeCount int
	CreateAt  time.Time
}
