package processing

import (
	"strings"
)

func AnalizeSentiment(text string) string {
	if strings.Contains(text, "ดี") || strings.Contains(text, "น่าสนใจครับ") {
		return "positive"
	}
	if strings.Contains(text, "แพง") || strings.Contains(text, "ช้า") {
		return "negative"
	}
	return "neutral"
}
