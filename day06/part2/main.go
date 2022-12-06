package main

import (
	"aoc2022/day06/common"
	"aoc2022/util"
	"fmt"
)

const inputFile = "../input.txt"

func main() {
	const bufferLimit = 14
	lines := util.ReadFileLinesAsArray(inputFile)
	markerIdx := common.FindMarker(lines[0], bufferLimit)
	fmt.Println(markerIdx)
}
