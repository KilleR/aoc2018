package main

import (
	"fmt"
	"inputReader"
	"time"
)

func main() {
	start := time.Now()
	input := inputReader.GetInput("src/day2/input")

	for i, line := range input {
		for j, compare := range input[i:] {
			//fmt.Println("Comparing ", i, j+i+1, line, compare)
			diffIndex := findDiff(line, compare)
			if diffIndex != -1 {
				fmt.Println("Match!", i, j+i, line, compare)
				fmt.Println(diffIndex, line[:diffIndex]+line[diffIndex+1:])
			}
		}
	}

	fmt.Println("Done", "in", time.Since(start))
}

func findDiff(s1 string, s2 string) int {
	b1 := []byte(s1)
	b2 := []byte(s2)
	nDiff := 0
	diffIndex := -1

	for i, v := range b1 {
		if b2[i] != v {
			nDiff++
			diffIndex = i
		}
	}

	if nDiff > 1 {
		diffIndex = -1
	}

	return diffIndex
}
