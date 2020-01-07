package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/my0k/godict/matcher"
	"github.com/my0k/godict/statik"
)

// go:generate statik -src dictext

func main() {
	var (
		words = flag.String("key", "", "search keywords (AND: \"word1 word2 word3 ..\")")
	)
	flag.Parse()
	if *words == "" {
		fmt.Fprint(os.Stdout, "please input a word\n")
		return
	}
	wordSlice := strings.Split(*words, " ")

	dictionary, err := statik.LoadDictionary()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	for _, line := range dictionary {
		sline := strings.ToLower(string(line))
		if matcher.JudgeRow(sline, wordSlice) {
			cols := strings.Split(line, "\t")

			// "\x1b[1m": bold, "\x1b[0m": default, "\x1b[31m": red, "\x1b[39m"
			str := fmt.Sprintf("\x1b[31m\x1b[1m [%s] \x1b[39m\x1b[0m\n%s\n", cols[0], cols[1])
			fmt.Printf(str)
		}
	}
}
