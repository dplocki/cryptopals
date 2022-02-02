package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

func loadFiles(fileName string) string {
	builder := strings.Builder{}

	file, readError := os.Open(fileName)
	if readError != nil {
		fmt.Println(readError)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		builder.WriteString(scanner.Text())
	}

	return builder.String()
}

func findKeyLenght(encryptedMessage []byte) int {
	theBestScore := 0
	theBestKeyLength := 0

	for keyLength := 2; keyLength <= 40; keyLength++ {
		currentScore := HammingDistance(encryptedMessage[0:keyLength], encryptedMessage[keyLength:keyLength*2]) / keyLength
		if currentScore > theBestScore {
			theBestScore = currentScore
			theBestKeyLength = keyLength
		}
	}

	return theBestKeyLength
}

func MainSet1Challenge06() {
	content := loadFiles("6.txt")

	originalStringBytes, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		fmt.Printf("Some error occured during base64 decode. Error %q", err.Error())
	}

	key := findKeyLenght(originalStringBytes)

	fmt.Println(key)
}
