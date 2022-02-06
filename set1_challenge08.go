package main

import (
	"encoding/hex"
	"fmt"
)

func MainSet1Challenge08() {
	content := LoadFileContentAsStringsArray("8.txt")

	for _, v := range content {
		inputByte, _ := hex.DecodeString(v)
		fmt.Println(inputByte[0:16])
	}
}
