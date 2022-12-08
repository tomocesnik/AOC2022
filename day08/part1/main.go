package main

import (
	"aoc2022/day08/common"
	"aoc2022/util"
	"fmt"
)

const inputFile = "../input.txt"

func calcVisibilities1D(getNextEl func() (bool, int, byte), numElems int) []bool {
	visibilities := make([]bool, numElems)
	maxVal := -1
	for valid, elIdx, value := getNextEl(); valid; valid, elIdx, value = getNextEl() {
		intVal := int(value)
		visible := false
		if intVal > maxVal {
			maxVal = intVal
			visible = true
		}
		visibilities[elIdx] = visible
	}
	return visibilities
}

func calcVisibilities2D(grid [][]byte) [][]bool {
	var visibilities [][]bool = make([][]bool, len(grid))
	for row := range visibilities {
		visibilities[row] = make([]bool, len(grid[row]))
	}

	for row := range grid {
		rowLen := len(grid[row])
		visibilitiesRow := calcVisibilities1D(common.IterRow(grid, row), rowLen)
		visibilitiesRowReverse := calcVisibilities1D(common.IterRowReverse(grid, row), rowLen)
		for col := 0; col < rowLen; col++ {
			visibilities[row][col] = visibilities[row][col] || visibilitiesRow[col] || visibilitiesRowReverse[col]
		}
	}

	colLen := len(grid)
	for col := range grid[0] {
		visibilitiesCol := calcVisibilities1D(common.IterCol(grid, col), colLen)
		visibilitiesColReverse := calcVisibilities1D(common.IterColReverse(grid, col), colLen)
		for row := 0; row < colLen; row++ {
			visibilities[row][col] = visibilities[row][col] || visibilitiesCol[row] || visibilitiesColReverse[row]
		}
	}
	return visibilities
}

func countAllVisibles(visibilities [][]bool) int {
	count := 0
	for _, row := range visibilities {
		for _, col := range row {
			if col {
				count++
			}
		}
	}
	return count
}

func main() {
	lines := util.ReadFileLinesAsArray(inputFile)
	grid := common.ParseLinesToByteGrid(lines)
	visibilities := calcVisibilities2D(grid)
	visibles := countAllVisibles(visibilities)
	fmt.Println(visibles)
}
