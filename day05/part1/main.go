package main

import (
	"aoc2022/day05/common"
	"aoc2022/util"
	"fmt"
)

const inputFile = "../input.txt"

func main() {
	lines := util.ReadFileLinesAsArray(inputFile)
	stacks, commands := common.ParseInput(lines)
	common.ApplyCommands(stacks, commands, func(targetStack, mvStack common.Stack) common.Stack {
		for i := range mvStack {
			targetStack = append(targetStack, mvStack[len(mvStack)-1-i])
		}
		return targetStack
	})
	stacksTops := common.GetStacksTops(stacks)
	fmt.Println(stacksTops)
}
