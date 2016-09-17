package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(file)
	lines := []string{}
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		lines = append(lines, strings.TrimSpace(string(line)))
	}

	return lines, nil
}
