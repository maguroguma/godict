package statik

import (
	"strings"
	"testing"
)

// Check whether the dictionary can be loaded without errors.
func TestLoadDictionary(t *testing.T) {
	if _, err := LoadDictionary(); err != nil {
		t.Errorf("Failure of loading the dictionary: %v", err)
	}
}

// Check formats of all rows in the dictionary.
func TestAllRowsHaveOneTab(t *testing.T) {
	dictionary, _ := LoadDictionary()
	expect := 2

	for i, str := range dictionary {
		cols := strings.Split(str, "\t")
		actual := len(cols)
		if expect < actual {
			t.Errorf("%d row do not have any tabs", i)
		} else if expect > actual {
			t.Errorf("%d row have more than one tab", i)
		}

		if i >= 1000 && i%1000 == 0 {
			t.Logf("%d done\n", i)
		}
	}
}
