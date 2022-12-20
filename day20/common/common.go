package common

import "aoc2022/util"

type DlItem struct {
	Value int
	Prev  *DlItem
	Next  *DlItem
}

func CreateDoubleLinkedList(lines []string) []DlItem {
	var items []DlItem = make([]DlItem, len(lines))
	for i, v := range lines {
		prevItem := &items[len(items)-1]
		if i > 0 {
			prevItem = &items[i-1]
		}
		nextItem := &items[0]
		if i < (len(items) - 1) {
			nextItem = &items[i+1]
		}
		items[i] = DlItem{Value: util.StringToNum(v), Prev: prevItem, Next: nextItem}
	}
	return items
}

func MixValues(dlItems []DlItem) {
	for i := range dlItems {
		dlItem := &dlItems[i]
		// count item as removed from the list!
		moveSteps := dlItem.Value % (len(dlItems) - 1)
		if moveSteps < 0 {
			moveItemBackwards(dlItem, -moveSteps)
		} else if moveSteps > 0 {
			moveItemForward(dlItem, moveSteps)
		}
	}
}

func moveItemForward(item *DlItem, numSteps int) {
	currentItem := item
	for step := 0; step < numSteps; step++ {
		currentItem = currentItem.Next
	}
	newPrev := currentItem
	newNext := currentItem.Next
	moveItem(item, newPrev, newNext)
}

func moveItemBackwards(item *DlItem, numSteps int) {
	currentItem := item
	for step := 0; step < numSteps; step++ {
		currentItem = currentItem.Prev
	}
	newPrev := currentItem.Prev
	newNext := currentItem
	moveItem(item, newPrev, newNext)
}

func moveItem(item *DlItem, newPrev *DlItem, newNext *DlItem) {
	if (item == newPrev) || (item == newNext) {
		return
	}
	newPrev.Next = item
	newNext.Prev = item
	oldPrev := item.Prev
	oldNext := item.Next
	oldPrev.Next = oldNext
	oldNext.Prev = oldPrev
	item.Prev = newPrev
	item.Next = newNext
}

func GetGroveCoordinates(dlItems []DlItem) []int {
	var zerothItem *DlItem = nil
	for i := range dlItems {
		dlItem := &dlItems[i]
		if dlItem.Value == 0 {
			zerothItem = dlItem
			break
		}
	}

	var groveCoordinates []int
	currentItem := zerothItem
	for i := 0; i < 3000; i++ {
		currentItem = currentItem.Next
		if (i == 999) || (i == 1999) || (i == 2999) {
			groveCoordinates = append(groveCoordinates, currentItem.Value)
		}
	}
	return groveCoordinates
}
