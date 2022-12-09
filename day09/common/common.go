package common

import (
	"aoc2022/util"
	"math"
	"strings"
)

type Command struct {
	Direction string
	Amount    int
}

type Position struct {
	X int
	Y int
}

func ParseCommands(lines []string) []Command {
	var commands []Command
	for _, line := range lines {
		parts := strings.Split(line, " ")
		commands = append(commands, Command{parts[0], util.StringToNum(parts[1])})
	}
	return commands
}

func MoveHead(pos *Position, direction string) *Position {
	switch direction {
	case "U":
		return &Position{pos.X, pos.Y + 1}
	case "D":
		return &Position{pos.X, pos.Y - 1}
	case "L":
		return &Position{pos.X - 1, pos.Y}
	case "R":
		return &Position{pos.X + 1, pos.Y}
	}
	return pos
}

func MoveLink(pos *Position, headPos *Position) *Position {
	if pos.X == headPos.X {
		diff := headPos.Y - pos.Y
		if math.Abs(float64(diff)) >= 2 {
			return &Position{X: pos.X, Y: pos.Y + numToUnit(diff)}
		}
	} else if pos.Y == headPos.Y {
		diff := headPos.X - pos.X
		if math.Abs(float64(diff)) >= 2 {
			return &Position{X: pos.X + numToUnit(diff), Y: pos.Y}
		}
	} else {
		diffX := headPos.X - pos.X
		diffY := headPos.Y - pos.Y
		if (math.Abs(float64(diffX)) >= 2) || (math.Abs(float64(diffY)) >= 2) {
			return &Position{X: pos.X + numToUnit(diffX), Y: pos.Y + numToUnit(diffY)}
		}
	}
	return pos
}

func numToUnit(num int) int {
	if num > 0 {
		return 1
	}
	return -1
}

func RemoveDuplicateValues(positions []*Position) []*Position {
	keys := make(map[Position]bool)
	list := []*Position{}

	for _, entry := range positions {
		if _, value := keys[*entry]; !value {
			keys[*entry] = true
			list = append(list, entry)
		}
	}
	return list
}
