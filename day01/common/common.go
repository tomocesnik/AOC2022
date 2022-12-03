package common

import (
	"log"
	"strconv"
)

func SumLinesToNumbers(lines []string) []int {
	var numbers []int
	var currentNum = 0

	for _, line := range lines {
		if line == "" {
			numbers = append(numbers, currentNum)
			currentNum = 0
		} else {
			num, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			currentNum += num
		}
	}
	return numbers
}
