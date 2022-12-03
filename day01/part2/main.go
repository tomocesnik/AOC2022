package main

import (
	"aoc2022/day01/common"
	"aoc2022/util"
	"fmt"
)

const inputFile = "../input.txt"

func main() {
	lines := util.ReadFileLinesAsArray(inputFile)
	numbers := common.SumLinesToNumbers(lines)

	var minIdx = 0
	maxNumbers := [3]int{0, 0, 0}

	for _, num := range numbers {
		if num > maxNumbers[minIdx] {
			maxNumbers[minIdx] = num
			minIdx, _ = util.MinNumberInArray(maxNumbers[:])
		}
	}

	sum := 0
	for _, num := range maxNumbers {
		sum += num
	}
	fmt.Println(sum)
}
