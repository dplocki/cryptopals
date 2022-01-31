package main

import (
	"bufio"
	"fmt"
	"os"
)

func Main_set1_challenge04() {
	file, readError := os.Open("4.txt")
	if readError != nil {
		fmt.Println(readError)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
