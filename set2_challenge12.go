package main

import (
	"crypto/aes"
	"encoding/base64"
	"strings"
)

func buildString(lenght int) string {
	return strings.Repeat("A", lenght)
}

func encrypt(message, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic("cannot create cipher")
	}

	buffer := AddPaddingToBlock(message, len(key))
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

func findLetterForBlock(mask string, block, key []byte) rune {
	compeareSize := len(mask)
	for letter := rune(0); letter <= 255; letter++ {
		blockForLetter := encrypt([]byte(mask+string(letter)), key)

		if compeare(block, blockForLetter, compeareSize) {
			return letter
		}
	}

	panic("coudn't find the letter")
}

func byteAtTimeDecryption(secretMessage, key []byte, blockSize int) string {
	result := strings.Builder{}

	for {
		for index := 1; index <= blockSize; index++ {
			prefixPadding := buildString(blockSize - index)
			currentBase := prefixPadding + result.String()
			encryptedMessage := encrypt([]byte(prefixPadding+string(secretMessage)), key)
			letter := findLetterForBlock(currentBase, encryptedMessage, key)

			if result.Len() == len(secretMessage) {
				return result.String()
			}

			result.WriteRune(letter)
		}
	}
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

	brakeEncryptedMessage := byteAtTimeDecryption(secretMessageAsBase64, key, keySize)
	if brakeEncryptedMessage != string(secretMessageAsBase64) {
		panic("Unable to decrypt message")
	}

	println("Decrypted message:", brakeEncryptedMessage)
}
