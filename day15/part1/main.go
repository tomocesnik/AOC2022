package main

import (
	"aoc2022/day15/common"
	"aoc2022/util"
	"fmt"
)

const inputFile = "../input.txt"

func countFilledCoords(lineSegments util.AxisAlignedLines) int {
	sum := 0
	for _, ls := range lineSegments {
		sum += ls.Max - ls.Min
	}
	return sum
}

func main() {
	const row = 2000000
	lines := util.ReadFileLinesAsArray(inputFile)
	sensorReadings := common.ParseSensorReadings(lines)
	lineSegments := common.CalcLineSegmentsForRow(row, sensorReadings)
	lineSegments = lineSegments.Optimize()
	count := countFilledCoords(lineSegments)
	fmt.Println(count)
}
