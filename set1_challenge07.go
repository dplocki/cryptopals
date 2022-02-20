package main

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
)

func MainSet1Challenge07() {
	println("AES in ECB mode")

	content := LoadFileContentAsString("7.txt")

	originalStringBytes, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		panic(fmt.Errorf("some error occured during base64 decode. Error %q", err.Error()))
	}

	block, err := aes.NewCipher([]byte("YELLOW SUBMARINE"))
	if err != nil {
		panic("cannot create cipher")
	}

	decryptedBytes := make([]byte, len(originalStringBytes))

	DecryptAES128ECB(block, decryptedBytes, originalStringBytes)

	println("The decrypted message:")
	println(string(decryptedBytes))
}
