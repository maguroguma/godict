package statik

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/rakyll/statik/fs"
)

func LoadDictionary() ([]string, error) {
	dictionary := []string{}

	statikFs, err := fs.New()
	if err != nil {
		return nil, fmt.Errorf("[Error] cannot create statikFs: error=%v", err)
	}

	r, err := statikFs.Open("/ejdic-hand-utf8.txt")
	if err != nil {
		return nil, fmt.Errorf("[Error] cannot open dictionary file: error=%v", err)
	}
	defer r.Close()

	reader := bufio.NewReaderSize(r, 4096)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, fmt.Errorf("[Error] cannot read line: error=%v", err)
		}

		sline := strings.ToLower(string(line))
		dictionary = append(dictionary, sline)
	}

	return dictionary, nil
}
