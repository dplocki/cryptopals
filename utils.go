package main

func CompeareBlock(message []byte, firstIndex, secondIndex, blockSize int) bool {
	for i := 0; i < blockSize; i++ {
		if message[firstIndex*blockSize+i] != message[secondIndex*blockSize+i] {
			return false
		}
	}

	return true
}
