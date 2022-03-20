package main

import (
	"errors"
	"os"
	"strconv"
)

var Solutions = map[int]func(){
	// Set 1
	3: MainSet1Challenge03,
	4: MainSet1Challenge04,
	6: MainSet1Challenge06,
	7: MainSet1Challenge07,
	8: MainSet1Challenge08,

	// Set 2
	10: MainSet2Challenge10,
	11: MainSet2Challenge11,
	12: MainSet2Challenge12,
}

func parseNumberArgument(parameter string) (int, error) {
	if parameter == "last" {
		lastKey := 0
		for key := range Solutions {
			if key > lastKey {
				lastKey = key
			}
		}

		return lastKey, nil
	}

	number, err := strconv.Atoi(parameter)
	if err != nil {
		return 0, err
	}

	_, exists := Solutions[number]
	if exists {
		return number, nil
	}

	return 0, errors.New("solution doesn't exist")
}

func main() {
	if len(os.Args) != 2 {
		println("Please provide the challane number as paramater to program")
		return
	}

	number, err := parseNumberArgument(os.Args[1])
	if err != nil {
		println("Provided argument is not recognise or solution is not exist")
	} else {
		function := Solutions[number]
		function()
	}
}
