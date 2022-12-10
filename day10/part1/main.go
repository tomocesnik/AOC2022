package main

import (
	"aoc2022/day10/common"
	"aoc2022/util"
	"fmt"
	"strings"
)

const inputFile = "../input.txt"

func simulateCpu(lines []string) []int {
	regValue := 1
	cycle := 1
	cmdIdx := 0
	var importantSignalVals []int

	for cmdIdx < len(lines) {
		line := lines[cmdIdx]
		cmdIdx++

		moveCycles := common.CyclesNoop
		var afterCycleCb func() = nil

		if strings.HasPrefix(line, common.CmdAddx) {
			num := util.StringToNum(line[len(common.CmdAddx):])
			moveCycles = common.CyclesAddx
			afterCycleCb = func() {
				regValue += num
			}
		}

		for i := 0; i < moveCycles; i++ {
			if isImportantCycle(cycle) {
				importantSignalVals = append(importantSignalVals, regValue*cycle)
			}
			cycle++
		}
		if afterCycleCb != nil {
			afterCycleCb()
		}
	}
	return importantSignalVals
}

func importantCycleIterator() func() int {
	cycle := 20
	return func() int {
		c := cycle
		cycle += 40
		return c
	}
}

func isImportantCycle(cycle int) bool {
	iter := importantCycleIterator()
	importantCycle := iter()

	for importantCycle < cycle {
		importantCycle = iter()
	}
	return cycle == importantCycle
}

func main() {
	lines := util.ReadFileLinesAsArray(inputFile)
	importantSignalVals := simulateCpu(lines)
	sum := util.SumArray(importantSignalVals)
	fmt.Println(sum)
}
