package decorator

import (
	"strings"
)

// ReplaceToDecorated returns the text where words are decorated those.
func ReplaceToDecorated(
	text string, words []string, decorator func(a ...interface{}) string,
) string {
	for _, w := range words {
		text = strings.Replace(text, w, decorator(w), -1)
	}

	return text
}
