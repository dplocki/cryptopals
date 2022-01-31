package main

import (
	"fmt"
)

func MainSet1Challenge03() {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	bestOutput, key, _ := CheckAllCombinationOfSingleKey(input)

	fmt.Println("Message:", bestOutput)
	fmt.Println("Key:", key)
}
