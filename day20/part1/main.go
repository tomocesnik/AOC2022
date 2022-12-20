package main

import (
	"aoc2022/util"
	"fmt"
)

const inputFile = "../input.txt"

type dlItem struct {
	value int
	prev  *dlItem
	next  *dlItem
}

func createDoubleLinkedList(lines []string) []dlItem {
	var items []dlItem = make([]dlItem, len(lines))
	for i, v := range lines {
		// count item as removed from the list!
		prevItem := &items[len(items)-1]
		if i > 0 {
			prevItem = &items[i-1]
		}
		nextItem := &items[0]
		if i < (len(items) - 1) {
			nextItem = &items[i+1]
		}
		items[i] = dlItem{value: util.StringToNum(v), prev: prevItem, next: nextItem}
	}
	return items
}

func mixValues(dlItems []dlItem) {
	for i := range dlItems {
		dlItem := &dlItems[i]
		moveSteps := dlItem.value % (len(dlItems) - 1)
		if moveSteps < 0 {
			moveItemBackwards(dlItem, -moveSteps)
		} else if moveSteps > 0 {
			moveItemForward(dlItem, moveSteps)
		}
	}
}

func moveItemForward(item *dlItem, numSteps int) {
	currentItem := item
	for step := 0; step < numSteps; step++ {
		currentItem = currentItem.next
	}
	newPrev := currentItem
	newNext := currentItem.next
	moveItem(item, newPrev, newNext)
}

func moveItemBackwards(item *dlItem, numSteps int) {
	currentItem := item
	for step := 0; step < numSteps; step++ {
		currentItem = currentItem.prev
	}
	newPrev := currentItem.prev
	newNext := currentItem
	moveItem(item, newPrev, newNext)
}

func moveItem(item *dlItem, newPrev *dlItem, newNext *dlItem) {
	if (item == newPrev) || (item == newNext) {
		return
	}
	newPrev.next = item
	newNext.prev = item
	oldPrev := item.prev
	oldNext := item.next
	oldPrev.next = oldNext
	oldNext.prev = oldPrev
	item.prev = newPrev
	item.next = newNext
}

func getGroveCoordinates(dlItems []dlItem) []int {
	var zerothItem *dlItem = nil
	for i := range dlItems {
		dlItem := &dlItems[i]
		if dlItem.value == 0 {
			zerothItem = dlItem
			break
		}
	}

	var groveCoordinates []int
	currentItem := zerothItem
	for i := 0; i < 3000; i++ {
		currentItem = currentItem.next
		if (i == 999) || (i == 1999) || (i == 2999) {
			groveCoordinates = append(groveCoordinates, currentItem.value)
		}
	}
	return groveCoordinates
}

func main() {
	lines := util.ReadFileLinesAsArray(inputFile)
	dlItems := createDoubleLinkedList(lines)
	mixValues(dlItems)
	groveCoordinates := getGroveCoordinates(dlItems)
	sum := util.SumArray(groveCoordinates)
	fmt.Println(sum)
}
