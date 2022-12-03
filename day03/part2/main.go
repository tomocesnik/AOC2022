package main

import (
	"aoc2022/day03/common"
	"aoc2022/util"
	"fmt"
)

const inputFile = "../input.txt"

func findSameChars(string1 string, string2 string) string {
	var sameChars string
	for _, c1 := range string1 {
		for _, c2 := range string2 {
			if c1 == c2 {
				sameChars += string(c1)
			}
		}
	}
	return sameChars
}

func findSameChar(eg elfGroup) byte {
	sameChars := findSameChars(eg[0], eg[1])
	return common.FindSameChar(sameChars, eg[2])
}

type elfGroup [3]string

func fileAsElfGroups(lines []string) []elfGroup {
	groupIdx := 0
	currentElfGroup := elfGroup{"", "", ""}
	var elfGroups []elfGroup
	for _, line := range lines {
		currentElfGroup[groupIdx] = line
		groupIdx = (groupIdx + 1) % 3
		if groupIdx == 0 {
			elfGroups = append(elfGroups, currentElfGroup)
			currentElfGroup = elfGroup{"", "", ""}
		}
	}
	return elfGroups
}

func main() {
	lines := util.ReadFileLinesAsArray(inputFile)
	elfGroups := fileAsElfGroups(lines)

	sum := 0
	for _, elfGroup := range elfGroups {
		sameChar := findSameChar(elfGroup)
		sum += common.CharToPriority(sameChar)
	}

	fmt.Println(sum)
}
