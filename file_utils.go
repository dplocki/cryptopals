package main

import (
	"bufio"
	"os"
	"strings"
)

func LoadFileContentAsString(fileName string) string {
	builder := strings.Builder{}

	file, readError := os.Open(fileName)
	if readError != nil {
		panic("cannot load file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		builder.WriteString(scanner.Text())
	}

	return builder.String()
}
