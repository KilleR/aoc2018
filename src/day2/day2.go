package main

import (
	"fmt"
	"inputReader"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	var twoCount, threeCount, checksum int

	input := inputReader.GetInput("src/day2/input")

	twoCount = 0
	threeCount = 0

	for _, line := range input {
		if isN(line, 2) {
			twoCount++
		}

		if isN(line, 3) {
			threeCount++
		}
	}

	checksum = twoCount * threeCount

	fmt.Println("Done", twoCount, threeCount, checksum, "in", time.Since(start))
}

func isN(s string, n int) bool {
	for _, v := range []byte(s) {
		if strings.Count(s, string(v)) == n {
			return true
		}
	}

	return false
	//if(strings.Count(s, ""))
}
