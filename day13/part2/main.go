package main

import (
	"aoc2022/day13/common"
	"aoc2022/util"
	"fmt"
	"sort"
)

const inputFile = "../input.txt"

func gatherAllPackets(packetPairs []common.PacketPair) []common.PacketNode {
	var packets []common.PacketNode
	for _, pp := range packetPairs {
		packets = append(packets, pp.LPacket)
		packets = append(packets, pp.RPacket)
	}
	return packets
}

func createDividerPacket(num int) common.PacketNode {
	elementPacket := common.PacketNode{IsLeaf: true, Value: num, Children: nil}
	var inbetweenPacket common.PacketNode
	inbetweenPacket.Children = append(inbetweenPacket.Children, elementPacket)
	var dividerPacket common.PacketNode
	dividerPacket.Children = append(dividerPacket.Children, inbetweenPacket)
	return elementPacket
}

func findPacket(packet common.PacketNode, packets []common.PacketNode) int {
	for i, pkt := range packets {
		if common.ComparePackets(pkt, packet) == common.Same {
			return i
		}
	}
	return -1
}

func main() {
	lines := util.ReadFileLinesAsArray(inputFile)
	packetPairs := common.ParsePackets(lines)
	packets := gatherAllPackets(packetPairs)
	divPkt2 := createDividerPacket(2)
	divPkt6 := createDividerPacket(6)
	packets = append(packets, divPkt2, divPkt6)
	sort.Slice(packets, func(i, j int) bool {
		return common.ComparePackets(packets[i], packets[j]) != common.More
	})
	idx2 := findPacket(divPkt2, packets) + 1
	idx6 := findPacket(divPkt6, packets) + 1
	fmt.Println(idx2 * idx6)
}
