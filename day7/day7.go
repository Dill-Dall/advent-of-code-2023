package day7

import (
	"bufio"
	"embed"
	"fmt"
	"log"
	"os"
	"strconv"

	"strings"
)

var (
	//go:embed puzzleinput
	puzzleinput embed.FS

	game      = newGame()
	labelsMap = map[rune]int{'2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9, 'T': 10, 'J': 1, 'Q': 12, 'K': 13, 'A': 14}
	isPartTwo = false
)

type HandType int

const (
	HighCard HandType = iota + 1
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Game struct {
	Hands     []Hand
	HandTypes map[HandType][]int
}

type Hand struct {
	id      int
	labels  [5]rune
	rank    int
	bet     int
	handSet HandType
}

func newHand(id int, labels [5]rune, bet int) *Hand {

	set := defineHandSet(labels)

	return &Hand{
		id:      id,
		labels:  labels,
		bet:     bet,
		handSet: set,
	}
}

func newGame() *Game {

	return &Game{
		Hands:     make([]Hand, 0),
		HandTypes: make(map[HandType][]int),
	}
}

func RunPart_1() {
	data, err := puzzleinput.ReadFile("puzzleinput")
	if err != nil {
		log.Fatal(err)
	}

	input := string(data)

	result := Execute(input, false)

	fmt.Printf("Day 7_1 answer is: %d\n", result)

}

func RunPart_2() {
	data, err := puzzleinput.ReadFile("puzzleinput")
	if err != nil {
		log.Fatal(err)
	}

	input := string(data)

	result := Execute(input, true)

	fmt.Printf("Day 7_2 answer is: %d\n", result)

}

func Execute(input string, thisIsPartTwo bool) int {
	isPartTwo = thisIsPartTwo
	if !isPartTwo {
		labelsMap['J'] = 11
	}
	scanner := bufio.NewScanner(strings.NewReader(input))
	counter := 0
	for scanner.Scan() {
		line := scanner.Text()
		hand := parseHand(&line, counter)
		game.HandTypes[hand.handSet] = append(game.HandTypes[hand.handSet], counter)
		game.Hands = append(game.Hands, *hand)
		counter++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error occurred: %v\n", err)
	}

	orderedRank := calculateRanks()
	sum := 0
	for i := 0; i < len(orderedRank); i++ {
		rank := len(orderedRank) - i
		bet := game.Hands[orderedRank[i]].bet
		sum += bet * rank
	}

	return sum

}

func calculateRanks() []int {
	orderedRank := make([]int, 0)

	for handType, idsOfHandsWithTheHandType := range game.HandTypes {
		if len(idsOfHandsWithTheHandType) == 0 {
			continue
		}

		for _, handId := range idsOfHandsWithTheHandType {
			insertIdx := len(orderedRank)
			for i, existingHandId := range orderedRank {
				if handType > game.Hands[existingHandId].handSet {
					insertIdx = i
					break
				} else if handType == game.Hands[existingHandId].handSet {
					if isHigher(handId, existingHandId) {
						insertIdx = i
						break
					}
				}
			}
			// Insert handId at the found position
			orderedRank = append(orderedRank[:insertIdx], append([]int{handId}, orderedRank[insertIdx:]...)...)
		}
	}

	return orderedRank
}

func isHigher(handid1, handid2 int) bool {
	hand1 := game.Hands[handid1]
	hand2 := game.Hands[handid2]

	for i := range hand1.labels {
		if labelsMap[hand1.labels[i]] > labelsMap[hand2.labels[i]] {
			return true
		} else if labelsMap[hand1.labels[i]] < labelsMap[hand2.labels[i]] {
			return false
		}
	}
	return false

}

func parseHand(line *string, counter int) *Hand {
	inputSlice := strings.Split(*line, " ")
	if len(inputSlice[0]) != 5 {
		fmt.Println("Invalid input" + inputSlice[0])
		os.Exit(1)
	}

	var labels [5]rune
	for i, r := range inputSlice[0] {
		labels[i] = r
	}

	var bet, err = strconv.Atoi(inputSlice[1])
	if err != nil {
		fmt.Println("Invalid input" + inputSlice[1])
		os.Exit(1)
	}

	return newHand(counter, labels, bet)

}

func defineHandSet(labels [5]rune) HandType {
	groups := make(map[rune]int)
	jsInLabelsSet := 0
	for i := 0; i < 5; i++ {
		if isPartTwo && labels[i] == 'J' {
			jsInLabelsSet++
		} else {
			groups[labels[i]]++
		}
	}

	highest := '0'
	for k, v := range groups {
		if v > groups[highest] {
			highest = k
		}
	}

	groups[highest] += jsInLabelsSet

	hasOnePair := false
	hasThreeOfAKind := false

	for _, v := range groups {
		if v == 5 {
			return FiveOfAKind
		}
		if v == 4 {
			return FourOfAKind
		}

		if v == 3 {
			hasThreeOfAKind = true
			if hasOnePair {
				return FullHouse
			}
		}
		if v == 2 {
			if hasOnePair {
				return TwoPair
			} else if hasThreeOfAKind {
				return FullHouse
			} else {
				hasOnePair = true
			}
		}
	}

	if hasOnePair {
		return OnePair
	}
	if hasThreeOfAKind {
		return ThreeOfAKind
	}
	return HighCard
}
