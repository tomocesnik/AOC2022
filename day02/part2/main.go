package main

import (
	"aoc2022/day02/common"
	"aoc2022/util"
	"fmt"
	"log"
	"strings"
)

const inputFile = "../input.txt"

func parseSecondColumn(identifier string) common.RpsResult {
	switch identifier {
	case "X":
		return common.Loss
	case "Y":
		return common.Draw
	case "Z":
		return common.Win
	}
	log.Fatal("No mapping for value: " + identifier)
	return common.Loss
}

func calcMoveForResult(mv1 common.RpsMove, res common.RpsResult) common.RpsMove {
	return common.RpsMove((int(res) + int(mv1) - 1 + 3) % 3)
}

func main() {
	lines := util.ReadFileLinesAsArray(inputFile)
	sum := 0
	for _, line := range lines {
		identifiers := strings.Split(line, " ")
		opMove := common.ParseMove(identifiers[0])
		result := parseSecondColumn(identifiers[1])
		meMove := calcMoveForResult(opMove, result)
		sum += meMove.GetScore() + result.GetScore()
	}

	fmt.Println(sum)
}
