package insight

import (
	"social-listening-backend-golang/internal/domain"
)

type OverviewInsight struct {
	Positive int `json:"positive"`
	Neutral  int `json:"neutral"`
	Negative int `json:"negative"`
}

func GenerateOverview(analyses []domain.CommentAnalysis) OverviewInsight {
	var o OverviewInsight
	for _, a := range analyses {
		switch a.Sentiment {
		case domain.SentimentPositive:
			o.Positive++
		case domain.SentimentNeutral:
			o.Neutral++
		case domain.SentimentNegative:
			o.Negative++
		}
	}
	return o
}
