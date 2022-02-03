package main

type EmptyType struct{}
type Set map[byte]EmptyType

var englishMostCommon Set = Set{
	' ': EmptyType{},
	'e': EmptyType{},
	't': EmptyType{},
	'a': EmptyType{},
	'o': EmptyType{},
	'i': EmptyType{},
	'n': EmptyType{},
	's': EmptyType{},
	'h': EmptyType{},
	'r': EmptyType{},
	'd': EmptyType{},
	'l': EmptyType{},
	'u': EmptyType{},
}

func checkKey(scoringTable Set, inputByte []byte, key byte) ([]byte, int) {
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

func CheckAllCombinationOfSingleKey(input []byte) ([]byte, byte, int) {
	var bestKey byte
	var bestOutput []byte
	scoreOfBestKey := 0

	for key := byte(0); key < 255; key++ {
		outputByte, currentScore := checkKey(englishMostCommon, input, key)

		if currentScore > scoreOfBestKey {
			bestKey = key
			bestOutput = outputByte
			scoreOfBestKey = currentScore
		}
	}

	return bestOutput, bestKey, scoreOfBestKey
}
