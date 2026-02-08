package processing

import "strings"

func DetectIntent(text string) string {
	if strings.Contains(text, "ราคา") {
		return "purchase"
	}
	if strings.Contains(text, "ช้า") {
		return "complataint"
	}
	if strings.Contains(text, "?") {
		return "question"
	}
	return "other"
}
