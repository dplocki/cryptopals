package main

import (
	"bytes"
	"crypto/aes"
	"strings"
)

func decrypt(encryptedMessage, key []byte) string {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic("cannot create cipher")
	}

	result := make([]byte, len(encryptedMessage))

	DecryptAES128ECB(block, result, encryptedMessage)

	size := len(key)
	length := len(result)
	padLen := int(result[length-1])
	ref := bytes.Repeat([]byte{byte(padLen)}, padLen)
	if padLen > size || padLen == 0 || !bytes.HasSuffix(result, ref) {
		panic("invalid block size")
	}

	return string(result[:length-padLen])
}

func fromQueryParameters(input string) map[string]string {
	result := make(map[string]string)
	input_split := strings.Split(input, "&")

	for _, value := range input_split {
		valueSplit := strings.Split(value, "=")
		result[valueSplit[0]] = valueSplit[1]
	}

	return result
}

func toQueryParameters(input map[string]string) string {
	result := strings.Builder{}
	index := 0

	for key, value := range input {
		key = strings.ReplaceAll(key, "&", "")
		key = strings.ReplaceAll(key, "=", "")

		result.WriteString(key)
		result.WriteRune('=')

		value = strings.ReplaceAll(value, "&", "")
		value = strings.ReplaceAll(value, "=", "")

		result.WriteString(value)
		index++

		if index != len(input) {
			result.WriteRune('&')
		}
	}

	return result.String()
}

func profileFor(email string) string {
	result := make(map[string]string)

	result["email"] = email
	result["uid"] = "10"
	result["role"] = "user"

	return toQueryParameters(result)
}

func prepearAdminAsEmail() string {
	spacesUpToNewBlock := strings.Repeat(" ", 16-len("email="))
	adminBlock := string(AddPaddingToBlock([]byte("admin"), 16))

	return spacesUpToNewBlock + adminBlock
}

func MainSet2Challenge13() {
	println("ECB cut-and-paste")

	key := GenerateAESKey()

	//     1234567890ABCDEF
	// --------------------
	// 0 | email=
	// 0 | email=  foo@bar.
	// 1 | com&uid=10&role=
	// 2 | admin
	// 2 | user

	adminBlock := encrypt([]byte(profileFor(prepearAdminAsEmail())), key)[16:32]
	baseForAtack := encrypt([]byte(profileFor(" foo@bar.com ")), key)

	buffer := make([]byte, 16*3)

	for i := 0; i < 32; i++ {
		buffer[i] = baseForAtack[i]
	}

	for i := 0; i < 16; i++ {
		buffer[32+i] = adminBlock[i]
	}

	result := decrypt(buffer, key)
	queryParamMap := fromQueryParameters(string(result))

	println("{")
	for k, v := range queryParamMap {
		println("\t", k, ":", v)
	}

	println("}")
}
