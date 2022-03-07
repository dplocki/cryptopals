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

func TestAddPaddingToBlock(t *testing.T) {
	input := []byte("YELLOW SUBMARINE")
	result := AddPaddingToBlock(input, 20)

	if len(result) != 20 {
		t.Error("innocorrect size of result", result)
	}

	if string(result) != "YELLOW SUBMARINE\x04\x04\x04\x04" {
		t.Error("innocorrect result", result)
	}
}

func TestAddPaddingToBlock_OverBlockSize(t *testing.T) {
	input := []byte("YELLOW SUBMARINE|YELLOW SUBMARINE")
	result := AddPaddingToBlock(input, 20)

	if len(result) != 40 {
		t.Error("innocorrect size of result", result)
	}

	if string(result) != "YELLOW SUBMARINE|YELLOW SUBMARINE\x07\x07\x07\x07\x07\x07\x07" {
		t.Error("innocorrect result", result)
	}
}
