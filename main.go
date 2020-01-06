package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	_ "github.com/my0k/go-dict/statik"
	"github.com/rakyll/statik/fs"
)

// go:generate statik -src dictext

func main() {
	var (
		word = flag.String("key", "", "search key word")
	)
	flag.Parse()
	if *word == "" {
		fmt.Fprint(os.Stdout, "please input a word\n")
		return
	}

	statikFs, err := fs.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Error] cannot create statikFs: error=%v\n", err)
		return
	}

	r, err := statikFs.Open("/ejdic-hand-utf8.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Error] cannot open dictionary file: error=%v\n", err)
		return
	}
	defer r.Close()

	reader := bufio.NewReaderSize(r, 4096)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "[Error] cannot read line: error=%v\n", err)
			return
		}

		sline := strings.ToLower(string(line))
		if idx := strings.Index(sline, *word); idx != -1 {
			fmt.Println(string(line))
		}
	}
}
