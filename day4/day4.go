package day4

import (
	"bufio"
	"embed"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

//go:embed puzzleinput
var puzzleinput embed.FS

var (
	numbers = regexp.MustCompile(`([0-9]+)`)
)

func RunDay4_1() {
	data, err := puzzleinput.ReadFile("puzzleinput")
	if err != nil {
		log.Fatal(err)
	}

	input := string(data)

	result := Part1(input)

	fmt.Printf("Day 4-1 answer is: %d\n", result)
	//RunDay4_1
}

func RunDay4_2() {
	data, err := puzzleinput.ReadFile("puzzleinput")
	if err != nil {
		log.Fatal(err)
	}

	input := string(data)

	result := Part2(input)

	fmt.Printf("Day 4-2 answer is: %d\n", result)
	//RunDay4_1
}

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

func Part2(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))

	var sum = 0
	var counter = 0
	for scanner.Scan() {
		line := scanner.Text()
		FindWinningSum2(line)
		counter++
	}

	for _, val := range resultList {
		sum += val
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error occurred: %v\n", err)
	}
	return sum

}

func FindWinningSum(input string) int {
	scratchedNumbers, winningNumbers := ParseFields(input)

	sum := 0
	for _, scratchedNumber := range scratchedNumbers {
	scratch:
		for i := 0; i < len(winningNumbers); i++ {
			winningNumber := winningNumbers[i]
			if scratchedNumber == winningNumber {
				if sum == 0 {
					sum = 1
				} else {
					sum *= 2
				}
				break scratch
			}
		}
	}
	return sum
}

func FindWinningSum2(input string) {
	winningNumbers, scratchedNumbers := ParseFields(input)

	wins := 0
	for _, scratchedNumber := range scratchedNumbers {
	scratch:
		for i := 1; i < len(winningNumbers); i++ {
			winningNumber := winningNumbers[i]
			if scratchedNumber == winningNumber {
				wins++
				break scratch
			}
		}
	}
	calculateNumberOfCopiesForCard(winningNumbers[0], wins)

}

/*
Parses format "Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19"
To a card number (2), a list of scratched numbers (13, 32, 20, 16...) and a list of winning numbers (61, 30, 68...)
*/
func ParseFields(input string) ([]int, []int) {
	cardContent := strings.Split(input, "|")

	winningNumbersChan := make(chan []int, 1)   // Buffer size of 1
	scratchedNumbersChan := make(chan []int, 1) // Buffer size of 1

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		winningNumbersChan <- processNumerals(cardContent[0])
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		scratchedNumbersChan <- processNumerals(cardContent[1])
	}()

	wg.Wait()
	close(winningNumbersChan)
	close(scratchedNumbersChan)

	winningNumbers, scratchedNumbers := <-winningNumbersChan, <-scratchedNumbersChan

	return winningNumbers, scratchedNumbers
}

func processNumerals(numerals string) []int {
	nums := numbers.FindAllString(numerals, -1)
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

var resultList []int

func calculateNumberOfCopiesForCard(cardNumber int, wins int) {
	// Ensure the slice is large enough
	requiredLength := cardNumber + wins
	if len(resultList) < requiredLength {
		// Grow the slice to the required length
		additionalElements := make([]int, requiredLength*2-len(resultList))
		resultList = append(resultList, additionalElements...)
	}

	resultList[cardNumber-1]++
	for y := 0; y < resultList[cardNumber-1]; y++ {
		for i := 1; i <= wins; i++ {
			index := cardNumber - 1 + i
			if index < len(resultList) {
				resultList[index]++
			}
		}
	}
}
