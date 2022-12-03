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
	_, maxNum := util.MaxNumberInArray(numbers)

	fmt.Println(maxNum)
}
