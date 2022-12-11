package common

import (
	"aoc2022/util"
	"strings"
)

const numMonkeyLines = 7

type Monkey struct {
	items           []int
	operation       operation
	testDivisor     int
	testTargetTrue  int
	testTargetFalse int
	numInspections  int
}

func (m *Monkey) TestDivisor() int {
	return m.testDivisor
}

type operation struct {
	operator string
	argument string
}

func ParseMonkeys(lines []string) []Monkey {
	var monkeys []Monkey

	lineIdx := 0
	for lineIdx < len(lines) {
		lineStartingItems := strings.TrimPrefix(lines[lineIdx+1], "  Starting items: ")
		lineOperation := strings.TrimPrefix(lines[lineIdx+2], "  Operation: new = old ")
		lineTestDivisor := strings.TrimPrefix(lines[lineIdx+3], "  Test: divisible by ")
		lineTestTargetTrue := strings.TrimPrefix(lines[lineIdx+4], "    If true: throw to monkey ")
		lineTestTargetFalse := strings.TrimPrefix(lines[lineIdx+5], "    If false: throw to monkey ")

		startingItems := parseStartingItems(lineStartingItems)
		operation := parseOperation(lineOperation)
		testDivisor := util.StringToNum(lineTestDivisor)
		testTargetTrue := util.StringToNum(lineTestTargetTrue)
		testTargetFalse := util.StringToNum(lineTestTargetFalse)
		monkeys = append(monkeys, Monkey{startingItems, operation, testDivisor, testTargetTrue, testTargetFalse, 0})

		lineIdx += numMonkeyLines
	}
	return monkeys
}

func parseStartingItems(line string) []int {
	var items []int
	parts := strings.Split(line, ", ")
	for _, p := range parts {
		items = append(items, util.StringToNum(p))
	}
	return items
}

func parseOperation(line string) operation {
	parts := strings.Split(line, " ")
	return operation{parts[0], parts[1]}
}

func DoRounds(monkeys []Monkey, numRounds int, reliefFunc func(int) int) {
	for round := 0; round < numRounds; round++ {
		doRound(monkeys, reliefFunc)
	}
}

func doRound(monkeys []Monkey, reliefFunc func(int) int) {
	for monkeyIdx, monkey := range monkeys {
		for _, item := range monkey.items {
			worryLevel := item
			worryLevel = applyOperation(monkey.operation, worryLevel)
			worryLevel = reliefFunc(worryLevel)

			targetMonkeyIdx := monkey.testTargetFalse
			if (worryLevel % monkey.testDivisor) == 0 {
				targetMonkeyIdx = monkey.testTargetTrue
			}
			monkeys[targetMonkeyIdx].items = append(monkeys[targetMonkeyIdx].items, worryLevel)
		}
		monkey.numInspections += len(monkey.items)
		monkey.items = monkey.items[:0]
		monkeys[monkeyIdx] = monkey
	}
}

func applyOperation(op operation, oldValue int) int {
	val := 0
	if op.argument == "old" {
		val = oldValue
	} else {
		val = util.StringToNum(op.argument)
	}
	switch op.operator {
	case "+":
		return oldValue + val
	case "*":
		return oldValue * val
	}
	return oldValue
}

func CalcMonkeyBusiness(monkeys []Monkey) int {
	var inspections []int
	for _, monkey := range monkeys {
		inspections = append(inspections, monkey.numInspections)
	}
	idx1, max1 := util.MaxNumberInArray(inspections)
	inspections2 := append(inspections[:idx1], inspections[idx1+1:]...)
	_, max2 := util.MaxNumberInArray(inspections2)
	return max1 * max2
}
