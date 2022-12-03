package main

import (
	"aoc2022/day02/common"
	"aoc2022/util"
	"fmt"
	"log"
	"strings"
)

const inputFile = "../input.txt"

func parseSecondColumn(identifier string) common.RpsMove {
	switch identifier {
	case "X":
		return common.Rock
	case "Y":
		return common.Paper
	case "Z":
		return common.Scissors
	}
	log.Fatal("No mapping for value: " + identifier)
	return common.Rock
}

func calcRoundResult(mv1 common.RpsMove, mv2 common.RpsMove) common.RpsResult {
	return common.RpsResult((int(mv2) - int(mv1) + 1 + 3) % 3)
}

func main() {
	lines := util.ReadFileLinesAsArray(inputFile)
	sum := 0
	for _, line := range lines {
		identifiers := strings.Split(line, " ")
		opMove := common.ParseMove(identifiers[0])
		meMove := parseSecondColumn(identifiers[1])
		result := calcRoundResult(opMove, meMove)
		sum += meMove.GetScore() + result.GetScore()
	}

	fmt.Println(sum)
}
