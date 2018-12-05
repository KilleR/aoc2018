package main

import (
	"fmt"
	"inputReader"
	"log"
	"regexp"
	"strconv"
)

func main() {
	input := inputReader.GetInput("src/day3/input")

	var fabric [1000][1000]int

	rex := regexp.MustCompile("#([0-9]+) @ ([0-9]+),([0-9]+): ([0-9]+)x([0-9]+)")
	for _, v := range input {
		matches := rex.FindStringSubmatch(v)
		id, _ := strconv.Atoi(matches[1])
		startX, _ := strconv.Atoi(matches[2])
		startY, _ := strconv.Atoi(matches[3])
		sizeX, _ := strconv.Atoi(matches[4])
		sizeY, _ := strconv.Atoi(matches[5])

		// claim the squares
		for x := startX; x < startX+sizeX; x++ {
			for y := startY; y < startY+sizeY; y++ {
				switch fabric[x][y] {
				case 0: // unclaimed
					fabric[x][y] = id
				case id: // claimed by current ?!?
					log.Println("duplicate claim at", x, y)
				case -1: // already contested
					continue
				default: // claimed by someone else
					fabric[x][y] = -1
				}
			}
		}
	}

	var contestCount int
	for _, row := range fabric {
		for _, col := range row {
			if col == -1 {
				contestCount++
			}
			//fmt.Printf("%4d", col)
		}
	}
	fmt.Println(contestCount, "contested")

inputLoop:
	for _, v := range input {
		matches := rex.FindStringSubmatch(v)
		id, _ := strconv.Atoi(matches[1])
		startX, _ := strconv.Atoi(matches[2])
		startY, _ := strconv.Atoi(matches[3])
		sizeX, _ := strconv.Atoi(matches[4])
		sizeY, _ := strconv.Atoi(matches[5])

		// check claim intact
		for x := startX; x < startX+sizeX; x++ {
			for y := startY; y < startY+sizeY; y++ {
				if fabric[x][y] == -1 {
					continue inputLoop
				}
			}
		}
		fmt.Println("Claim", id, "uncontested")
	}
}
