package main

import (
	util01 "aoc2022/day01/util"
	util02 "aoc2022/day02/util"
	"fmt"
	"log"
	"strings"
)

const inputFile = "../input.txt"

func parseSecondColumn(identifier string) util02.RpsMove {
	switch identifier {
	case "X":
		return util02.Rock
	case "Y":
		return util02.Paper
	case "Z":
		return util02.Scissors
	}
	log.Fatal("No mapping for value: " + identifier)
	return util02.Rock
}

func calcRoundResult(mv1 util02.RpsMove, mv2 util02.RpsMove) util02.RpsResult {
	return util02.RpsResult((int(mv2) - int(mv1) + 1 + 3) % 3)
}

func main() {
	lines := util01.ReadFileLinesAsArray(inputFile)
	sum := 0
	for _, line := range lines {
		identifiers := strings.Split(line, " ")
		opMove := util02.ParseMove(identifiers[0])
		meMove := parseSecondColumn(identifiers[1])
		result := calcRoundResult(opMove, meMove)
		sum += meMove.GetScore() + result.GetScore()
	}

	fmt.Println(sum)
}
