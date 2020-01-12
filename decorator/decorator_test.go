package decorator

import (
	"strings"
	"testing"

	"github.com/fatih/color"
)

// TestColorPackageAndStringsReplace tests how to use Replace and color package functions.
func TestColorPackageAndStringsReplace(t *testing.T) {
	cases := []struct {
		in struct {
			rowText string
			words   []string
		}
	}{
		{
			in: struct {
				rowText string
				words   []string
			}{
				rowText: "apple black cookie death ear",
				words:   []string{"black", "deatch"},
			},
		},
		{
			in: struct {
				rowText string
				words   []string
			}{
				rowText: "東西南北、古今東西",
				words:   []string{"東西", "古"},
			},
		},
	}

	info := color.New(color.FgWhite, color.BgGreen).SprintFunc()

	for _, c := range cases {
		t.Run(c.in.rowText, func(t *testing.T) {
			var out string = c.in.rowText
			for _, w := range c.in.words {
				out = strings.Replace(out, w, info(w), -1)
			}

			t.Logf("%s - %s", c.in.rowText, out)
			if c.in.rowText != out {
				t.Errorf("contents turn to be another one")
			}
		})
	}
}
