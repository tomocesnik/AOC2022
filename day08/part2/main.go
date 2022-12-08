package main

import (
	"aoc2022/day08/common"
	"aoc2022/util"
	"fmt"
	"math"
)

const inputFile = "../input.txt"

const numValues = 10

func calcScenicScores1D(getNextEl func() (bool, int, byte), numElems int, maxIdx int) []int {
	idxsForVal := make([]int, numValues)
	for i := range idxsForVal {
		idxsForVal[i] = maxIdx
	}

	scenicScores := make([]int, numElems)
	for valid, elIdx, value := getNextEl(); valid; valid, elIdx, value = getNextEl() {
		intVal := int(value)
		scenicScores[elIdx] = int(math.Abs(float64(elIdx - idxsForVal[intVal])))

		for i := 0; i <= intVal; i++ {
			idxsForVal[i] = elIdx
		}
	}
	return scenicScores
}

func calcScenicScores2D(grid [][]byte) [][]int {
	var scenicScores [][]int = make([][]int, len(grid))
	for row := range scenicScores {
		scenicScores[row] = make([]int, len(grid[row]))
	}

	for row := range grid {
		rowLen := len(grid[row])
		scenicScoresRow := calcScenicScores1D(common.IterRow(grid, row), rowLen, 0)
		scenicScoresRowReverse := calcScenicScores1D(common.IterRowReverse(grid, row), rowLen, rowLen-1)
		for col := 0; col < rowLen; col++ {
			scenicScores[row][col] = scenicScoresRow[col] * scenicScoresRowReverse[col]
		}
	}

	colLen := len(grid)
	for col := range grid[0] {
		scenicScoresCol := calcScenicScores1D(common.IterCol(grid, col), colLen, 0)
		scenicScoresColReverse := calcScenicScores1D(common.IterColReverse(grid, col), colLen, colLen-1)
		for row := 0; row < colLen; row++ {
			scenicScores[row][col] = scenicScores[row][col] * scenicScoresCol[row] * scenicScoresColReverse[row]
		}
	}
	return scenicScores
}

func findMaxScenicScore(scenicScores [][]int) int {
	max := 0
	for _, row := range scenicScores {
		for _, col := range row {
			if col > max {
				max = col
			}
		}
	}
	return max
}

func main() {
	lines := util.ReadFileLinesAsArray(inputFile)
	grid := common.ParseLinesToByteGrid(lines)
	scenicScores := calcScenicScores2D(grid)
	maxScenicScore := findMaxScenicScore(scenicScores)
	fmt.Println(maxScenicScore)
}
