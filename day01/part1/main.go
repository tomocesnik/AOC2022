package main

import (
	"aoc2022/day01/util"
	"fmt"
)

const inputFile = "../input.txt"

func main() {
	lines := util.ReadFileLinesAsArray(inputFile)
	numbers := util.SumLinesToNumbers(lines)
	_, maxNum := util.MaxNumberInArray(numbers)

	fmt.Println(maxNum)
}
