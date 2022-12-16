package util

import (
	"bufio"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

func ReadFileLinesAsArray(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	var fileLines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	file.Close()
	return fileLines
}

func StringToNum(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func MinNumberInArray(arr []int) (int, int) {
	var minNum = math.MaxInt
	var idx = 0
	for i, num := range arr {
		if num < minNum {
			minNum = num
			idx = i
		}
	}
	return idx, minNum
}

func MaxNumberInArray(arr []int) (int, int) {
	var maxNum = math.MinInt
	var idx = 0
	for i, num := range arr {
		if num > maxNum {
			maxNum = num
			idx = i
		}
	}
	return idx, maxNum
}

func SumArray(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}

type Coordinate struct {
	X int
	Y int
}

func ManhattanDistance(c1 Coordinate, c2 Coordinate) int {
	return int(math.Abs(float64(c1.X-c2.X)) + math.Abs(float64(c1.Y-c2.Y)))
}

type AxisAlignedLine struct {
	Pos int
	Min int
	Max int
}

type AxisAlignedLines []AxisAlignedLine

func (aals AxisAlignedLines) Sort() {
	sort.Slice(aals, axisAlignedLinesComparator(aals))
}

func axisAlignedLinesComparator(aals AxisAlignedLines) func(i, j int) bool {
	return func(i, j int) bool {
		if aals[i].Min < aals[j].Min {
			return true
		} else if aals[i].Min == aals[j].Min {
			return aals[i].Max < aals[j].Max
		}
		return false
	}
}

func (aals AxisAlignedLines) Optimize() AxisAlignedLines {
	aalsCopy := aals
	aalsCopy.Sort()

	var optimizedAals AxisAlignedLines
	var currentAal AxisAlignedLine
	for i, aal := range aalsCopy {
		if i == 0 {
			currentAal = aal
			continue
		}

		if DoIntersect(currentAal, aal) {
			currentAal = Union(currentAal, aal)
		} else {
			optimizedAals = append(optimizedAals, currentAal)
			currentAal = aal
		}
	}
	optimizedAals = append(optimizedAals, currentAal)
	return optimizedAals
}

func DoIntersect(aal1 AxisAlignedLine, aal2 AxisAlignedLine) bool {
	return (aal1.Max >= aal2.Min) && (aal2.Max >= aal1.Min)
}

func Union(aal1 AxisAlignedLine, aal2 AxisAlignedLine) AxisAlignedLine {
	min := int(math.Min(float64(aal1.Min), float64(aal2.Min)))
	max := int(math.Max(float64(aal1.Max), float64(aal2.Max)))
	return AxisAlignedLine{aal1.Pos, min, max}
}

func Intersection(aal1 AxisAlignedLine, aal2 AxisAlignedLine) AxisAlignedLine {
	min := int(math.Max(float64(aal1.Min), float64(aal2.Min)))
	max := int(math.Min(float64(aal1.Max), float64(aal2.Max)))
	return AxisAlignedLine{aal1.Pos, min, max}
}
