package main

import (
	"aoc2022/day21/common"
	"aoc2022/util"
	"fmt"
)

const inputFile = "../input.txt"

func main() {
	lines := util.ReadFileLinesAsArray(inputFile)
	expressions := common.ParseExpressions(lines)
	value := expressions["root"].CalcValue(expressions)
	fmt.Println(value)
}
