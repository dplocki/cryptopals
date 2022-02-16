package main

import (
	"crypto/aes"
	"fmt"
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
		print("Encrypt by AES128 ECB")
		EncryptAES128ECB(block, result, message)
	} else {
		print("Encrypt by AES128 CBC")
		iv, err := GenRandomBytes(block.BlockSize())
		if err != nil {
			panic("cannot generate iv")
		}

		EncryptAES128CBC(block, result, message, iv)
	}

	return result
}

func MainSet2Challenge11() {
	rand.Seed(time.Now().UnixNano())

	plainText := "Write a function to generate a random AES key; that's just 16 random bytes."
	messageWithPadding := addRandomPadding(plainText)
	key := GenerateAESKey()

	encryptedMessage := EncryptionOracle(key, messageWithPadding)

	fmt.Println(encryptedMessage)
}
