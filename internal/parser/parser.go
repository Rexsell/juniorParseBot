package parser

import (
	"strings"
)

func FindKeyword(text string, keywords []string) bool {
	text = strings.ToLower(text)
	for _, keyword := range keywords {
		if res := strings.Contains(text, keyword); res {
			return true
		}
	}
	return false
}
