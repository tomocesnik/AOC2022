package main

import (
	"aoc2022/day14/common"
	"aoc2022/util"
	"fmt"
)

const inputFile = "../input.txt"

func createMinSegmentFinder() func(util.AxisAlignedLines, int) (*util.AxisAlignedLine, bool) {
	return func(solidSegments util.AxisAlignedLines, minLimit int) (*util.AxisAlignedLine, bool) {
		for idx := range solidSegments {
			ss := &solidSegments[idx]
			if ss.Min > minLimit {
				return ss, true
			}
			if ss.Max >= minLimit {
				return nil, true
			}
		}
		return nil, false
	}
}

func main() {
	sourceCoord := util.Coordinate{X: 500, Y: 0}
	lines := util.ReadFileLinesAsArray(inputFile)
	vSegments, hSegments := common.ParseSolidSegments(lines)
	world := common.ConstructWorld(vSegments, hSegments)
	minSegmentFinder := createMinSegmentFinder()
	sandCoords := common.SimulateSand(sourceCoord, world, minSegmentFinder)
	fmt.Println(len(sandCoords))
}
