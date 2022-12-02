package main

import (
	util01 "aoc2022/day01/util"
	util02 "aoc2022/day02/util"
	"fmt"
	"log"
	"strings"
)

const inputFile = "../input.txt"

func parseSecondColumn(identifier string) util02.RpsResult {
	switch identifier {
	case "X":
		return util02.Loss
	case "Y":
		return util02.Draw
	case "Z":
		return util02.Win
	}
	log.Fatal("No mapping for value: " + identifier)
	return util02.Loss
}

func calcMoveForResult(mv1 util02.RpsMove, res util02.RpsResult) util02.RpsMove {
	return util02.RpsMove((int(res) + int(mv1) - 1 + 3) % 3)
}

func main() {
	lines := util01.ReadFileLinesAsArray(inputFile)
	sum := 0
	for _, line := range lines {
		identifiers := strings.Split(line, " ")
		opMove := util02.ParseMove(identifiers[0])
		result := parseSecondColumn(identifiers[1])
		meMove := calcMoveForResult(opMove, result)
		sum += meMove.GetScore() + result.GetScore()
	}

	fmt.Println(sum)
}
