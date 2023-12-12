package day11

import (
	"bufio"
	"embed"
	"fmt"
	"log"
	"math"
	"strings"
)

var (
	//go:embed puzzleinput
	puzzleinput embed.FS
	isPartTwo   = false
	galaxy      = make([]string, 0)
	points      = make([]Point, 0)
	pointPairs  = make(map[Point]Point)
)

type Point struct {
	x int
	y int
}

func RunPart_1() {
	data, err := puzzleinput.ReadFile("puzzleinput")
	if err != nil {
		log.Fatal(err)
	}

	input := string(data)

	result := Execute(input, false)

	fmt.Printf("Day 11_1 answer is: %d\n", result)

}

func RunPart_2() {
	data, err := puzzleinput.ReadFile("puzzleinput")
	if err != nil {
		log.Fatal(err)
	}

	input := string(data)

	result := Execute(input, true)

	fmt.Printf("Day 11_2 answer is: %d\n", result)

}

var galaxyCounter = 1

func isEmptySpaceX(s string) (string, bool) {
	isEmptySpace := true
	firstChar := s[0]
	lineCopy := ""
	for i := 0; i < len(s); i++ {
		if s[i] != firstChar {
			isEmptySpace = false
		}

		lineCopy += string(s[i])

	}
	return lineCopy, isEmptySpace
}

func ifEmptySpaceAddColumnToGalaxy() {
	for col := 0; col < len(galaxy[0]); col++ {
		isEmptySpace := true

		// Check each row in the current column
		for _, row := range galaxy {
			if col >= len(row) || row[col] != '.' {
				isEmptySpace = false
				break
			}
		}

		// Duplicate the column if it's empty
		if isEmptySpace {
			for i := range galaxy {
				galaxy[i] = galaxy[i][:col+1] + "." + galaxy[i][col+1:]
			}
			col++ // Skip the newly added column
		}
	}
}

func createPairsAndReturnDistanceOfThem() int {
	distanceSum := 0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			point := points[i]
			otherPoint := points[j]

			// Check if the pair already exists (either as (point, otherPoint) or (otherPoint, point))
			if _, exists := pointPairs[point]; exists && pointPairs[point] == otherPoint {
				continue
			}
			if _, exists := pointPairs[otherPoint]; exists && pointPairs[otherPoint] == point {
				continue
			}

			pointPairs[point] = otherPoint
			distance := int(math.Abs(float64(point.x-otherPoint.x)) + math.Abs(float64(point.y-otherPoint.y)))
			distanceSum += distance
		}
	}
	return distanceSum // Return appropriate value or modify the function signature if needed
}

func findPoints() {
	for y, row := range galaxy {
		for x, char := range row {
			if char == '.' {
				continue
			}
			if char == '#' {
				point := Point{x, y}
				points = append(points, point)
			}
		}
	}
}

func Execute(input string, thisIsPartTwo bool) int {
	isPartTwo = thisIsPartTwo
	scanner := bufio.NewScanner(strings.NewReader(input))
	counter := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineCopy, isEmptySpace := isEmptySpaceX(line)
		galaxy = append(galaxy, lineCopy)
		if isEmptySpace {
			galaxy = append(galaxy, lineCopy)
		}
		counter++
	}
	ifEmptySpaceAddColumnToGalaxy()
	findPoints()
	return createPairsAndReturnDistanceOfThem()
}
