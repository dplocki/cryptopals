package main

import (
	"crypto/rand"
)

func GenRandomBytes(size int) (bulk []byte, err error) {
	bulk = make([]byte, size)
	_, err = rand.Read(bulk)
	return
}

func GenerateAESKey() []byte {
	key, err := GenRandomBytes(16)
	if err != nil {
		panic("cannot generate key")
	}

	return key
}
