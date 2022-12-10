package main

import (
	"aoc2022/day10/common"
	"aoc2022/util"
	"fmt"
	"strings"
)

const inputFile = "../input.txt"

const terminalWidth = 40

func simulateCpu(lines []string) []strings.Builder {
	regValue := 1
	cycle := 1
	cmdIdx := 0
	var strBuffers []strings.Builder
	var strBuffer *strings.Builder

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
			charToWrite := "."
			terminalPos := (cycle - 1) % terminalWidth
			if (regValue-1 == terminalPos) || (regValue == terminalPos) || (regValue+1 == terminalPos) {
				charToWrite = "#"
			}
			if terminalPos == 0 {
				strBuffers = append(strBuffers, strings.Builder{})
				strBuffer = &strBuffers[len(strBuffers)-1]
			}
			strBuffer.WriteString(charToWrite)
			cycle++
		}
		if afterCycleCb != nil {
			afterCycleCb()
		}
	}
	return strBuffers
}

func main() {
	lines := util.ReadFileLinesAsArray(inputFile)
	strBuffers := simulateCpu(lines)
	for _, strBuffer := range strBuffers {
		fmt.Println(strBuffer.String())
	}
}
