package main

import (
	"encoding/hex"
)

func MainSet1Challenge03() {
	println("Single-byte XOR cipher")

	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	inputByte, _ := hex.DecodeString(input)

	bestOutput, key, _ := CheckAllCombinationOfSingleKey(inputByte)

	println("Message:", string(bestOutput))
	println("Key:", string(key))
}
