package main

import (
	"crypto/aes"
	"math/rand"
	"time"
)

func addRandomPadding(input string) []byte {
	min, max := 5, 10
	leftPadding := rand.Intn(max-min) + min
	message := []byte(input)
	rightPadding := 16 - ((leftPadding + len(message)) % 16)

	result, err := GenRandomBytes(leftPadding + len(message) + rightPadding)
	if err != nil {
		panic("cannot generate random bytes")
	}

	for index, value := range message {
		result[leftPadding+index] = value
	}

	return result
}

func EncryptionOracle(key, message []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic("cannot create cipher")
	}

	result := make([]byte, len(message))
	flip := rand.Int() % 2
	if flip == 0 {
		println("Encrypt by AES128 ECB")
		EncryptAES128ECB(block, result, message)
	} else {
		println("Encrypt by AES128 CBC")
		iv, err := GenRandomBytes(block.BlockSize())
		if err != nil {
			panic("cannot generate iv")
		}

		EncryptAES128CBC(block, result, message, iv)
	}

	return result
}

func MainSet2Challenge11() {
	println("An ECB/CBC detection oracle")

	rand.Seed(time.Now().UnixNano())

	plainText := "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"
	messageWithPadding := addRandomPadding(plainText)
	key := GenerateAESKey()

	encryptedMessage := EncryptionOracle(key, messageWithPadding)

	if IsDecryptedByECB(encryptedMessage, 16) {
		println("Founded ECB")
	} else {
		println("Not founded ECB")
	}
}
