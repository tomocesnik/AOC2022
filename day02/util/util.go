package util

import "log"

type RpsMove int

const (
	Rock RpsMove = iota
	Paper
	Scissors
)

func (elem RpsMove) GetScore() int {
	return int(elem) + 1
}

type RpsResult int

const (
	Loss RpsResult = iota
	Draw
	Win
)

func (res RpsResult) GetScore() int {
	return int(res) * 3
}

func ParseMove(identifier string) RpsMove {
	switch identifier {
	case "A":
		return Rock
	case "B":
		return Paper
	case "C":
		return Scissors
	}
	log.Fatal("No mapping for value: " + identifier)
	return Rock
}
