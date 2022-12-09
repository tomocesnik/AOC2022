package main

import (
	"aoc2022/day09/common"
	"aoc2022/util"
	"fmt"
)

const inputFile = "../input.txt"

func applyCommands(commands []common.Command) []*common.Position {
	var tailPositions []*common.Position
	headPosition := &common.Position{X: 0, Y: 0}
	tailPosition := &common.Position{X: 0, Y: 0}

	for _, cmd := range commands {
		for i := 0; i < cmd.Amount; i++ {
			headPosition = common.MoveHead(headPosition, cmd.Direction)
			tailPosition = common.MoveLink(tailPosition, headPosition)
			tailPositions = append(tailPositions, tailPosition)
		}
	}
	return tailPositions
}

func main() {
	lines := util.ReadFileLinesAsArray(inputFile)
	commands := common.ParseCommands(lines)
	positions := applyCommands(commands)
	uniquePositions := common.RemoveDuplicateValues(positions)
	fmt.Println(len(uniquePositions))
}
