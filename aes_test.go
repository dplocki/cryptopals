package main

import (
	"crypto/aes"
	"testing"
)

func TestEncryptAES128CBC(t *testing.T) {
	planTextMessage := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hea"
	key := []byte("Yellow submarine")
	block, err := aes.NewCipher(key)
	if err != nil {
		panic("cannot create cipher")
	}

	vi, err := GenRandomBytes(16)
	if err != nil {
		panic("cannot crate iv")
	}

	cryptedMessage := make([]byte, len(planTextMessage))
	resultMessage := make([]byte, len(planTextMessage))

	EncryptAES128CBC(block, cryptedMessage, []byte(planTextMessage), vi)
	DecryptAES128CBC(block, resultMessage, cryptedMessage, vi)

	if string(resultMessage) != planTextMessage {
		t.Error("innocorrect result", resultMessage)
	}
}
