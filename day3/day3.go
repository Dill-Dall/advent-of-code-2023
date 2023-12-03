package day3

import (
	"bufio"
	"embed"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

//go:embed puzzleinput
var puzzleinput embed.FS

type Point struct {
	X, Y int
}

func RunDay3_1() {
	data, err := puzzleinput.ReadFile("puzzleinput")
	if err != nil {
		log.Fatal(err)
	}

	input := string(data)

	result := FindParts(input)

	fmt.Printf("Day 3-1 answer is: %d\n", result)
}

func RunDay3_2() {
	data, err := puzzleinput.ReadFile("puzzleinput")
	if err != nil {
		log.Fatal(err)
	}

	input := string(data)

	result := FindParts2(input)

	fmt.Printf("Day 3-2 answer is: %d\n", result)
}

func FindParts(input string) int {
	array := make2dArrayFromString(input)
	return sumNumbersWithAdjacentSymbols(array)
}

func FindParts2(input string) int {
	array := make2dArrayFromString(input)

	multiplyPointMap := getPointersWithTheirAdjacentNumbers(array)
	sum := multiplyAndSumPointsWithExcatleyTwoElements(&multiplyPointMap)

	return sum
}

func sumNumbersWithAdjacentSymbols(array [][]rune) int {
	var sum = 0
	for y, row := range array {
		var numeral = ""
		var adjacent = false
		for x, char := range row {

			if unicode.IsDigit(char) {
				numeral += string(char)
				if !adjacent {
					adjacent = isSymbolAdjacent(y, x, &array)
				}
			} else {
				if adjacent && numeral != "" {
					sum += convertAndSum(numeral)
				}
				numeral = ""
				adjacent = false
			}

			if x == len(row)-1 && adjacent && numeral != "" {
				sum += convertAndSum(numeral)
			}
		}
	}
	return sum
}

func getPointersWithTheirAdjacentNumbers(array [][]rune) map[Point][]string {
	pointOfMatches := make(map[Point][]string)
	for y, row := range array {
		var numeral = ""
		var point *Point
		for x, char := range row {
			if unicode.IsDigit(char) {
				numeral += string(char)
				if point == nil {
					point = getPointOfAdjacentMultiplier(y, x, &array)
				}

			} else {
				if numeral != "" && point != nil {
					pointOfMatches[*point] = append(pointOfMatches[*point], numeral)
				}
				numeral = ""
				point = nil
			}

			if x == len(row)-1 && numeral != "" && point != nil {
				pointOfMatches[*point] = append(pointOfMatches[*point], numeral)
			}
		}
	}
	return pointOfMatches
}

func multiplyAndSumPointsWithExcatleyTwoElements(pointOfMatches *map[Point][]string) int {
	sum := 0

	for _, elements := range *pointOfMatches {
		if len(elements) == 2 {
			sum += convertAndSum(elements[0]) * convertAndSum(elements[1])
		}
	}

	return sum
}

func convertAndSum(numeral string) int {
	completeNumber, err := strconv.Atoi(numeral)
	if err != nil {
		panic(err)
	}
	return completeNumber
}

func make2dArrayFromString(input string) [][]rune {

	var array [][]rune

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()

		newRow := make([]rune, len(line))
		for char, runeValue := range line {
			if char < len(line) {
				newRow[char] = runeValue
			}
		}
		array = append(array, newRow)
	}
	return array
}

func isSymbolAdjacent(y int, x int, array *[][]rune) bool {

	// Check left
	if x > 0 {
		if isNotDigitNorDot((*array)[y][x-1]) {
			return true
		}
		if y > 0 {
			if isNotDigitNorDot((*array)[y-1][x-1]) {
				return true
			}
		}
		if y < len(*array)-1 {
			if isNotDigitNorDot((*array)[y+1][x-1]) {
				return true
			}
		}
	}

	// Check Right
	if x < len((*array)[y])-1 {
		if isNotDigitNorDot((*array)[y][x+1]) {
			return true
		}
		if y > 0 {
			if isNotDigitNorDot((*array)[y-1][x+1]) {
				return true
			}
		}
		if y < len(*array)-1 {
			if isNotDigitNorDot((*array)[y+1][x+1]) {
				return true
			}
		}
	}

	// Check above and below
	if y > 0 {
		if isNotDigitNorDot((*array)[y-1][x]) {
			return true
		}
	}

	if y < len(*array)-1 {
		if isNotDigitNorDot((*array)[y+1][x]) {
			return true
		}
	}

	return false
}

func getPointOfAdjacentMultiplier(y int, x int, array *[][]rune) *Point {

	// Check left
	if x > 0 {

		if (*array)[y][x-1] == '*' {
			return &Point{
				X: x - 1,
				Y: y,
			}
		}
		if y > 0 {
			if (*array)[y-1][x-1] == '*' {
				return &Point{
					X: x - 1,
					Y: y - 1,
				}
			}
		}
		if y < len(*array)-1 {
			if (*array)[y+1][x-1] == '*' {
				return &Point{
					X: x - 1,
					Y: y + 1,
				}
			}
		}
	}

	// Check Right
	if x < len((*array)[y])-1 {
		if (*array)[y][x+1] == '*' {
			return &Point{
				X: x + 1,
				Y: y,
			}
		}
		if y > 0 {
			if (*array)[y-1][x+1] == '*' {
				return &Point{
					X: x + 1,
					Y: y - 1,
				}
			}
		}
		if y < len(*array)-1 {
			if (*array)[y+1][x+1] == '*' {
				return &Point{
					X: x + 1,
					Y: y + 1,
				}
			}
		}
	}

	// Check above and below
	if y > 0 {
		if (*array)[y-1][x] == '*' {
			return &Point{
				X: x,
				Y: y - 1,
			}
		}
	}

	if y < len(*array)-1 {
		if (*array)[y+1][x] == '*' {
			return &Point{
				X: x,
				Y: y + 1,
			}
		}
	}

	return nil
}

func isNotDigitNorDot(r rune) bool {
	return !unicode.IsDigit(r) && r != '.'
}
