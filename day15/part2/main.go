package main

import (
	"aoc2022/day15/common"
	"aoc2022/util"
	"fmt"
)

const inputFile = "../input.txt"

func findEmptySpaces(minCoord int, maxCoord int, sensorReadings []common.SensorReading) []util.Coordinate {
	var emptySpaces []util.Coordinate

	for i := minCoord; i <= maxCoord; i++ {
		searchLine := util.AxisAlignedLine{Pos: i, Min: minCoord, Max: maxCoord}
		lineSegments := common.CalcLineSegmentsForRow(i, sensorReadings)
		lineSegments = lineSegments.Optimize()
		var intersections util.AxisAlignedLines
		for _, ls := range lineSegments {
			if !util.DoIntersect(searchLine, ls) {
				continue
			}
			intersection := util.Intersection(searchLine, ls)
			intersections = append(intersections, intersection)
		}

		var emptySegments util.AxisAlignedLines
		pos := minCoord
		for _, intersection := range intersections {
			if pos == intersection.Min {
				pos = intersection.Max + 1
				continue
			}
			emptySegment := util.AxisAlignedLine{Pos: i, Min: pos, Max: intersection.Min - 1}
			emptySegments = append(emptySegments, emptySegment)
			pos = intersection.Max + 1
		}

		for _, es := range emptySegments {
			for j := es.Min; j <= es.Max; j++ {
				coord := util.Coordinate{X: j, Y: i}
				emptySpaces = append(emptySpaces, coord)
			}
		}
	}
	return emptySpaces
}

func calcTuningFrequency(coord util.Coordinate, multiplier int) int {
	return coord.X*multiplier + coord.Y
}

func main() {
	const minCoord = 0
	const maxCoord = 4000000

	lines := util.ReadFileLinesAsArray(inputFile)
	sensorReadings := common.ParseSensorReadings(lines)
	emptySpaces := findEmptySpaces(minCoord, maxCoord, sensorReadings)
	tuninFreq := calcTuningFrequency(emptySpaces[0], maxCoord)
	fmt.Println(tuninFreq)
}
