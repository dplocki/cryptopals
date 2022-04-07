package main

import "math/rand"

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func MainSet2Challenge14() {
	println("Byte-at-a-time ECB decryption (Harder)")

	key := GenerateAESKey()
	secretSuffix := RandomString(rand.Int())
	randomPrefix := RandomString(rand.Int())
	attackPattern := ""

	encrypted := encrypt([]byte(secretSuffix+attackPattern+randomPrefix), key)

	println(encrypted)
}
