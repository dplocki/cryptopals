package main

import (
	"crypto/aes"
	"encoding/base64"
	"strings"
)

func buildString(lenght int) string {
	result := make([]byte, lenght)

	for i := 0; i < lenght; i++ {
		result[i] = 'A'
	}

	return string(result)
}

func encrypt(encryptedMessage, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic("cannot create cipher")
	}

	buffer := AddPaddingToBlock(encryptedMessage, len(key))
	result := make([]byte, len(buffer))

	EncryptAES128ECB(block, result, buffer)

	return result
}

func findKeySize(secretMessageAsBase64, key []byte) int {
	originalLeght := len(encrypt(secretMessageAsBase64, key))

	for i := 0; i < 100; i++ {
		lenghtForI := len(encrypt([]byte(string(secretMessageAsBase64)+buildString(i)), key))

		if lenghtForI > originalLeght {
			return lenghtForI - originalLeght
		}
	}

	panic("cannot found key size")
}

func compeare(firstBlock, secondBlock []byte, blockSize int) bool {
	for i := 0; i < blockSize; i++ {
		if firstBlock[i] != secondBlock[i] {
			return false
		}
	}

	return true
}

func findLetterForBlock(mask string, block, key []byte, blockSize int) rune {
	for letter := rune(0); letter <= 255; letter++ {
		blockForLetter := encrypt([]byte(mask+string(letter)), key)

		if compeare(block, blockForLetter, blockSize) {
			return letter
		}
	}

	panic("coudn't find the letter")
}

func byteAtTimeDecryption(secretMessage, key []byte, blockSize int) string {
	result := strings.Builder{}

	for index := 1; index <= blockSize; index++ {
		currentBase := buildString(blockSize-index) + string(secretMessage[:index-1])
		firstBlock := currentBase + string(secretMessage[index-1])
		firstBlockEncrypted := encrypt([]byte(firstBlock), key)
		letter := findLetterForBlock(currentBase, firstBlockEncrypted, key, blockSize)

		result.WriteRune(letter)
	}

	return result.String()
}

func MainSet2Challenge12() {
	println("Byte-at-a-time ECB decryption (Simple)")

	key := GenerateAESKey()
	secretMessageAsBase64, _ := base64.StdEncoding.DecodeString("Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK")

	keySize := findKeySize(secretMessageAsBase64, key)
	println("Keysize is:", keySize)

	if IsDecryptedByECB(encrypt([]byte(buildString(keySize*2)+string(secretMessageAsBase64)), key), keySize) {
		println("Founded ECB")
	} else {
		panic("Not founded ECB")
	}

	println("Decrypted message:", byteAtTimeDecryption(secretMessageAsBase64, key, keySize))
}
