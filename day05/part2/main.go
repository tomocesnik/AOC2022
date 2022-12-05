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
		return append(targetStack, mvStack...)
	})
	stacksTops := common.GetStacksTops(stacks)
	fmt.Println(stacksTops)
}
