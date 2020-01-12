package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/my0k/godict/decorator"
	"github.com/my0k/godict/matcher"
	"github.com/my0k/godict/statik"
)

// go:generate statik -src dictext

func main() {
	var wordSlice []string
	flag.Parse()
	wordSlice = flag.Args()
	if len(wordSlice) == 0 {
		fmt.Fprint(os.Stdout, "please input some words\n")
	}
	for i, w := range wordSlice {
		wordSlice[i] = strings.ToLower(w)
	}

	dictionary, err := statik.LoadDictionary()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	enDecorator := color.New(color.FgRed, color.Bold, color.Underline).SprintFunc()
	jaDecorator := color.New(color.FgHiWhite, color.BgMagenta).SprintFunc()

	for _, line := range dictionary {
		sline := strings.ToLower(string(line))
		if matcher.JudgeRow(sline, wordSlice) {
			cols := strings.Split(line, "\t")

			enText := decorator.ReplaceToDecorated(cols[0], []string{cols[0]}, enDecorator)
			jaText := decorator.ReplaceToDecorated(cols[1], wordSlice, jaDecorator)
			fmt.Fprintf(os.Stdout, "%s - %s\n", enText, jaText)
		}
	}
}
