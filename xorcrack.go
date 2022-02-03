package main

type CommonCharacterTable map[byte]byte

var englishMostCommon CommonCharacterTable = CommonCharacterTable{
	' ': 13,
	'e': 12,
	't': 11,
	'a': 10,
	'o': 9,
	'i': 8,
	'n': 7,
	's': 6,
	'h': 5,
	'r': 4,
	'd': 3,
	'l': 2,
	'u': 1,
}

func checkKey(scoringTable CommonCharacterTable, inputByte []byte, key byte) ([]byte, int) {
	outputByte := make([]byte, len(inputByte))
	score := 0

	for i, value := range inputByte {
		decoded_value := value ^ key
		outputByte[i] = decoded_value
		value, exists := scoringTable[decoded_value]
		if exists {
			score += int(value)
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
