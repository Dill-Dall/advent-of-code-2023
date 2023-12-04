package main

import (
	"aoc-2023/day1"
	"aoc-2023/day2"
	"aoc-2023/day3"
	"aoc-2023/day4"
	"fmt"
	"strconv"
	"sync"
	"time"
)

// Timer function that takes a function as a parameter and times it
func timeFunction(functionToTime func()) {
	start := time.Now()
	functionToTime()
	duration := time.Since(start)
	fmt.Printf("Function took %s\n\n", duration)
}

func main() {
	dayArg := "all"
	partArg := "all"

	functions := map[string]func(){
		"1_1": day1.RunDay1_1,
		"1_2": day1.RunDay1_2,
		"2_2": day2.RunDay2_2,
		"3_1": day3.RunDay3_1,
		"3_2": day3.RunDay3_2,
		"4_1": day4.RunDay4_1,
		"4_2": day4.RunDay4_2,
	}

	if dayArg == "all" {
		var wg sync.WaitGroup
		for _, function := range functions {
			wg.Add(1)
			go func(f func()) {
				defer wg.Done()
				timeFunction(f)
			}(function)
		}
		wg.Wait()
	} else {
		if partArg == "all" {
			for part := 1; part <= 2; part++ {
				if function, exists := functions[dayArg+"_"+strconv.Itoa(part)]; exists {
					timeFunction(function)
				}
			}
		} else {
			key := dayArg + "_" + partArg
			if function, exists := functions[key]; exists {
				timeFunction(function)

			} else {
				print("Invalid day or part")
			}
		}
	}
}
