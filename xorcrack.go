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

	'E': 12,
	'T': 11,
	'A': 10,
	'O': 9,
	'I': 8,
	'N': 7,
	'S': 6,
	'H': 5,
	'R': 4,
	'D': 3,
	'L': 2,
	'U': 1,
}

func checkKey(scoringTable CommonCharacterTable, inputByte []byte, key byte) ([]byte, int) {
	outputByte := make([]byte, len(inputByte))
	score := 0

	for i, value := range inputByte {
		decodedValue := value ^ key
		outputByte[i] = decodedValue
		value, exists := scoringTable[decodedValue]
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
