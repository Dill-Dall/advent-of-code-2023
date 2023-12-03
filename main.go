package main

import (
	"aoc-2023/day1"
	"aoc-2023/day2"
	"aoc-2023/day3"
	"fmt"
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
	// Time each function by passing them to the timeFunction
	timeFunction(day1.RunDay1_1)
	timeFunction(day1.RunDay1_1_beta)
	timeFunction(day1.RunDay1_2)
	timeFunction(day2.RunDay2_2)
	timeFunction(day3.RunDay3_1)
	timeFunction(day3.RunDay3_2)
}
