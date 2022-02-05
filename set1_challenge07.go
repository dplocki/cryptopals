package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

func DecryptAES128ECB(block cipher.Block, dst, src []byte) {
	bs := block.BlockSize()

	if len(src)%bs != 0 {
		panic("Need a multiple of the blocksize")
	}

	for len(src) > 0 {
		block.Decrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
}

func MainSet1Challenge07() {
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
