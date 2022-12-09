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

	var knotsPositions [8]*common.Position
	numKnots := len(knotsPositions)
	for i := 0; i < numKnots; i++ {
		knotsPositions[i] = headPosition
	}
	tailPosition := headPosition

	for _, cmd := range commands {
		for i := 0; i < cmd.Amount; i++ {
			headPosition = common.MoveHead(headPosition, cmd.Direction)
			for i := 0; i < numKnots; i++ {
				fwdLink := headPosition
				if i > 0 {
					fwdLink = knotsPositions[i-1]
				}
				knotsPositions[i] = common.MoveLink(knotsPositions[i], fwdLink)
			}
			tailPosition = common.MoveLink(tailPosition, knotsPositions[numKnots-1])
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
