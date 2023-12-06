package day6

import (
	"aoc-2023/utils"
	"bufio"
	"fmt"
	"math"

	"strings"
)

var (
	highScores  = &HighScores{}
	puzzleInput = `Time:        58     99     64     69
Distance:   478   2232   1019   1071`
	puzzleInput2 = `Time:        58996469
Distance:   478223210191071`
)

type HighScores struct {
	times     []int
	distances []int
}

func RunDay6_1() {
	result := Execute(puzzleInput, part1Solution)

	fmt.Printf("Day 6-1 answer is: %d\n", result)
}

func RunDay6_2() {
	result := Execute(puzzleInput2, part2Solution)

	fmt.Printf("Day 6-2 answer is: %d\n", result)
}

func Execute(input string, partSolution func() int) int {
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		line := scanner.Text()
		ParseToHighscores(line, highScores)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error occurred: %v\n", err)
	}

	minLocation := partSolution()

	return minLocation

}

// d = v*t | v = t -> d = t

func part1Solution() int {
	waysToBeatMultiplied := 1
	for key, timeToBeat := range highScores.times {
		fmt.Printf("Time to beat: %d \n", timeToBeat)

		//I is amount of seconds to press the button
		//fmt.Printf("Distance: %d Time: %d \n", distance, speed)
		waysToBeat := findTimesThatCanBeat(timeToBeat, key)

		fmt.Printf("Ways to beat: %d \n", waysToBeat)
		if waysToBeat > 0 {
			waysToBeatMultiplied *= waysToBeat
		}

	}
	return waysToBeatMultiplied
}

func part2Solution() int {

	waysToBeatMultiplied := findAmountOfPossibleWins(float64(highScores.distances[0]), float64(highScores.times[0]))

	return waysToBeatMultiplied
}

func findTimesThatCanBeat(timeToBeat int, key int) int {
	speed := 0
	timeRemaining := timeToBeat - 1
	waysToBeat := 0

	for i := 1; i < timeToBeat; i++ {
		speed = i

		distance := timeRemaining * speed
		if distance > highScores.distances[key] {

			waysToBeat++
		}
		timeRemaining--
		if timeRemaining == 0 {
			break
		}

	}
	return waysToBeat
}

//t = time button is pressed

//v = tx
//T = Tot - tx | Tot - v
//d = v * (Tot - v)   | distance to beat
//d = vTot-v^2
//v^2 + Tot*v + d = 0
//ax^2 + bx + c = 0
//x = -(b +- sqrt(b^2 - 4ac)) / 2a
//x = -(Tot +- sqrt(Tot^2 - 4*1*d)) / 2*1

func findAmountOfPossibleWins(distanceToBeat, timeToBeat float64) int {
	//+1 because I think... that if time = time but distance is higher
	timeToBeat = timeToBeat + 1
	sqrt := math.Sqrt(timeToBeat*timeToBeat - 4*1*distanceToBeat)
	x1 := math.Ceil(-((timeToBeat + sqrt) / 2 * 1))
	x2 := math.Floor(-((timeToBeat - sqrt) / 2 * 1))
	fmt.Printf("X1: %f, X2: %f\n", x1, x2)
	return int(x2 - x1)

}

func ParseToHighscores(line string, highScores *HighScores) {
	if strings.HasPrefix(line, "Time") {
		highScores.times = utils.ProcessNumerals(line)
	} else if strings.HasPrefix(line, "Distance") {
		highScores.distances = utils.ProcessNumerals(line)
	}
}
