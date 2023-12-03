package day1

import (
	"bufio"
	"embed"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

//go:embed puzzleinput
var puzzleinput embed.FS

func RunDay1_1() {
	data, err := puzzleinput.ReadFile("puzzleinput")
	if err != nil {
		log.Fatal(err)
	}

	input := string(data)

	result := calculatedCalibration(input, false)

	fmt.Printf("Day 1-1 answer is: %d\n", result)

}
func RunDay1_2() {
	data, err := puzzleinput.ReadFile("puzzleinput")
	if err != nil {
		log.Fatal(err)
	}

	input := string(data)

	result := SumEdgeNumbersInStrings(input)

	fmt.Printf("Day 1-2 answer is: %d\n", result)

}

func SumEdgeNumbersInStrings(calibrationvalues string) int {
	return SumEdgeNumbersInStringsUsingNumerals(calibrationvalues, true)

}

func SumEdgeNumbersInStringsUsingNumerals(calibrationvalues string, readNumerals bool) int {

	scanner := bufio.NewScanner(strings.NewReader(calibrationvalues))

	var sum = 0
	var counter = 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += regexSum(line, readNumerals)
		/* testSum := regexSum(line)
		fmt.Printf("#%d: %s : %d\n", counter, line, testSum) */
		counter++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error occurred: %v\n", err)
	}
	return sum
}

var regexnumerals = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

var (
	numeralRegex    = regexp.MustCompile(`([1-9]|one|two|three|four|five|six|seven|eight|nine)`)
	nonNumeralRegex = regexp.MustCompile(`([1-9])`)
)

func regexSum(line string, useNumerals bool) int {
	regexp := numeralRegex
	if !useNumerals {
		regexp = nonNumeralRegex
	}
	sum := 0
	var matches []string
	for i := 0; i < len(line); i++ {
		// Finds the leftsmost match from the i: position
		loc := regexp.FindStringIndex(line[i:])
		if loc != nil {
			// overlap matches
			match := line[i:][loc[0]:loc[1]]

			// Convert word numerals to their numeric counterparts
			if useNumerals {
				if num, ok := regexnumerals[match]; ok {
					match = num
				}
			}
			matches = append(matches, match)

			// bump the index
			i += loc[0]
		}
	}

	sum, _ = strconv.Atoi(matches[0] + matches[len(matches)-1])

	return sum
}
