package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

func DecryptAES128CBC(block cipher.Block, dst, src []byte) {

}

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

	decripted := make([]byte, len(originalStringBytes))

	DecryptAES128CBC(block, decripted, originalStringBytes)

	println(string(decripted))
}
