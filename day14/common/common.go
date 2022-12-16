package common

import (
	"aoc2022/util"
	"fmt"
	"math"
	"strings"
)

type World map[int]util.AxisAlignedLines

func ParseSolidSegments(lines []string) (util.AxisAlignedLines, util.AxisAlignedLines) {
	var verticalSegments util.AxisAlignedLines
	var horizontalSegments util.AxisAlignedLines

	for _, line := range lines {
		coordsStrs := strings.Split(line, " -> ")
		var coordinates []util.Coordinate
		for _, cs := range coordsStrs {
			coordsParts := strings.Split(cs, ",")
			coordX := util.StringToNum(coordsParts[0])
			coordY := util.StringToNum(coordsParts[1])
			coordinates = append(coordinates, util.Coordinate{X: coordX, Y: coordY})
		}

		for i := 0; i < len(coordinates)-1; i++ {
			coord1 := coordinates[i]
			coord2 := coordinates[i+1]
			if coord1.X == coord2.X {
				min := coord1.Y
				max := coord2.Y
				if coord1.Y > coord2.Y {
					min = coord2.Y
					max = coord1.Y
				}
				verticalSegments = append(verticalSegments, util.AxisAlignedLine{Pos: coord1.X, Min: min, Max: max})
			} else if coord1.Y == coord2.Y {
				min := coord1.X
				max := coord2.X
				if coord1.X > coord2.X {
					min = coord2.X
					max = coord1.X
				}
				horizontalSegments = append(horizontalSegments, util.AxisAlignedLine{Pos: coord1.Y, Min: min, Max: max})
			}
		}
	}

	return verticalSegments, horizontalSegments
}

func ConstructWorld(verticalSegments util.AxisAlignedLines, horizontalSegments util.AxisAlignedLines) World {
	world := make(World)
	for _, vs := range verticalSegments {
		world[vs.Pos] = append(world[vs.Pos], vs)
	}

	for _, hs := range horizontalSegments {
		for i := hs.Min; i <= hs.Max; i++ {
			world[i] = append(world[i], util.AxisAlignedLine{Pos: i, Min: hs.Pos, Max: hs.Pos})
		}
	}

	for k, v := range world {
		world[k] = v.Optimize()
	}
	return world
}

func SimulateSand(sourceCoord util.Coordinate, world World, minSegmentFinder func(util.AxisAlignedLines, int) (*util.AxisAlignedLine, bool)) []util.Coordinate {
	var sandCoords []util.Coordinate
	for {
		// PrintWorld(world, sandCoords)
		restCoord, _ := simulateSandUnit(sourceCoord, world, minSegmentFinder)
		if restCoord == nil {
			break
		}
		// fmt.Println(*restCoord)
		sandCoords = append(sandCoords, *restCoord)
	}
	return sandCoords
}

func simulateSandUnit(coord util.Coordinate, world World, minSegmentFinder func(util.AxisAlignedLines, int) (*util.AxisAlignedLine, bool)) (*util.Coordinate, bool) {
	xvsss := world[coord.X]
	vss, found := minSegmentFinder(xvsss, coord.Y)
	if vss == nil {
		return nil, !found
	}

	restCoordLeft, rDone := simulateSandUnit(util.Coordinate{X: coord.X - 1, Y: vss.Min}, world, minSegmentFinder)
	if rDone {
		return nil, true
	}
	if restCoordLeft != nil {
		return restCoordLeft, false
	}
	restCoordRight, lDone := simulateSandUnit(util.Coordinate{X: coord.X + 1, Y: vss.Min}, world, minSegmentFinder)
	if lDone {
		return nil, true
	}
	if restCoordRight != nil {
		return restCoordRight, false
	}

	// grow segment by 1
	vss.Min = vss.Min - 1
	if !found {
		world[coord.X] = append(world[coord.X], *vss)
	}
	return &util.Coordinate{X: coord.X, Y: vss.Min}, false
}

func PrintWorld(world World, sandCoords []util.Coordinate) {
	minX := math.MaxInt
	maxX := 0
	maxY := 0
	for k, v := range world {
		if k < minX {
			minX = k
		}
		if k > maxX {
			maxX = k
		}
		for _, ss := range v {
			if ss.Max > maxY {
				maxY = ss.Max
			}
		}
	}

	for i := 0; i <= maxY; i++ {
		for j := minX; j <= maxX; j++ {
			ssl := world[j]
			char := "."
			for _, ss := range ssl {
				if (i >= ss.Min) && (i <= ss.Max) {
					char = "#"
				}
			}
			for _, sc := range sandCoords {
				impostorCoord := util.Coordinate{X: j, Y: i}
				if sc == impostorCoord {
					char = "o"
				}
			}
			fmt.Print(char)
		}
		fmt.Println()
	}
}
