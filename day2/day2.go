package day2

import (
	"bufio"
	"embed"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Game struct {
	GameNumber int
	Sets       []map[string]int
}

var (
	testRgx  = regexp.MustCompile(`([0-9]+) ([a-z]+) cubes`)
	gameRgx  = regexp.MustCompile(`Game ([0-9]+): (.+)`)
	colorRgx = regexp.MustCompile(`([0-9]+) (.+)`)
)

//go:embed puzzleinput
var puzzleinput embed.FS

func RunDay2_2() {
	data, err := puzzleinput.ReadFile("puzzleinput")
	if err != nil {
		log.Fatal(err)
	}

	input := string(data)

	result := findMinimumAmountOfPossibleStones(input)

	fmt.Printf("Day 2-2 answer is: %d\n", result)
}

func calculatePossibleStoneSets(input string, possibleamountofstones string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	testMap := parseStoneCriteria(possibleamountofstones)

	var sum = 0
	var counter = 0
	for scanner.Scan() {
		line := scanner.Text()
		game := parseGame(line)
		value := isValidGame(game, testMap)
		sum += value
		counter++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error occurred: %v\n", err)
	}
	return sum
}

func findMinimumAmountOfPossibleStones(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))

	var sum = 0
	var counter = 0
	for scanner.Scan() {
		line := scanner.Text()
		game := parseGame(line)
		value := multiplyIfValid(game)
		sum += value
		counter++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error occurred: %v\n", err)
	}
	return sum
}

// isValidGame checks if a game's stone count matches the test set.
func isValidGame(game Game, testMap map[string]int) int {
	for _, set := range game.Sets {

		for k, v := range set {
			if testMap[k] < v {
				return 0
			}
		}
	}
	return game.GameNumber
}

// multiplyIfValid multiples the minimun required values of each color in a game.
func multiplyIfValid(game Game) int {
	maxValues := make(map[string]int)

	for _, set := range game.Sets {
		for k, v := range set {
			if maxValues[k] < v {
				maxValues[k] = v
			}
		}
	}

	var sum = 1
	for _, val := range maxValues {
		sum *= val
	}
	if len(maxValues) == 0 {
		return 0
	}
	return sum
}

func parseGame(input string) Game {
	gamematch := gameRgx.FindStringSubmatch(input)

	rawSets := strings.Split(gamematch[2], ";")
	sets := make([]map[string]int, 0)

	gameNumber, _ := strconv.Atoi(gamematch[1])

	for _, rawset := range rawSets {
		set := make(map[string]int)
		items := strings.Split(rawset, ",")

		for _, item := range items {
			colorMatches := colorRgx.FindStringSubmatch(item)
			set[colorMatches[2]], _ = strconv.Atoi(colorMatches[1])

		}
		sets = append(sets, set)
	}

	game := Game{
		GameNumber: gameNumber,
		Sets:       sets,
	}

	return game

}

func parseStoneCriteria(testSet string) map[string]int {
	var testMap = make(map[string]int)

	testmatch := testRgx.FindAllStringSubmatch(testSet, -1)

	for _, item := range testmatch {
		testMap[item[2]], _ = strconv.Atoi(item[1])
	}

	return testMap
}
