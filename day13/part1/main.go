package main

import (
	"aoc2022/day13/common"
	"aoc2022/util"
	"fmt"
)

const inputFile = "../input.txt"

func comparePacketPairs(packetPairs []common.PacketPair) []int {
	var correctlyOrderedIndices []int
	for i, pp := range packetPairs {
		if common.ComparePackets(pp.LPacket, pp.RPacket) == common.Less {
			correctlyOrderedIndices = append(correctlyOrderedIndices, i+1)
		}
	}
	return correctlyOrderedIndices
}

func main() {
	lines := util.ReadFileLinesAsArray(inputFile)
	packetPairs := common.ParsePackets(lines)
	correctlyOrderedIndices := comparePacketPairs(packetPairs)
	sumIndices := util.SumArray(correctlyOrderedIndices)
	fmt.Println(sumIndices)
}
