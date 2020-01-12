package matcher

import "strings"

func JudgeRow(line string, searchWords []string) bool {
	for _, w := range searchWords {
		if idx := strings.Index(line, w); idx == -1 {
			return false
		}
	}
	return true
}
