package main

import "crypto/rand"

const KeySize int = 16

func GenRandomBytes(size int) (bulk []byte, err error) {
	bulk = make([]byte, size)
	_, err = rand.Read(bulk)
	return
}

func GenerateAESKey() []byte {
	key, err := GenRandomBytes(KeySize)
	if err != nil {
		panic("cannot generate key")
	}

	return key
}
