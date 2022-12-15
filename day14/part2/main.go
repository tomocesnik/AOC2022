package main

import (
	"aoc2022/day14/common"
	"aoc2022/util"
	"fmt"
)

const inputFile = "../input.txt"

func findMaxY(world common.World) int {
	maxY := 0
	for _, v := range world {
		for _, ss := range v {
			if ss.MaxB > maxY {
				maxY = ss.MaxB
			}
		}
	}
	return maxY
}

func createMinSegmentFinder(max int) func(common.SolidSegmentsList, int) (*common.SolidSegment, bool) {
	return func(solidSegments common.SolidSegmentsList, minLimit int) (*common.SolidSegment, bool) {
		if minLimit >= max {
			return nil, true
		}
		for idx := range solidSegments {
			ss := &solidSegments[idx]
			if ss.MinB > minLimit {
				return ss, true
			}
			if ss.MaxB >= minLimit {
				return nil, true
			}
		}
		// PosA parameter does not matter anymore at this stage
		return &common.SolidSegment{PosA: 0, MinB: max, MaxB: max}, false
	}
}

func main() {
	sourceCoord := common.Coordinate{X: 500, Y: 0}
	lines := util.ReadFileLinesAsArray(inputFile)
	vSegments, hSegments := common.ParseSolidSegments(lines)
	world := common.ConstructWorld(vSegments, hSegments)
	maxY := findMaxY(world) + 2
	minSegmentFinder := createMinSegmentFinder(maxY)
	sandCoords := common.SimulateSand(sourceCoord, world, minSegmentFinder)
	fmt.Println(len(sandCoords))
}
