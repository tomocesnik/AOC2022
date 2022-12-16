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
			if ss.Max > maxY {
				maxY = ss.Max
			}
		}
	}
	return maxY
}

func createMinSegmentFinder(max int) func(util.AxisAlignedLines, int) (*util.AxisAlignedLine, bool) {
	return func(solidSegments util.AxisAlignedLines, minLimit int) (*util.AxisAlignedLine, bool) {
		if minLimit >= max {
			return nil, true
		}
		for idx := range solidSegments {
			ss := &solidSegments[idx]
			if ss.Min > minLimit {
				return ss, true
			}
			if ss.Max >= minLimit {
				return nil, true
			}
		}
		// PosA parameter does not matter anymore at this stage
		return &util.AxisAlignedLine{Pos: 0, Min: max, Max: max}, false
	}
}

func main() {
	sourceCoord := util.Coordinate{X: 500, Y: 0}
	lines := util.ReadFileLinesAsArray(inputFile)
	vSegments, hSegments := common.ParseSolidSegments(lines)
	world := common.ConstructWorld(vSegments, hSegments)
	maxY := findMaxY(world) + 2
	minSegmentFinder := createMinSegmentFinder(maxY)
	sandCoords := common.SimulateSand(sourceCoord, world, minSegmentFinder)
	fmt.Println(len(sandCoords))
}
