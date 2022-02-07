package main

import (
	"bufio"
	"os"
	"strings"
)

func readlines(path string) (<-chan string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	chanel := make(chan string)
	go func() {
		for scanner.Scan() {
			chanel <- scanner.Text()
		}
		close(chanel)
	}()

	return chanel, nil
}

func LoadFileContentAsString(fileName string) string {
	reader, readError := readlines(fileName)
	if readError != nil {
		panic("cannot load file")
	}

	builder := strings.Builder{}
	for line := range reader {
		builder.WriteString(line)
	}

	return builder.String()
}

func LoadFileContentAsStringsArray(fileName string) []string {
	reader, readError := readlines(fileName)
	if readError != nil {
		panic("cannot load file")
	}

	result := []string{}
	for line := range reader {
		result = append(result, line)
	}

	return result
}
