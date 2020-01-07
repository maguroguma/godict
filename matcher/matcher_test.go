package matcher

import (
	"testing"

	"github.com/my0k/godict/statik"
)

func BenchmarkJudgeRow(b *testing.B) {
	dictionary, _ := statik.LoadDictionary()
	for i := 0; i < b.N; i++ {
		for _, line := range dictionary {
			JudgeRow(line, []string{"apple", "企業"})
		}
	}
}
