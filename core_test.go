package main

import (
	"encoding/hex"
	"testing"
)

func TestHexDecodeToBase64(t *testing.T) {
	result := HexDecodeToBase64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")

	if result != "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t" {
		t.Error("innocorrect result", result)
	}
}

func TestEncodeFixedXor(t *testing.T) {
	input, _ := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	key, _ := hex.DecodeString("686974207468652062756c6c277320657965")

	result := EncodeFixedXor(input, key)

	if hex.EncodeToString(result) != "746865206b696420646f6e277420706c6179" {
		t.Error("innocorrect result", result)
	}
}

func TestEncodeRepeatingXor(t *testing.T) {
	inputByte := []byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal")
	inputKey := []byte("ICE")

	result := EncodeRepeatingXor(inputByte, inputKey)

	if hex.EncodeToString(result) != "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f" {
		t.Error("innocorrect result", result)
	}
}
