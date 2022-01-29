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

func EncodeFixedXor(input string, key string) string {
	input_byte, _ := hex.DecodeString(input)
	key_byte, _ := hex.DecodeString(key)
	output_byte := make([]byte, len(input_byte))

	for i, value := range input_byte {
		output_byte[i] = value ^ key_byte[i]
	}

	result := hex.EncodeToString(output_byte)
	return result
}
