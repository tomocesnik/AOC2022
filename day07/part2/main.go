package main

import (
	"aoc2022/day07/common"
	"aoc2022/util"
	"fmt"
)

const inputFile = "../input.txt"

func findSmallestDirSizeToDelete(maxDiskSpace int, neededFreeSpace int, rootNode *common.FileNode, dirsSizes map[*common.FileNode]int) int {
	usedSpace := dirsSizes[rootNode]
	freeSpace := maxDiskSpace - usedSpace
	minSpaceToDelete := neededFreeSpace - freeSpace

	var goodDirSizes []int
	for _, val := range dirsSizes {
		if val >= minSpaceToDelete {
			goodDirSizes = append(goodDirSizes, val)
		}
	}
	_, bestDirSize := util.MinNumberInArray(goodDirSizes)
	return bestDirSize
}

func main() {
	const maxDiskSpace = 70000000
	const neededFreeSpace = 30000000
	lines := util.ReadFileLinesAsArray(inputFile)
	rootNode := common.ParseLinesToFiletree(lines)
	var dirsSizes = make(map[*common.FileNode]int)
	common.FindDirsSizes(rootNode, &dirsSizes)
	smallestDirSize := findSmallestDirSizeToDelete(maxDiskSpace, neededFreeSpace, rootNode, dirsSizes)
	fmt.Println(smallestDirSize)
}
