package common

import (
	"math"
	"strconv"
)

const lenPacketsSegment = 3

type PacketNode struct {
	IsLeaf   bool
	Value    int
	Children []PacketNode
}

type PacketPair struct {
	LPacket PacketNode
	RPacket PacketNode
}

type comparisonResult int

const (
	Less comparisonResult = iota
	More
	Same
)

func ParsePackets(lines []string) []PacketPair {
	var packetPairs []PacketPair
	lineIdx := 0
	for lineIdx < len(lines) {
		lPacket, _ := parsePacket(lines[lineIdx])
		rPacket, _ := parsePacket(lines[lineIdx+1])
		packetPairs = append(packetPairs, PacketPair{*lPacket, *rPacket})
		lineIdx += lenPacketsSegment
	}
	return packetPairs
}

func parsePacket(line string) (*PacketNode, int) {
	node := PacketNode{}
	lenConsumed := 1
	var child *PacketNode = nil
	for i := 0; i < len(line); i += lenConsumed {
		c := line[i]
		lenConsumed = 1

		switch c {
		case '[':
			child, lenConsumed = parsePacket(line[i+1:])
			lenConsumed++
		case ']':
			if child != nil {
				node.Children = append(node.Children, *child)
			}
			return &node, i + 1
		case ',':
			node.Children = append(node.Children, *child)
		default:
			child, lenConsumed = parseNum(line[i:])
		}
	}
	return child, len(line)
}

func parseNum(line string) (*PacketNode, int) {
	value := 0
	for i, c := range line {
		num, err := strconv.Atoi(string(c))
		if err != nil {
			return &PacketNode{true, value, nil}, i
		}
		value = value*10 + num
	}
	return &PacketNode{true, value, nil}, len(line)
}

func ComparePackets(lPacket PacketNode, rPacket PacketNode) comparisonResult {
	if lPacket.IsLeaf && rPacket.IsLeaf {
		if lPacket.Value < rPacket.Value {
			return Less
		} else if lPacket.Value > rPacket.Value {
			return More
		}
		return Same
	} else if !lPacket.IsLeaf && !rPacket.IsLeaf {
		i := 0
		lenLpc := len(lPacket.Children)
		lenRpc := len(rPacket.Children)
		for i < int(math.Max(float64(lenLpc), float64(lenRpc))) {
			lpktDone := i >= lenLpc
			rpktDone := i >= lenRpc
			if lpktDone && !rpktDone {
				return Less
			} else if !lpktDone && rpktDone {
				return More
			}

			cResult := ComparePackets(lPacket.Children[i], rPacket.Children[i])
			if cResult != Same {
				return cResult
			}
			i++
		}
		return Same
	}

	var impostorNode PacketNode
	if lPacket.IsLeaf {
		impostorNode.Children = append(impostorNode.Children, lPacket)
		return ComparePackets(impostorNode, rPacket)
	} else if rPacket.IsLeaf {
		impostorNode.Children = append(impostorNode.Children, rPacket)
		return ComparePackets(lPacket, impostorNode)
	}
	return Same
}
