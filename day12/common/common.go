package common

import (
	"math"
)

type Node struct {
	value     byte
	neighbors []*Node
	distance  int
	prev      *Node
	marked    bool
}

func (n *Node) Value() byte {
	return n.value
}

func CreateGraph(lines []string) ([]Node, *Node, *Node) {
	height := len(lines)
	width := len(lines[0])
	nodes := make([]Node, height*width)

	var startNode *Node
	var endNode *Node

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			value := lines[i][j]
			newNode := Node{value, []*Node{}, math.MaxInt, nil, false}
			idx := i*width + j
			if value == 'S' {
				newNode.value = 'a'
				startNode = &nodes[idx]
			}
			if value == 'E' {
				newNode.value = 'z'
				endNode = &nodes[idx]
			}
			nodes[idx] = newNode
		}
	}
	linkAllNeighbors(height, width, nodes)
	return nodes, startNode, endNode
}

func linkAllNeighbors(height int, width int, nodes []Node) {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			currentNode := &nodes[i*width+j]
			if i > 0 {
				nNode := &nodes[(i-1)*width+j]
				checkAndLinkNeighbors(currentNode, nNode)
			}
			if j > 0 {
				nNode := &nodes[i*width+(j-1)]
				checkAndLinkNeighbors(currentNode, nNode)
			}
		}
	}
}

func checkAndLinkNeighbors(n1 *Node, n2 *Node) {
	if canMove(n1, n2) {
		n1.neighbors = append(n1.neighbors, n2)
	}
	if canMove(n2, n1) {
		n2.neighbors = append(n2.neighbors, n1)
	}
}

func canMove(source *Node, target *Node) bool {
	return (source.value + 1) >= target.value
}

func FindShortestPath(nodes []Node, startNode *Node, endNode *Node) []*Node {
	resetNodes(nodes, startNode)

	minNode := startNode
	found := false
	for minNode != nil {
		for _, neighbor := range minNode.neighbors {
			if neighbor.marked {
				continue
			}

			newDist := minNode.distance + 1
			if newDist < neighbor.distance {
				neighbor.distance = newDist
				neighbor.prev = minNode
			}
		}

		minNode = markNodeWithSmallestDist(nodes)
		if minNode == endNode {
			found = true
			break
		}
	}

	if !found {
		return nil
	}

	var shortestPath []*Node
	n := endNode
	for n != nil {
		shortestPath = append(shortestPath, n)
		n = n.prev
	}
	return shortestPath
}

func resetNodes(allNodes []Node, startNode *Node) {
	for i := range allNodes {
		node := &allNodes[i]
		node.distance = math.MaxInt
		node.prev = nil
		node.marked = false
		if node == startNode {
			node.marked = true
		}
	}
}

func markNodeWithSmallestDist(allNodes []Node) *Node {
	minDist := math.MaxInt
	minIdx := -1
	for i, n := range allNodes {
		if n.marked {
			continue
		}
		if n.distance < minDist {
			minDist = n.distance
			minIdx = i
		}
	}
	if minIdx < 0 {
		return nil
	}
	minNode := &allNodes[minIdx]
	minNode.marked = true
	return minNode
}
