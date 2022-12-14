package main

import (
	"aoc2022/day20/common"
	"aoc2022/util"
	"fmt"
)

const inputFile = "../input.txt"

func main() {
	lines := util.ReadFileLinesAsArray(inputFile)
	dlItems := common.CreateDoubleLinkedList(lines)
	common.MixValues(dlItems)
	groveCoordinates := common.GetGroveCoordinates(dlItems)
	sum := util.SumArray(groveCoordinates)
	fmt.Println(sum)
}
