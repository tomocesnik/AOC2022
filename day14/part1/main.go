package main

import (
	"aoc2022/day14/common"
	"aoc2022/util"
	"fmt"
)

const inputFile = "../input.txt"

func createMinSegmentFinder() func(common.SolidSegmentsList, int) (*common.SolidSegment, bool) {
	return func(solidSegments common.SolidSegmentsList, minLimit int) (*common.SolidSegment, bool) {
		for idx := range solidSegments {
			ss := &solidSegments[idx]
			if ss.MinB > minLimit {
				return ss, true
			}
			if ss.MaxB >= minLimit {
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
