package common

import (
	"aoc2022/util"
	"strings"
)

func ParseLinesToByteGrid(lines []string) [][]byte {
	var grid [][]byte
	for _, line := range lines {
		lineParts := strings.Split(line, "")
		gridLine := make([]byte, len(lineParts))
		for i, lp := range lineParts {
			gridLine[i] = byte(util.StringToNum(lp))
		}
		grid = append(grid, gridLine)
	}
	return grid
}

func IterRow(grid [][]byte, rowIdx int) func() (bool, int, byte) {
	elIdx := 0
	return func() (bool, int, byte) {
		if elIdx >= len(grid[rowIdx]) {
			return false, elIdx, 0
		}
		el := grid[rowIdx][elIdx]
		idx := elIdx
		elIdx++
		return true, idx, el
	}
}

func IterRowReverse(grid [][]byte, rowIdx int) func() (bool, int, byte) {
	elIdx := len(grid[rowIdx]) - 1
	return func() (bool, int, byte) {
		if elIdx < 0 {
			return false, elIdx, 0
		}
		el := grid[rowIdx][elIdx]
		idx := elIdx
		elIdx--
		return true, idx, el
	}
}

func IterCol(grid [][]byte, colIdx int) func() (bool, int, byte) {
	elIdx := 0
	return func() (bool, int, byte) {
		if elIdx >= len(grid) {
			return false, elIdx, 0
		}
		el := grid[elIdx][colIdx]
		idx := elIdx
		elIdx++
		return true, idx, el
	}
}

func IterColReverse(grid [][]byte, colIdx int) func() (bool, int, byte) {
	elIdx := len(grid) - 1
	return func() (bool, int, byte) {
		if elIdx < 0 {
			return false, elIdx, 0
		}
		el := grid[elIdx][colIdx]
		idx := elIdx
		elIdx--
		return true, idx, el
	}
}
