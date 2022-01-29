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
