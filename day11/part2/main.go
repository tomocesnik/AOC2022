package main

import (
	"aoc2022/day11/common"
	"aoc2022/util"
	"fmt"
)

const inputFile = "../input.txt"

func calcCommonDivisor(monkeys []common.Monkey) int {
	multDivisors := 1
	for _, m := range monkeys {
		multDivisors *= m.TestDivisor()
	}
	return multDivisors
}

func main() {
	const numRounds = 10000
	lines := util.ReadFileLinesAsArray(inputFile)
	monkeys := common.ParseMonkeys(lines)
	commonDiv := calcCommonDivisor(monkeys)
	reliefFunc := func(worryLevel int) int { return worryLevel % commonDiv }
	common.DoRounds(monkeys, numRounds, reliefFunc)
	monkeyBusiness := common.CalcMonkeyBusiness(monkeys)
	fmt.Println(monkeyBusiness)
}
