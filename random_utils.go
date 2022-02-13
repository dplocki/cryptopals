package main

import "crypto/rand"

func GenerateAESKey() []byte {
	key := make([]byte, 16)
	_, err := rand.Read(key)
	if err != nil {
		panic("cannot generate key")
	}

	return key
}
