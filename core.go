package main

import (
	"encoding/base64"
	"encoding/hex"
)

func HexDecodeToBase64(input string) string {
	input_byte, _ := hex.DecodeString(input)
	result := base64.StdEncoding.EncodeToString(input_byte)

	return result
}

func EncodeFixedXor(input []byte, key []byte) []byte {
	outputByte := make([]byte, len(input))

	for i, value := range input {
		outputByte[i] = value ^ key[i]
	}

	return outputByte
}

func EncodeRepeatingXor(input []byte, key []byte) []byte {
	keySize := len(key)
	outputByte := make([]byte, len(input))

	for i, value := range input {
		outputByte[i] = value ^ key[i%keySize]
	}

	return outputByte
}

func countSetBits(value byte) byte {
	result := byte(0)

	for value > 0 {
		result += value & 1
		value >>= 1
	}

	return result
}

func HammingDistance(firstInput []byte, secondInput []byte) int {
	result := 0

	for i := 0; i < len(firstInput); i++ {
		result += int(countSetBits(firstInput[i] ^ secondInput[i]))
	}

	return result
}
