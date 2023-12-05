package day5

import (
	"bufio"
	"embed"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var (
	//go:embed puzzleinput
	puzzleinput embed.FS
	seeds       []int
	numbersRgx  = regexp.MustCompile(`([0-9]+)`)
	mapDefRgx   = regexp.MustCompile(`([a-z]+)-to-([a-z]+)`)
	seedMappers = make([]Mapper, 0)
)

type Mapper struct {
	from     string
	to       string
	rangeDef [][3]int
}

func RunDay5_1() {
	data, err := puzzleinput.ReadFile("puzzleinput")
	if err != nil {
		log.Fatal(err)
	}

	input := string(data)

	result := Execute(input, part1Solution)

	fmt.Printf("Day 5-1 answer is: %d\n", result)
}

func RunDay5_2() {
	data, err := puzzleinput.ReadFile("puzzleinput")
	if err != nil {
		log.Fatal(err)
	}

	input := string(data)

	result := Execute(input, part2Solution)

	fmt.Printf("Day 5-1 answer is: %d\n", result)
}

func Execute(input string, partSolution func() int) int {

	scanner := bufio.NewScanner(strings.NewReader(input))

	var sum = 0
	var counter = 0
	activeMapperIndex := -1
	for scanner.Scan() {

		line := scanner.Text()

		if len(line) == 0 {
			continue
		}
		if counter == 0 {
			seeds = processNumerals(line)
		}
		counter++

		if strings.HasSuffix(line, ":") {
			activeMapperIndex++
			matches := mapDefRgx.FindAllStringSubmatch(line, -1)
			seedMappers = append(seedMappers, Mapper{
				from: matches[0][1],
				to:   matches[0][2],
			})
			continue
		}

		if unicode.IsDigit(rune(line[0])) {
			numbers := processNumerals(line)
			seedMappers[activeMapperIndex].rangeDef = append(seedMappers[activeMapperIndex].rangeDef, [3]int(numbers))
		}

		sum += 0
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error occurred: %v\n", err)
	}

	minLocation := partSolution()

	return minLocation

}

func part1Solution() int {
	var minLocation = -1

	for i := 0; i < len(seeds); i++ {

		locationOfSeed := parseSeedThroughMap(seeds[i])
		if minLocation == -1 {
			minLocation = locationOfSeed
			continue
		} else if locationOfSeed < minLocation {
			minLocation = locationOfSeed
		}
	}
	return minLocation
}

func part2Solution() int {
	var minLocation = -1

	for i := 0; i < len(seeds)-1; i += 2 {
		for j := 0; j < seeds[i+1]; j++ {
			locationOfSeed := parseSeedThroughMap(seeds[i] + j)
			if minLocation == -1 {
				minLocation = locationOfSeed
				continue
			} else if locationOfSeed < minLocation {
				minLocation = locationOfSeed
			}
		}
	}
	return minLocation
}

type Result struct {
	Location int
}

func parseSeedThroughMap(seed int) int {
	seedState := seed
	for _, seedMapper := range seedMappers {
		for _, rangeDef := range seedMapper.rangeDef {
			if seedState >= rangeDef[1] && seedState <= rangeDef[1]+rangeDef[2] {
				seedState = seedState - rangeDef[1] + rangeDef[0]
				break
			}
		}
	}
	return seedState
}

func processNumerals(numerals string) []int {
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
