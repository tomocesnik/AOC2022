package common

import (
	"aoc2022/util"
	"fmt"
	"math"
	"sort"
	"strings"
)

type SolidSegment struct {
	PosA int
	MinB int
	MaxB int
}

type Coordinate struct {
	X int
	Y int
}

type SolidSegmentsList []SolidSegment

type World map[int]SolidSegmentsList

func ParseSolidSegments(lines []string) ([]SolidSegment, []SolidSegment) {
	var verticalSegments []SolidSegment
	var horizontalSegments []SolidSegment

	for _, line := range lines {
		coordsStrs := strings.Split(line, " -> ")
		var coordinates []Coordinate
		for _, cs := range coordsStrs {
			coordsParts := strings.Split(cs, ",")
			coordX := util.StringToNum(coordsParts[0])
			coordY := util.StringToNum(coordsParts[1])
			coordinates = append(coordinates, Coordinate{coordX, coordY})
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
				verticalSegments = append(verticalSegments, SolidSegment{coord1.X, min, max})
			} else if coord1.Y == coord2.Y {
				min := coord1.X
				max := coord2.X
				if coord1.X > coord2.X {
					min = coord2.X
					max = coord1.X
				}
				horizontalSegments = append(horizontalSegments, SolidSegment{coord1.Y, min, max})
			}
		}
	}

	return verticalSegments, horizontalSegments
}

func (list SolidSegmentsList) Len() int {
	return len(list)
}

func (list SolidSegmentsList) Less(i, j int) bool {
	if list[i].MinB < list[j].MinB {
		return true
	} else if list[i].MinB == list[j].MinB {
		return list[i].MaxB < list[j].MaxB
	}
	return false
}

func (list SolidSegmentsList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func ConstructWorld(verticalSegments []SolidSegment, horizontalSegments []SolidSegment) World {
	world := make(World)
	for _, vs := range verticalSegments {
		world[vs.PosA] = append(world[vs.PosA], vs)
	}

	for _, hs := range horizontalSegments {
		for i := hs.MinB; i <= hs.MaxB; i++ {
			world[i] = append(world[i], SolidSegment{i, hs.PosA, hs.PosA})
		}
	}

	for k, v := range world {
		sort.Sort(v)

		var ssList SolidSegmentsList
		var ss SolidSegment
		for i, vss := range v {
			if i == 0 {
				ss = vss
				continue
			}

			if doIntersect(ss, vss) {
				ss = union(ss, vss)
			} else {
				ssList = append(ssList, ss)
				ss = vss
			}
		}
		ssList = append(ssList, ss)

		world[k] = ssList
	}
	return world
}

func doIntersect(ss1 SolidSegment, ss2 SolidSegment) bool {
	return (ss1.MaxB >= ss2.MinB) && (ss2.MaxB >= ss1.MinB)
}

func union(ss1 SolidSegment, ss2 SolidSegment) SolidSegment {
	min := int(math.Min(float64(ss1.MinB), float64(ss2.MinB)))
	max := int(math.Max(float64(ss1.MaxB), float64(ss2.MaxB)))
	return SolidSegment{ss1.PosA, min, max}
}

func SimulateSand(sourceCoord Coordinate, world World, minSegmentFinder func(SolidSegmentsList, int) (*SolidSegment, bool)) []Coordinate {
	var sandCoords []Coordinate
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

func simulateSandUnit(coord Coordinate, world World, minSegmentFinder func(SolidSegmentsList, int) (*SolidSegment, bool)) (*Coordinate, bool) {
	xvsss := world[coord.X]
	vss, found := minSegmentFinder(xvsss, coord.Y)
	if vss == nil {
		return nil, !found
	}

	restCoordLeft, rDone := simulateSandUnit(Coordinate{coord.X - 1, vss.MinB}, world, minSegmentFinder)
	if rDone {
		return nil, true
	}
	if restCoordLeft != nil {
		return restCoordLeft, false
	}
	restCoordRight, lDone := simulateSandUnit(Coordinate{coord.X + 1, vss.MinB}, world, minSegmentFinder)
	if lDone {
		return nil, true
	}
	if restCoordRight != nil {
		return restCoordRight, false
	}

	// grow segment by 1
	vss.MinB = vss.MinB - 1
	if !found {
		world[coord.X] = append(world[coord.X], *vss)
	}
	return &Coordinate{coord.X, vss.MinB}, false
}

func PrintWorld(world World, sandCoords []Coordinate) {
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
			if ss.MaxB > maxY {
				maxY = ss.MaxB
			}
		}
	}

	for i := 0; i <= maxY; i++ {
		for j := minX; j <= maxX; j++ {
			ssl := world[j]
			char := "."
			for _, ss := range ssl {
				if (i >= ss.MinB) && (i <= ss.MaxB) {
					char = "#"
				}
			}
			for _, sc := range sandCoords {
				impostorCoord := Coordinate{j, i}
				if sc == impostorCoord {
					char = "o"
				}
			}
			fmt.Print(char)
		}
		fmt.Println()
	}
}
