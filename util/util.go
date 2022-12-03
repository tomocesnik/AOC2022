package util

import (
	"bufio"
	"log"
	"math"
	"os"
)

func ReadFileLinesAsArray(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	var fileLines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	file.Close()
	return fileLines
}

func MinNumberInArray(arr []int) (int, int) {
	var minNum = math.MaxInt
	var idx = 0
	for i, num := range arr {
		if num < minNum {
			minNum = num
			idx = i
		}
	}
	return idx, minNum
}

func MaxNumberInArray(arr []int) (int, int) {
	var maxNum = math.MinInt
	var idx = 0
	for i, num := range arr {
		if num > maxNum {
			maxNum = num
			idx = i
		}
	}
	return idx, maxNum
}
