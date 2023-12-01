package day1

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func calculatedCalibration(calibrationvalues string, useStringNumerals bool) int {
	scanner := bufio.NewScanner(strings.NewReader(calibrationvalues))

	var sum = 0
	var counter = 0
	for scanner.Scan() {
		line := scanner.Text()
		sum = calculateTheSum(line, useStringNumerals, sum)
		testSum := calculateTheSum(line, true, 0)
		fmt.Printf("#%d: %s : %d\n", counter, line, testSum)
		counter++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error occurred: %v\n", err)
	}
	return sum
}

func calculateTheSum(line string, numerals bool, sum int) int {
	leftIndex, numeralLeft := findLeftNumber(line, numerals)
	rightIndex, numeralRight := findRightNumber(line, numerals)

	var result int
	var leftResult string
	var rightResult string
	if numeralLeft != -1 {
		leftResult = strconv.Itoa(numeralLeft)
	} else {
		leftResult = string(line[leftIndex])
	}
	if numeralRight != -1 {
		rightResult = strconv.Itoa(numeralRight)
	} else {
		rightResult = string(line[rightIndex])
	}

	if leftIndex == rightIndex && leftResult == rightResult {
		result, _ = strconv.Atoi(leftResult + leftResult)
	} else {
		result, _ = strconv.Atoi(leftResult + rightResult)
	}

	sum += result
	return sum
}

func findRightNumber(line string, numerals bool) (int, int) {
	numeral := ""

	var rightIndex int
	var numeralRight = -1
	for b := len(line) - 1; b >= 0; b-- {

		_, isInt := CharToNum(rune(line[b]))
		if isInt {
			rightIndex = b
			break
		} else {
			if numerals {

				numeral = string(rune(line[b])) + numeral
				intval, isPossible := isPossibleNumeral(numeral)
				if isPossible && intval > -1 {
					numeralRight = intval
					rightIndex = b
					break
				} else if !isPossible {
					if len(numeral) > 1 {
						numeral = string(rune(line[b]))
					} else {
						numeral = ""
					}
				}
			}
		}
	}
	return rightIndex, numeralRight
}

func findLeftNumber(line string, numerals bool) (int, int) {
	var leftIndex int
	var numeralLeft = -1
	var numeral = ""
	for a := 0; a < len(line); a++ {
		_, isInt := CharToNum(rune(line[a]))
		if isInt {
			leftIndex = a
			break
		} else {
			if numerals {

				numeral += string(rune(line[a]))
				intval, isPossible := isPossibleNumeral(numeral)
				if isPossible && intval > -1 {
					numeralLeft = intval
					leftIndex = a
					break
				} else if !isPossible {
					if len(numeral) > 1 {
						numeral = string(rune(line[a]))
					} else {
						numeral = ""
					}
				}
			}
		}
	}
	return leftIndex, numeralLeft
}

func CharToNum(r rune) (int, bool) {
	if '0' <= r && r <= '9' {
		return int(r) - '0', true
	}
	return 0, false
}

var numerals = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func isPossibleNumeral(val string) (int, bool) {
	for key, num := range numerals {
		if val == key {
			return num, true
		}

		if strings.HasPrefix(key, val) {
			return -1, true
		}
		if strings.HasSuffix(key, val) {
			return -1, true
		}
	}
	return -1, false
}
