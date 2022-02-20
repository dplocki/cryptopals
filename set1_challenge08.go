package main

import "encoding/hex"

type ProductResult struct {
	FirstField  int
	SecondField int
}

func product(n int) <-chan ProductResult {
	chanel := make(chan ProductResult)
	go func() {
		for firstIndex := 0; firstIndex < n-1; firstIndex++ {
			for secondIndex := firstIndex + 1; secondIndex < n; secondIndex++ {
				result := new(ProductResult)
				result.FirstField = firstIndex
				result.SecondField = secondIndex

				chanel <- *result
			}
		}

		close(chanel)
	}()

	return chanel
}

func compeareBlock(message []byte, firstIndex, secondIndex, blockSize int) bool {
	for i := 0; i < blockSize; i++ {
		if message[firstIndex*blockSize+i] != message[secondIndex*blockSize+i] {
			return false
		}
	}

	return true
}

func IsDecryptedByECB(message []byte, blockSize int) bool {
	for productResult := range product(len(message) / blockSize) {
		if compeareBlock(message, productResult.FirstField, productResult.SecondField, blockSize) {
			return true
		}
	}

	return false
}

func MainSet1Challenge08() {
	println("Detect AES in ECB mode")

	content := LoadFileContentAsStringsArray("8.txt")

	for lineNumber, value := range content {
		inputByte, _ := hex.DecodeString(value)

		if IsDecryptedByECB(inputByte, 16) {
			println("Line number", lineNumber)
			println(value)
		}
	}
}
