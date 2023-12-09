package day9

import (
	"aoc-2023/utils"
	"bufio"
	"embed"
	"fmt"
	"log"

	"strings"
)

var (
	//go:embed puzzleinput
	puzzleinput embed.FS
	isPartTwo   = false
)

func RunPart_1() {
	data, err := puzzleinput.ReadFile("puzzleinput")
	if err != nil {
		log.Fatal(err)
	}

	input := string(data)

	result := Execute(input, false)

	fmt.Printf("Day 9_1 answer is: %d\n", result)

}

func RunPart_2() {
	data, err := puzzleinput.ReadFile("puzzleinput")
	if err != nil {
		log.Fatal(err)
	}

	input := string(data)

	result := Execute(input, true)

	fmt.Printf("Day 9_2 answer is: %d\n", result)

}

func Execute(input string, thisIsPartTwo bool) int {
	isPartTwo = thisIsPartTwo
	scanner := bufio.NewScanner(strings.NewReader(input))
	counter := 0
	value := 0
	for scanner.Scan() {
		line := scanner.Text()
		numbers := utils.ProcessNumerals(line)
		if isPartTwo {
			value += findValueOfHistory2(numbers)
		} else {
			value += findValueOfHistory(numbers)
		}
		counter++
	}

	return value

}

func findValueOfHistory(numbers []int) int {
	nextRow := make([]int, 0)

	group := make(map[int]struct{})
	for i := 1; i < len(numbers); i++ {
		nextRow = append(nextRow, numbers[i]-numbers[i-1])
		group[numbers[i]-numbers[i-1]] = struct{}{}
	}

	if len(group) == 1 && nextRow[0] == 0 {
		return numbers[len(numbers)-1]
	} else {
		returnValue := findValueOfHistory(nextRow)
		return returnValue + numbers[len(numbers)-1]
	}
}

func findValueOfHistory2(numbers []int) int {
	nextRow := make([]int, 0)

	group := make(map[int]struct{})
	for i := 1; i < len(numbers); i++ {
		nextRow = append(nextRow, numbers[i]-numbers[i-1])
		group[numbers[i]-numbers[i-1]] = struct{}{}
	}

	if len(group) == 1 && nextRow[0] == 0 {
		return numbers[len(numbers)-1]
	} else {
		returnValue := findValueOfHistory2(nextRow)
		return -returnValue + numbers[0]
	}
}
