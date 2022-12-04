package common

import (
	"aoc2022/util"
	"strings"
)

type SectionsAssignment struct {
	Min int
	Max int
}

type SectionsAssignmentPair [2]SectionsAssignment

func FileAsSectionsAssignmentPairs(lines []string) []SectionsAssignmentPair {
	var saPairs []SectionsAssignmentPair

	for _, line := range lines {
		sectionsAssignments := strings.Split(line, ",")
		sa1raw := strings.Split(sectionsAssignments[0], "-")
		sa2raw := strings.Split(sectionsAssignments[1], "-")

		sa1 := SectionsAssignment{util.StringToNum(sa1raw[0]), util.StringToNum(sa1raw[1])}
		sa2 := SectionsAssignment{util.StringToNum(sa2raw[0]), util.StringToNum(sa2raw[1])}
		saPairs = append(saPairs, SectionsAssignmentPair{sa1, sa2})
	}
	return saPairs
}
