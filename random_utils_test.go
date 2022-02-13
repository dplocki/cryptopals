package main

import (
	"testing"
)

func TestGenerateAESKey(t *testing.T) {
	key := GenerateAESKey()

	if len(key) != 16 {
		t.Error("innocorrect key size result", len(key))
	}
}
