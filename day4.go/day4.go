package day4

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

type Card struct {
	WinningNumbers   []int
	ScratchedNumbers []int
	Points           int
}

func RunDay4_1() {
	data, err := puzzleinput.ReadFile("puzzleinput")
	if err != nil {
		log.Fatal(err)
	}

	input := string(data)

	result := Part1(input)

	fmt.Printf("Day 4-1 answer is: %d\n", result)
}

var (
	game    = regexp.MustCompile(`([0-9]+):`)
	numbers = regexp.MustCompile(`([0-9]+)`)
)

func Part1(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))

	var sum = 0
	var counter = 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += FindWinningSum(line)
		counter++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error occurred: %v\n", err)
	}
	return sum

}

func FindWinningSum(input string) int {
	cardRound := game.FindStringSubmatch(input)
	input = strings.Replace(input, "Card "+cardRound[0], "", 1)

	cardContent := strings.Split(input, "|")
	winningNumerals := numbers.FindAllString(cardContent[0], -1)
	scratchedNumerals := numbers.FindAllString(cardContent[1], -1)

	winningNumbers := make([]int, len(winningNumerals))

	for i := range winningNumerals {
		winningNumbers[i], _ = strconv.Atoi(winningNumerals[i])
	}

	sum := 0
	for _, scratchedNumber := range scratchedNumerals {
		scratchedNumber, _ := strconv.Atoi(scratchedNumber)
		for i := 0; i < len(winningNumbers); i++ {
			winningNumber := winningNumbers[i]
			if scratchedNumber == winningNumber {
				if sum == 0 {
					sum = 1
				} else {
					sum *= 2
				}
			}
		}
	}
	println(cardRound[1] + ": " + strconv.Itoa(sum))
	return sum
}
