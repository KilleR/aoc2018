package main

import (
	"fmt"
	"inputReader"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type event struct {
	time      time.Time
	minute    int
	guardId   string
	eventType string
}

func main() {
	input := inputReader.GetInput("src/day4/input")

	guardRex := regexp.MustCompile("#([0-9]+)")
	timeRex := regexp.MustCompile("[0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:([0-9]{2})")

	guards := make(map[string]int)

	var events []event

	for _, line := range input {
		eventTime := timeRex.FindString(line)
		if eventTime != "" {
			var e event
			t, err := time.Parse("2006-01-02 15:04", eventTime)
			if err != nil {
				log.Fatalln(err)
			}

			e.time = t

			e.minute, _ = strconv.Atoi(timeRex.FindStringSubmatch(line)[1])

			guardMatch := guardRex.FindStringSubmatch(line)
			if len(guardMatch) > 0 {
				e.guardId = guardMatch[1]
			}
			switch {
			case strings.Contains(line, "asleep"):
				e.eventType = "sleep"
			case strings.Contains(line, "up"):
				e.eventType = "wake"
			default:
				e.eventType = "new"
			}

			events = append(events, e)
		}
	}

	// sort times
	sort.Slice(events, func(i, j int) bool {
		return events[i].time.Before(events[j].time)
	})

	var guardId string
	var asleep int
	for i, event := range events {
		//fmt.Println("Event line:", event)

		if event.guardId != "" {
			guardId = event.guardId
			if asleep != 0 {
				fmt.Println("Previous guard was not awake!")
			}

			_, ok := guards[guardId]
			if !ok {
				guards[guardId] = 0
				//fmt.Println("New guard:", guardId)
			} else {
				//fmt.Println("Existing guard:", guardId)
			}
		}

		event.guardId = guardId

		if event.eventType == "sleep" {
			asleep = event.minute
			//fmt.Println("sleep start at:", asleep)
		}
		if event.eventType == "wake" {
			wakeTime := event.minute
			//fmt.Println("Sleep end at:", wakeTime, wakeTime-asleep)

			guards[guardId] += wakeTime - asleep
			asleep = 0
		}

		events[i] = event
	}

	longestSleeper := ""
	for id, sleepTime := range guards {
		if guards[longestSleeper] < sleepTime {
			longestSleeper = id
		}
	}

	fmt.Println("Longest sleeper:", longestSleeper)

	minutes := make(map[int]int, 60)
	for _, event := range events {
		if event.guardId != longestSleeper {
			continue
		}
		if event.eventType == "sleep" {
			asleep = event.minute
			//fmt.Println("sleep start at:", asleep)
		}
		if event.eventType == "wake" {
			wakeTime := event.minute

			for i := asleep; i < wakeTime; i++ {
				minutes[i]++
			}

			asleep = 0
		}
	}

	var maxMinute, maxMinuteCount int
	for minute, count := range minutes {
		if count > maxMinuteCount {
			maxMinute = minute
			maxMinuteCount = count
		}
	}

	fmt.Printf("Max minute is %s at %d with %d\n", longestSleeper, maxMinute, maxMinuteCount)

	longestSleeperId, _ := strconv.Atoi(longestSleeper)
	fmt.Println(maxMinute * longestSleeperId)
}
