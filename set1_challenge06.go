package main

import (
	"encoding/base64"
	"fmt"
	"math"
)

const MaxUint uint = ^uint(0)

func findKeyLength(encryptedMessage []byte) byte {
	theBestScore := math.MaxFloat64
	theBestKeyLength := byte(0)

	const blockSize byte = 3
	for keyLength := byte(2); keyLength <= 40; keyLength++ {
		currentScore := float64(HammingDistance(encryptedMessage[0:keyLength*blockSize], encryptedMessage[keyLength*blockSize:keyLength*blockSize*2])) / float64(keyLength*blockSize)
		if currentScore >= theBestScore {
			continue
		}

		theBestScore = currentScore
		theBestKeyLength = keyLength
	}

	return theBestKeyLength
}

func MainSet1Challenge06() {
	println("Break repeating-key XOR")

	content := LoadFileContentAsString("6.txt")
	originalStringBytes, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		panic(fmt.Errorf("some error occured during base64 decode. Error %q", err.Error()))
	}

	keyLength := findKeyLength(originalStringBytes)
	key := make([]byte, keyLength)

	sampleSize := len(originalStringBytes) / int(keyLength)
	for i := byte(0); i < keyLength; i++ {
		column := make([]byte, sampleSize)

		for j := 0; j < sampleSize; j++ {
			column[j] = originalStringBytes[int(i)+j*int(keyLength)]
		}

		_, singleKeyCharacter, _ := CheckAllCombinationOfSingleKey(column)
		key[i] = singleKeyCharacter
	}

	result := EncodeRepeatingXor(originalStringBytes, key)

	println("Key:", string(key))
	println("Message:\n", string(result))
}
