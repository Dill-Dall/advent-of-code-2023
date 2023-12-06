package utils

import (
	"log"
	"regexp"
	"strconv"
)

var (
	numbersRgx = regexp.MustCompile(`([0-9]+)`)
)

func ProcessNumerals(numerals string) []int {
	nums := numbersRgx.FindAllString(numerals, -1)
	parsedNums := make([]int, len(nums))
	for i, numeral := range nums {
		num, err := strconv.Atoi(numeral)
		if err != nil {
			log.Fatal(err)
		}
		parsedNums[i] = num
	}
	return parsedNums
}
