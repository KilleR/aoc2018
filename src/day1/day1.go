package main

import (
	"fmt"
	"inputReader"
	"strconv"
)

func main() {
	input := inputReader.GetInput("src/day1/input")

	frequencies := make(map[int]bool)

	result := 0
	loops := 0
	frequencies[0] = true
repLoop:
	for {
		loops++
		for _, v := range input {
			inc, _ := strconv.Atoi(v)

			result += inc
			_, ok := frequencies[result]
			if !ok {
				frequencies[result] = true
			} else {
				fmt.Println("breaking at:", result)
				break repLoop
			}
		}
	}
	fmt.Println(result, loops)
}
