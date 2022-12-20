package main

import (
	"aoc2022/day20/common"
	"aoc2022/util"
	"fmt"
)

const inputFile = "../input.txt"

func multiplyByDecryptionKey(dlItems []common.DlItem, decryptionKey int) {
	for i := range dlItems {
		dlItem := &dlItems[i]
		dlItem.Value *= decryptionKey
	}
}

func main() {
	const decryptionKey = 811589153
	const numMixes = 10

	lines := util.ReadFileLinesAsArray(inputFile)
	dlItems := common.CreateDoubleLinkedList(lines)
	multiplyByDecryptionKey(dlItems, decryptionKey)
	for i := 0; i < numMixes; i++ {
		common.MixValues(dlItems)
	}

	groveCoordinates := common.GetGroveCoordinates(dlItems)
	sum := util.SumArray(groveCoordinates)
	fmt.Println(sum)
}
