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

func EncodeRepeatingXor(input string, key string) string {
	inputByte := []byte(input)
	inputKey := []byte(key)
	keySize := len(inputKey)
	outputByte := make([]byte, len(inputByte))

	for i, value := range inputByte {
		outputByte[i] = value ^ inputKey[i%keySize]
	}

	result := hex.EncodeToString(outputByte)
	return result
}
