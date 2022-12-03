package main

import (
	"aoc2022/day03/common"
	"aoc2022/util"
	"fmt"
)

const inputFile = "../input.txt"

func main() {
	lines := util.ReadFileLinesAsArray(inputFile)

	sum := 0
	for _, line := range lines {
		len := len(line)
		halfLen := len / 2
		firstHalf := line[:halfLen]
		secondHalf := line[halfLen:]
		sameChar := common.FindSameChar(firstHalf, secondHalf)
		sum += common.CharToPriority(sameChar)
	}

	fmt.Println(sum)
}
