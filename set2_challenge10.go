package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

func DecryptAES128CBC(block cipher.Block, dst, src []byte) {
	bs := block.BlockSize()

	if len(src)%bs != 0 {
		panic("Need a multiple of the blocksize")
	}

	previous := make([]byte, bs)
	for len(src) > 0 {
		block.Decrypt(dst, src[:bs])

		EncodeFixedXor(dst, dst, previous[:bs])

		previous = src[:bs]
		src = src[bs:]
		dst = dst[bs:]
	}
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

	plaintext := make([]byte, len(originalStringBytes))

	DecryptAES128CBC(block, plaintext, originalStringBytes)

	println(string(plaintext))
}
