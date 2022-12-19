package main

import (
	"aoc2022/day16/common"
	"aoc2022/util"
	"fmt"
)

const inputFile = "../input.txt"

func visitIValve(valve common.ImportantValve, unvisitedValves []common.ImportantValve, stepsLimit int) (int, []common.ImportantValve) {
	maxVal := 0
	var bestValves []common.ImportantValve
	for _, uv := range unvisitedValves {
		dist := valve.Distances[uv.Valve.Name]
		remainingSteps := stepsLimit - dist
		if remainingSteps <= 0 {
			continue
		}
		_, newUnvisitedValves := common.ExtractValveByName(unvisitedValves, uv.Valve.Name)
		val, bv := visitIValve(uv, newUnvisitedValves, remainingSteps)
		if val > maxVal {
			maxVal = val
			bestValves = bv
		}
	}
	internalVal := stepsLimit * valve.Valve.FlowRate
	return maxVal + internalVal, append(bestValves, valve)
}

func main() {
	const maxSteps = 30

	lines := util.ReadFileLinesAsArray(inputFile)
	valves := common.ParseValves(lines)
	importantValves := common.FindImportantValves(valves, common.StartValveName)
	common.CalcShortestDistances(valves, importantValves)
	startValve, unvisitedValves := common.ExtractValveByName(importantValves, common.StartValveName)
	maxVal, _ := visitIValve(startValve, unvisitedValves, maxSteps)
	fmt.Println(maxVal)
}
