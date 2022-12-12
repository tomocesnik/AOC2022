package main

import (
	"aoc2022/day12/common"
	"aoc2022/util"
	"fmt"
	"log"
	"math"
)

const inputFile = "../input.txt"

func getLowestNodes(nodes []common.Node) []*common.Node {
	var lowestNodes []*common.Node
	for i := range nodes {
		n := &nodes[i]
		if n.Value() == 'a' {
			lowestNodes = append(lowestNodes, n)
		}
	}
	return lowestNodes
}

func lenShortestPathFromAnyStartingNode(nodes []common.Node, startNodes []*common.Node, endNode *common.Node) int {
	minPathLen := math.MaxInt
	for i, sn := range startNodes {
		log.Printf("%d of %d\n", i, len(startNodes))
		shortestPath := common.FindShortestPath(nodes, sn, endNode)
		if shortestPath != nil {
			minPathLen = int(math.Min(float64(minPathLen), float64(len(shortestPath))))
		}
	}
	return minPathLen - 1
}

func main() {
	lines := util.ReadFileLinesAsArray(inputFile)
	nodes, _, endNode := common.CreateGraph(lines)
	lowestNodes := getLowestNodes(nodes)
	minLen := lenShortestPathFromAnyStartingNode(nodes, lowestNodes, endNode)
	fmt.Println(minLen)
}
