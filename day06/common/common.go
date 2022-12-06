package common

import "math"

func HasDuplicateChars(text string) bool {
	for i := range text {
		for j := i + 1; j < len(text); j++ {
			if text[i] == text[j] {
				return true
			}
		}
	}
	return false
}

func FindMarker(line string, bufferLimit int) int {
	for i := range line {
		lowerBound := int(math.Max(0, float64(i-bufferLimit+1)))
		upperBound := i + 1
		if (i >= bufferLimit) && !HasDuplicateChars(line[lowerBound:upperBound]) {
			return i + 1
		}
	}
	return -1
}
