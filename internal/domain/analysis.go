package domain

type Sentiment string
type Intent string

const (
	SentimentPositive Sentiment = "positive"
	SentimentNeutral  Sentiment = "neutral"
	SentimentNegative Sentiment = "negative"
)

const (
	IntentPurchase  Intent = "purchase"
	IntentComplaint Intent = "complaint"
	IntentQuestion  Intent = "question"
	IntentOther     Intent = "other"
)

type CommentAnalysis struct {
	CommentID string
	Sentiment Sentiment
	Intent    Intent
	Keywords  []string
}
