package common

import (
	"aoc2022/util"
	"strings"
)

type SensorReading struct {
	coordSensor util.Coordinate
	coordBeacon util.Coordinate
}

func ParseSensorReadings(lines []string) []SensorReading {
	var sensorReadings []SensorReading
	for _, line := range lines {
		lineParts := strings.Split(line, ": ")
		coordStrSensor := strings.TrimPrefix(lineParts[0], "Sensor at ")
		coordSensor := parseCoordStr(coordStrSensor)
		coordStrBeacon := strings.TrimPrefix(lineParts[1], "closest beacon is at ")
		coordBeacon := parseCoordStr(coordStrBeacon)
		sensorReadings = append(sensorReadings, SensorReading{coordSensor: coordSensor, coordBeacon: coordBeacon})
	}
	return sensorReadings
}

func parseCoordStr(coordStr string) util.Coordinate {
	parts := strings.Split(coordStr, ", ")
	xStr := strings.TrimPrefix(parts[0], "x=")
	x := util.StringToNum(xStr)
	yStr := strings.TrimPrefix(parts[1], "y=")
	y := util.StringToNum(yStr)
	return util.Coordinate{X: x, Y: y}
}

func CalcLineSegmentsForRow(row int, sensorReadings []SensorReading) util.AxisAlignedLines {
	var lineSegments util.AxisAlignedLines
	for _, sr := range sensorReadings {
		mDist := util.ManhattanDistance(sr.coordSensor, sr.coordBeacon)
		minY := sr.coordSensor.Y - mDist
		maxY := sr.coordSensor.Y + mDist
		if (row < minY) || (row > maxY) {
			continue
		}
		distMin := row - minY
		distMax := maxY - row
		dist := distMin
		if distMax < distMin {
			dist = distMax
		}
		lineSegments = append(lineSegments, util.AxisAlignedLine{Pos: row, Min: sr.coordSensor.X - dist, Max: sr.coordSensor.X + dist})
	}
	return lineSegments
}
