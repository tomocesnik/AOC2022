package main

import (
	"aoc2022/day12/common"
	"aoc2022/util"
	"fmt"
)

const inputFile = "../input.txt"

func main() {
	lines := util.ReadFileLinesAsArray(inputFile)
	nodes, startNode, endNode := common.CreateGraph(lines)
	shortestPath := common.FindShortestPath(nodes, startNode, endNode)
	fmt.Println(len(shortestPath) - 1)
}
