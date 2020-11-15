package main

import (
	"fmt"
	"os"
)

// Function takes an array of inputs, the values for the array should be between 0 < value < 100
// Counts how many times a value is repeated. If two Counts are equal, return the count of the
// lowest value

func printUniqueValue(arr []int) string {
	//Create a dictionary of values for each element
	var dict = make(map[int]int)
	for _, num := range arr {
		if (num >= 0) && (num <= 100) {
			dict[num] = dict[num] + 1
		} else {
			fmt.Print("Error, array value out of range!")
			os.Exit(1)
		}
	}
	var result string
	currentNumber, currentCount := arr[0], dict[arr[0]]
	for number, count := range dict {
		fmt.Printf("[%d]: [%d]\n", currentNumber, currentCount)
		if number < currentNumber {
			if count >= currentCount {
				result = fmt.Sprintf("number [%d] repeated [%d] times\n", number, count)
			}
		}

		currentNumber = number
		if count >= currentCount {
			currentCount = count
		}
	}
	return result
}
