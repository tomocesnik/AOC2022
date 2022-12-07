package main

import (
	"aoc2022/day07/common"
	"aoc2022/util"
	"fmt"
)

const inputFile = "../input.txt"

func sumDirsOfMaxSize(maxSize int, dirsSizes map[*common.FileNode]int) int {
	sum := 0
	for _, val := range dirsSizes {
		if val < maxSize {
			sum += val
		}
	}
	return sum
}

func main() {
	const maxDirSize = 100000
	lines := util.ReadFileLinesAsArray(inputFile)
	rootNode := common.ParseLinesToFiletree(lines)
	var dirsSizes = make(map[*common.FileNode]int)
	common.FindDirsSizes(rootNode, &dirsSizes)
	sumDirsSizes := sumDirsOfMaxSize(maxDirSize, dirsSizes)
	fmt.Println(sumDirsSizes)
}
