package main

import (
	"aoc2022/day16/common"
	"aoc2022/util"
	"fmt"
)

const inputFile = "../input.txt"

type visit struct {
	valve            common.ImportantValve
	remainingPathLen int
}

type partialVisitOption struct {
	valve           common.ImportantValve
	unvisitedValves []common.ImportantValve
	dist            int
}

type visitOption struct {
	pVisit          *visit
	eVisit          *visit
	unvisitedValves []common.ImportantValve
	stepsLimit      int
}

func visitIValve(pVisit *visit, eVisit *visit, unvisitedValves []common.ImportantValve, stepsLimit int) int {
	maxVal := 0
	internalVal := 0
	var visitOptions []visitOption
	handlePVisit := handleVisit(pVisit)
	handleEVisit := handleVisit(eVisit)

	var pPartialVisitOptions []partialVisitOption
	if handlePVisit {
		internalVal = internalVal + (stepsLimit * pVisit.valve.Valve.FlowRate)

		pPartialVisitOptions = findPartialVisitOptions(pVisit.valve, unvisitedValves, stepsLimit)
		visitOptions = createSingleVisitOptions(eVisit, pPartialVisitOptions, func(oldVisit, newVisit *visit, unvisitedValves []common.ImportantValve, stepsLimit int) visitOption {
			return visitOption{newVisit, oldVisit, unvisitedValves, stepsLimit}
		}, stepsLimit)
	}

	var ePartialVisitOptions []partialVisitOption
	if handleEVisit {
		internalVal = internalVal + (stepsLimit * eVisit.valve.Valve.FlowRate)

		ePartialVisitOptions = findPartialVisitOptions(eVisit.valve, unvisitedValves, stepsLimit)
		visitOptions = createSingleVisitOptions(pVisit, ePartialVisitOptions, func(oldVisit, newVisit *visit, unvisitedValves []common.ImportantValve, stepsLimit int) visitOption {
			return visitOption{oldVisit, newVisit, unvisitedValves, stepsLimit}
		}, stepsLimit)
	}

	if handlePVisit && handleEVisit {
		visitOptionsP := createDoubleVisitOptions(pPartialVisitOptions, eVisit.valve, func(newVisit, newVisitDerived *visit, unvisitedValves []common.ImportantValve, stepsLimit int) visitOption {
			return visitOption{newVisit, newVisitDerived, unvisitedValves, stepsLimit}
		}, stepsLimit)
		visitOptionsE := createDoubleVisitOptions(ePartialVisitOptions, pVisit.valve, func(newVisit, newVisitDerived *visit, unvisitedValves []common.ImportantValve, stepsLimit int) visitOption {
			return visitOption{newVisitDerived, newVisit, unvisitedValves, stepsLimit}
		}, stepsLimit)
		visitOptions = append(append([]visitOption{}, visitOptionsP...), visitOptionsE...)
	}

	for _, visitOption := range visitOptions {
		val := visitIValve(visitOption.pVisit, visitOption.eVisit, visitOption.unvisitedValves, visitOption.stepsLimit)

		if val > maxVal {
			maxVal = val
		}
	}

	return maxVal + internalVal
}

func handleVisit(visit *visit) bool {
	return (visit != nil) && (visit.remainingPathLen <= 0)
}

func findPartialVisitOptions(valve common.ImportantValve, unvisitedValves []common.ImportantValve, stepsLimit int) []partialVisitOption {
	var partialVisitOptions []partialVisitOption
	for _, uv := range unvisitedValves {
		dist := valve.Distances[uv.Valve.Name]
		remainingSteps := stepsLimit - dist
		if remainingSteps <= 0 {
			continue
		}
		_, newUnvisitedValves := common.ExtractValveByName(unvisitedValves, uv.Valve.Name)
		partialVisitOptions = append(partialVisitOptions, partialVisitOption{uv, newUnvisitedValves, dist})
	}
	return partialVisitOptions
}

func createSingleVisitOptions(existingVisit *visit, partialVisitOptions []partialVisitOption, visitOptionCreator func(*visit, *visit, []common.ImportantValve, int) visitOption, stepsLimit int) []visitOption {
	var visitOptions []visitOption
	for _, pvo := range partialVisitOptions {
		minDist := pvo.dist
		if (existingVisit != nil) && (existingVisit.remainingPathLen < minDist) {
			minDist = existingVisit.remainingPathLen
		}
		var oldVisit *visit = nil
		if existingVisit != nil {
			oldVisit = &visit{existingVisit.valve, existingVisit.remainingPathLen - minDist}
		}
		newVisit := &visit{pvo.valve, pvo.dist - minDist}
		visitOption := visitOptionCreator(oldVisit, newVisit, pvo.unvisitedValves, stepsLimit-minDist)
		visitOptions = append(visitOptions, visitOption)
	}

	if (len(visitOptions) <= 0) && (existingVisit != nil) {
		oldVisit := &visit{existingVisit.valve, 0}
		visitOption := visitOptionCreator(oldVisit, nil, []common.ImportantValve{}, stepsLimit-existingVisit.remainingPathLen)
		visitOptions = append(visitOptions, visitOption)
	}
	return visitOptions
}

func createDoubleVisitOptions(partialVisitOptions []partialVisitOption, valve common.ImportantValve, visitOptionCreator func(*visit, *visit, []common.ImportantValve, int) visitOption, stepsLimit int) []visitOption {
	var visitOptions []visitOption
	for _, pvo := range partialVisitOptions {
		derivedPartialVisitOptions := findPartialVisitOptions(valve, pvo.unvisitedValves, stepsLimit)
		for _, dPvo := range derivedPartialVisitOptions {
			minDist := pvo.dist
			if dPvo.dist < minDist {
				minDist = dPvo.dist
			}
			newVisit := &visit{pvo.valve, pvo.dist - minDist}
			newVisitDerived := &visit{dPvo.valve, dPvo.dist - minDist}
			visitOption := visitOptionCreator(newVisit, newVisitDerived, dPvo.unvisitedValves, stepsLimit-minDist)
			visitOptions = append(visitOptions, visitOption)
		}
	}
	return visitOptions
}

func main() {
	const maxSteps = 26

	lines := util.ReadFileLinesAsArray(inputFile)
	valves := common.ParseValves(lines)
	importantValves := common.FindImportantValves(valves, common.StartValveName)
	common.CalcShortestDistances(valves, importantValves)
	startValve, unvisitedValves := common.ExtractValveByName(importantValves, common.StartValveName)
	pVisit := &visit{startValve, 0}
	eVisit := &visit{startValve, 0}
	maxVal := visitIValve(pVisit, eVisit, unvisitedValves, maxSteps)
	fmt.Println(maxVal)
}
