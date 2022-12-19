package common

import (
	"aoc2022/util"
	"log"
	"math"
	"strings"
)

const StartValveName = "AA"

type Valve struct {
	// static
	Name      string
	FlowRate  int
	Neighbors []*Valve
	// for shortest path calculation
	marked   bool
	distance int
	prev     *Valve
}

type ImportantValve struct {
	Valve     *Valve
	Distances map[string]int
}

const valveNameSize = 2

func ParseValves(lines []string) []Valve {
	valves := make([]Valve, len(lines))
	var allNeighborNames [][]string
	for i, line := range lines {
		lineParts := strings.Split(line, "; ")
		valves[i] = parseValve(lineParts[0])
		neighborNames := parseNeighborNames(lineParts[1])
		allNeighborNames = append(allNeighborNames, neighborNames)
	}

	linkValves(valves, allNeighborNames)
	return valves
}

func linkValves(valves []Valve, neighborNames [][]string) {
	for i := range valves {
		valve := &valves[i]
		nNames := neighborNames[i]
		for _, nName := range nNames {
			neighbor := findValveWithName(nName, valves)
			if neighbor == nil {
				log.Fatalf("Cannot find neighbor for valve %v with name %s\n", valve, nName)
				continue
			}
			valve.Neighbors = append(valve.Neighbors, neighbor)
		}
	}
}

func findValveWithName(name string, valves []Valve) *Valve {
	for i, v := range valves {
		if v.Name == name {
			return &valves[i]
		}
	}
	return nil
}

func parseValve(valveStr string) Valve {
	restStr := strings.TrimPrefix(valveStr, "Valve ")
	valveName := restStr[:valveNameSize]
	flowRateStr := strings.TrimPrefix(restStr[valveNameSize:], " has flow rate=")
	flowRate := util.StringToNum(flowRateStr)
	return Valve{Name: valveName, FlowRate: flowRate}
}

func parseNeighborNames(neighborsStr string) []string {
	trimmedStr := strings.TrimPrefix(neighborsStr, "tunnel leads to valve ")
	trimmedStr = strings.TrimPrefix(trimmedStr, "tunnels lead to valves ")
	return strings.Split(trimmedStr, ", ")
}

func findShortestPath(nodes []Valve, startNode *Valve, endNode *Valve) []*Valve {
	resetNodes(nodes, startNode)

	minNode := startNode
	found := false
	for minNode != nil {
		for _, neighbor := range minNode.Neighbors {
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

	var shortestPath []*Valve
	n := endNode
	for n != nil {
		shortestPath = append(shortestPath, n)
		n = n.prev
	}
	return shortestPath
}

func resetNodes(allNodes []Valve, startNode *Valve) {
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

func markNodeWithSmallestDist(allNodes []Valve) *Valve {
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

func FindImportantValves(valves []Valve, startValveName string) []ImportantValve {
	var importantValves []ImportantValve
	for i := range valves {
		valve := &valves[i]
		if (valve.FlowRate > 0) || (valve.Name == startValveName) {
			importantValves = append(importantValves, ImportantValve{Valve: valve, Distances: map[string]int{}})
		}
	}
	return importantValves
}

func CalcShortestDistances(valves []Valve, importantValves []ImportantValve) {
	for i := 0; i < len(importantValves); i++ {
		for j := i + 1; j < len(importantValves); j++ {
			v1 := &importantValves[i]
			v2 := &importantValves[j]
			shortestPath := findShortestPath(valves, v1.Valve, v2.Valve)
			lenSp := len(shortestPath)
			v1.Distances[v2.Valve.Name] = lenSp
			v2.Distances[v1.Valve.Name] = lenSp
		}
	}
}

func ExtractValveByName(importantValves []ImportantValve, valveName string) (ImportantValve, []ImportantValve) {
	var extractedValve ImportantValve
	var remainingValves []ImportantValve
	for _, iv := range importantValves {
		if iv.Valve.Name != valveName {
			remainingValves = append(remainingValves, iv)
		} else {
			extractedValve = iv
		}
	}
	return extractedValve, remainingValves
}
