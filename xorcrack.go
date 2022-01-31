package main

import "encoding/hex"

var englishMostCommon []byte = []byte{' ', 'e', 't', 'a', 'o', 'i', 'n', 's', 'h', 'r', 'd', 'l', 'u'}

func buildScoreMap(base []byte) map[byte]struct{} {
	result := map[byte]struct{}{}

	for _, value := range base {
		result[value] = struct{}{}
	}

	return result
}

func checkKey(scoringTable map[byte]struct{}, inputByte []byte, key byte) ([]byte, int) {
	outputByte := make([]byte, len(inputByte))
	score := 0

	for i, value := range inputByte {
		decoded_value := value ^ key
		outputByte[i] = decoded_value
		_, exists := scoringTable[decoded_value]
		if exists {
			score++
		}
	}

	return outputByte, score
}

func CheckAllCombinationOfSingleKey(input string) (string, byte, int) {
	inputByte, _ := hex.DecodeString(input)
	letterCount := buildScoreMap(englishMostCommon)

	var bestKey byte
	var bestOutput []byte
	scoreOfBestKey := 0

	for key := byte(0); key < 255; key++ {
		outputByte, currentScore := checkKey(letterCount, inputByte, key)

		if currentScore > scoreOfBestKey {
			bestKey = key
			bestOutput = outputByte
			scoreOfBestKey = currentScore
		}
	}

	return string(bestOutput), bestKey, scoreOfBestKey
}
