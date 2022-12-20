package main

import (
	"aoc2022/util"
	"fmt"
	"strings"
)

const inputFile = "../input.txt"

type point struct {
	x int
	y int
	z int
}

func parsePoints(lines []string) []point {
	points := make([]point, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, ",")
		points[i] = point{x: util.StringToNum(parts[0]), y: util.StringToNum(parts[1]), z: util.StringToNum(parts[2])}
	}
	return points
}

func countExposedSides(points []point) int {
	pointsExposedSides := initExposedSidesMap(points)
	var xmap map[int][]*point = make(map[int][]*point)
	var ymap map[int][]*point = make(map[int][]*point)

	for i := range points {
		pnt := &points[i]
		neighborsX := findNeighbors(pnt.x, xmap, pnt.y, pnt.z, func(p *point) (int, int) {
			return p.y, p.z
		})
		neighborsY := findNeighbors(pnt.y, ymap, pnt.x, pnt.z, func(p *point) (int, int) {
			return p.x, p.z
		})

		var neighbors []*point
		neighbors = append(neighbors, neighborsX...)
		for _, n := range neighborsY {
			existing := false
			for _, en := range neighbors {
				if n == en {
					existing = true
				}
			}
			if !existing {
				neighbors = append(neighbors, n)
			}
		}

		pointsExposedSides[pnt] -= len(neighbors)
		for _, n := range neighbors {
			pointsExposedSides[n] -= 1
		}

		xmap[pnt.x] = append(xmap[pnt.x], pnt)
		ymap[pnt.y] = append(ymap[pnt.y], pnt)
	}

	sum := 0
	for _, v := range pointsExposedSides {
		sum += v
	}
	return sum
}

func initExposedSidesMap(points []point) map[*point]int {
	var pointsExposedSides map[*point]int = make(map[*point]int)
	for i := range points {
		point := &points[i]
		pointsExposedSides[point] = 6
	}
	return pointsExposedSides
}

func findNeighbors(key int, coordsMap map[int][]*point, a int, b int, pointCoordsExtractor func(*point) (int, int)) []*point {
	var neighbors []*point
	pointsWithSameCoord := coordsMap[key]
	for _, p := range pointsWithSameCoord {
		pa, pb := pointCoordsExtractor(p)
		if (a == pa) && ((b+1 == pb) || (b-1 == pb)) {
			neighbors = append(neighbors, p)
		}
		if (b == pb) && ((a+1 == pa) || (a-1 == pa)) {
			neighbors = append(neighbors, p)
		}
	}
	return neighbors
}

func main() {
	lines := util.ReadFileLinesAsArray(inputFile)
	points := parsePoints(lines)
	exposedSides := countExposedSides(points)
	fmt.Println(exposedSides)
}
