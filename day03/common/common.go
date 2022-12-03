package common

import "log"

func FindSameChar(string1 string, string2 string) byte {
	for _, c1 := range string1 {
		for _, c2 := range string2 {
			if c1 == c2 {
				return byte(c1)
			}
		}
	}
	log.Fatal("No identical character")
	return string1[0]
}

func CharToPriority(char byte) int {
	if char >= 97 {
		return int(char - 96)
	}
	return int(char - 38)
}
