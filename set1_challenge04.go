package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
)

func findTheEncodedLine(scanner *bufio.Scanner) ([]byte, byte) {
	var theBestEncodedLine []byte
	var thebestKey byte
	theBestScore := 0

	for scanner.Scan() {
		line := scanner.Text()
		inputByte, _ := hex.DecodeString(line)

		encodedLine, key, score := CheckAllCombinationOfSingleKey(inputByte)

		if score > theBestScore {
			theBestScore = score
			theBestEncodedLine = encodedLine
			thebestKey = key
		}
	}

	return theBestEncodedLine, thebestKey
}

func MainSet1Challenge04() {
	println("Detect single-character XOR")

	file, readError := os.Open("4.txt")
	if readError != nil {
		fmt.Println(readError)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	encodedLine, key := findTheEncodedLine(scanner)

	if err := scanner.Err(); err != nil {
		println(err)
	} else {
		println("Encoded line:", string(encodedLine))
		println("Key:", string(key))
	}
}
