package main

import (
	"aoc2022/day04/common"
	"aoc2022/util"
	"fmt"
)

const inputFile = "../input.txt"

func saPartlyContainedIn(sa1 common.SectionsAssignment, sa2 common.SectionsAssignment) bool {
	return ((sa1.Min >= sa2.Min) && (sa1.Min <= sa2.Max)) || ((sa1.Max >= sa2.Min) && (sa1.Max <= sa2.Max))
}

func main() {
	lines := util.ReadFileLinesAsArray(inputFile)

	sum := 0
	saPairs := common.FileAsSectionsAssignmentPairs(lines)
	for _, saPair := range saPairs {
		if saPartlyContainedIn(saPair[0], saPair[1]) || saPartlyContainedIn(saPair[1], saPair[0]) {
			sum++
		}
	}

	fmt.Println(sum)
}
