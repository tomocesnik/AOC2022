package main

import (
	"aoc2022/day11/common"
	"aoc2022/util"
	"fmt"
	"math"
)

const inputFile = "../input.txt"

func main() {
	const numRounds = 20
	lines := util.ReadFileLinesAsArray(inputFile)
	monkeys := common.ParseMonkeys(lines)
	reliefFunc := func(worryLevel int) int { return int(math.Floor(float64(worryLevel) / 3.0)) }
	common.DoRounds(monkeys, numRounds, reliefFunc)
	monkeyBusiness := common.CalcMonkeyBusiness(monkeys)
	fmt.Println(monkeyBusiness)
}
