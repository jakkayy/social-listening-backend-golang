package collector

import (
	"social-listening-backend-golang/internal/domain"
	"time"
)

func CollectorMockComments() []domain.Comment {
	return []domain.Comment{
		{
			ID:        "c1",
			PostID:    "p1",
			Message:   "ราคาดีมาก น่าสนใจ",
			LikeCount: 10,
			CreateAt:  time.Now(),
		},
		{
			ID:        "c2",
			PostID:    "p2",
			Message:   "แพงจังเลย",
			LikeCount: 34,
			CreateAt:  time.Now(),
		},
	}
}
