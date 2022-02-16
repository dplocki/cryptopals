package main

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
)

func MainSet2Challenge10() {
	content := LoadFileContentAsString("10.txt")

	originalStringBytes, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		panic(fmt.Errorf("some error occured during base64 decode. Error %q", err.Error()))
	}

	block, err := aes.NewCipher([]byte("YELLOW SUBMARINE"))
	if err != nil {
		panic("cannot create cipher")
	}

	plaintext := make([]byte, len(originalStringBytes))

	iv := make([]byte, aes.BlockSize)
	DecryptAES128CBC(block, plaintext, originalStringBytes, iv)

	println(string(plaintext))
}
