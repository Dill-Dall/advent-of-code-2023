package main

import (
	"aoc-2023/day1"
	"aoc-2023/day10"
	"aoc-2023/day11"
	"aoc-2023/day2"
	"aoc-2023/day3"
	"aoc-2023/day4"
	"aoc-2023/day5"
	"aoc-2023/day6"
	"aoc-2023/day7"
	"aoc-2023/day8"
	"aoc-2023/day9"
	"fmt"
	"os"
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

	dayArg = os.Args[1]
	partArg = os.Args[2]

	functions := map[string]func(){
		"1_1":  day1.RunDay1_1,
		"1_2":  day1.RunDay1_2,
		"2_2":  day2.RunDay2_2,
		"3_1":  day3.RunDay3_1,
		"3_2":  day3.RunDay3_2,
		"4_1":  day4.RunDay4_1,
		"4_2":  day4.RunDay4_2,
		"5_1":  day5.RunDay5_1,
		"5_2":  day5.RunDay5_2,
		"6_1":  day6.RunDay6_2,
		"6_2":  day6.RunDay6_2,
		"7_1":  day7.RunPart_1,
		"7_2":  day7.RunPart_2,
		"8_1":  day8.RunPart_1,
		"8_2":  day8.RunPart_2,
		"9_1":  day9.RunPart_1,
		"9_2":  day9.RunPart_2,
		"10_1": day10.RunPart_1,
		"10_2": day10.RunPart_2,
		"11_1": day11.RunPart_1,
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
