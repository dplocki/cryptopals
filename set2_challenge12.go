package main

import (
	"crypto/aes"
)

func buildString(lenght int) string {
	result := make([]byte, lenght)

	for i := 0; i < lenght; i++ {
		result[i] = 'A'
	}

	return string(result)
}

func encrypt(encryptedMessage string, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic("cannot create cipher")
	}

	buffer := AddPaddingToBlock([]byte(encryptedMessage), len(key))
	result := make([]byte, len(buffer))

	EncryptAES128ECB(block, result, result)

	return result
}

func findKeySize(secretMessageAsBase64 string, key []byte) int {
	originalLeght := len(encrypt(secretMessageAsBase64, key))

	for i := 0; i < 100; i++ {
		lenghtForI := len(encrypt(secretMessageAsBase64+buildString(i), key))

		if lenghtForI > originalLeght {
			return lenghtForI - originalLeght
		}
	}

	panic("cannot found key size")
}

func MainSet2Challenge12() {
	println("Byte-at-a-time ECB decryption (Simple)")

	key := GenerateAESKey()
	secretMessageAsBase64 := HexDecodeToBase64("Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK")

	println("Keysize is: ", findKeySize(secretMessageAsBase64, key))
}
