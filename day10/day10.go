package day10

import (
	"bufio"
	"embed"
	"fmt"
	"log"
	"strings"

	"github.com/fatih/color"
)

var (
	//go:embed puzzleinput
	puzzleinput    embed.FS
	isPartTwo      = false
	labyrinthArray = make([]string, 0)
	startPoint     = Point{}
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

	fmt.Printf("Day 10_1 answer is: %d\n", result)

}

func RunPart_2() {
	data, err := puzzleinput.ReadFile("puzzleinput")
	if err != nil {
		log.Fatal(err)
	}

	input := string(data)

	result := Execute(input, true)

	fmt.Printf("Day 10_2 answer is: %d\n", result)

}

func Execute(input string, thisIsPartTwo bool) int {
	isPartTwo = thisIsPartTwo
	scanner := bufio.NewScanner(strings.NewReader(input))
	counter := 0
	for scanner.Scan() {
		line := scanner.Text()
		labyrinthArray = append(labyrinthArray, line)
		if startPoint == (Point{}) {
			for i, b := range line {
				if b == 'S' {
					startPoint = Point{x: i, y: counter}
				}
			}
		}
		counter++
	}
	copyLabyrinthArray()

	startDirections := findStartPointDirections(startPoint)
	travLengths := make([]int, 0)
	routes := make([][]Point, 0)
	travelLenght, route, err := traverseLabyrinth(startPoint, startDirections[0])
	if err != nil {
		fmt.Println(err)
	} else {
		travLengths = append(travLengths, travelLenght)
		routes = append(routes, route)
	}

	fmt.Println()

	if isPartTwo {
		return parseWrapped(routes[0])
	}
	return travLengths[0] / 2
}

func traverseLabyrinth(startPoint Point, nextPoint Point) (int, []Point, error) {
	counter := 1
	currentPoint := startPoint
	route := make([]Point, 0)
	for {
		route = append(route, currentPoint)
		counter++
		if !testConstraints(nextPoint) {
			return 0, nil, fmt.Errorf("out of bounds or dot")
		}
		if labyrinthArray[nextPoint.y][nextPoint.x] == 'S' {
			fmt.Printf("Routelength : %d\n", len(route))
			return counter - 1, route, nil
		}

		temPoint := nextPoint
		nextPoint = getNextPoint(currentPoint, nextPoint)
		currentPoint = temPoint

	}
}

func testConstraints(point Point) bool {
	withinWalls := point.x < 0 || point.y < 0 || point.x >= len(labyrinthArray[0]) || point.y >= len(labyrinthArray) || labyrinthArray[point.y][point.x] != '.'
	return withinWalls
}

func getNextPoint(from Point, to Point) Point {
	var point1 Point

	toVal := labyrinthArray[to.y][to.x]
	if toVal == '|' {
		if from.y < to.y {
			point1 = Point{x: to.x, y: to.y + 1}
		} else {
			point1 = Point{x: to.x, y: to.y - 1}
		}
	} else if toVal == '-' {
		if from.x < to.x {
			point1 = Point{x: to.x + 1, y: to.y}
		} else {
			point1 = Point{x: to.x - 1, y: to.y}
		}
		//S|F7||||JFL-7LJF-J7--LF-J7-LF--7|L7|LJ|FJ7-L|7---LFJF7LJFJ7LF77LF77LF77LF77LF77LF77LF77LF77LF7
	} else if toVal == 'L' {
		if from.y == to.y {
			point1 = Point{x: to.x, y: to.y - 1}
		} else {
			point1 = Point{x: to.x + 1, y: to.y}
		}
	} else if toVal == 'J' {
		if from.y == to.y {
			point1 = Point{x: to.x, y: to.y - 1}
		} else {
			point1 = Point{x: to.x - 1, y: to.y}
		}

	} else if toVal == '7' {
		if from.y == to.y {
			point1 = Point{x: to.x, y: to.y + 1}
		} else {
			point1 = Point{x: to.x - 1, y: to.y}
		}
	} else if toVal == 'F' {
		if from.y == to.y {
			point1 = Point{x: to.x, y: to.y + 1}
		} else {
			point1 = Point{x: to.x + 1, y: to.y}
		}

	} else if toVal == 'S' {
		point1 = Point{x: to.x, y: to.y}
	} else {
		return Point{}
	}
	return point1
}

func findStartPointDirections(point Point) []Point {

	toNorth := false
	toSouth := false
	toWest := false
	toEast := false

	if point.y > 0 {
		northRune := labyrinthArray[point.y-1][point.x]
		if northRune == '|' || northRune == '7' || northRune == 'F' {
			toNorth = true
		}
	}

	if point.y < len(labyrinthArray)-1 {
		southRune := labyrinthArray[point.y+1][point.x]
		if southRune == '|' || southRune == 'J' || southRune == 'L' {
			toSouth = true
		}
	}

	if point.x > 0 {
		westRune := labyrinthArray[point.y][point.x-1]
		if westRune == '-' || westRune == 'L' || westRune == 'F' {
			toWest = true
		}
	}

	if point.x < len(labyrinthArray[0])-1 {
		eastRune := labyrinthArray[point.y][point.x+1]
		if eastRune == '-' || eastRune == '7' || eastRune == 'J' {
			toEast = true
		}
	}

	paths := make([]Point, 0)

	if toNorth {
		northPath := Point{x: point.x, y: point.y - 1}
		paths = append(paths, northPath)
	}

	if toWest {
		westPath := Point{x: point.x - 1, y: point.y}
		paths = append(paths, westPath)
	}

	if toSouth {
		soutPath := Point{y: point.y + 1, x: point.x}
		paths = append(paths, soutPath)
	}

	if toEast {
		eastPath := Point{x: point.x + 1, y: point.y}
		paths = append(paths, eastPath)
	}
	return paths

}

// copy the labyrinthArray
var copyOfLabyrinthArray [][]rune

// copy the labyrinthArray and convert each string to a slice of runes
func copyLabyrinthArray() {
	copy := make([][]rune, len(labyrinthArray))
	for i, row := range labyrinthArray {

		copy[i] = []rune(row)
	}
	copyOfLabyrinthArray = copy
}

func parseWrapped(route []Point) int {
	for i := 0; i < len(route); i++ {
		copyOfLabyrinthArray[route[i].y][route[i].x] = 'X'
	}

	for y := 0; y < len(copyOfLabyrinthArray); y++ {
		leftMostX := -1
		rightMostX := -1
		for x := 0; x < len(copyOfLabyrinthArray[y]); x++ {
			if copyOfLabyrinthArray[y][x] == 'X' {
				if leftMostX == -1 {
					leftMostX = x
				} else if leftMostX != -1 {
					rightMostX = x
				}
			}
		}
		if rightMostX != -1 && leftMostX != -1 {
			for i := leftMostX + 1; i < rightMostX; i++ {
				if copyOfLabyrinthArray[y][i] != 'X' {
					copyOfLabyrinthArray[y][i] = '0'
				}
			}
		}
	}

	giveVirus(0, 0, 0)

	countZeros := 0
	for y := 0; y < len(copyOfLabyrinthArray); y++ {
		for x := 0; x < len(copyOfLabyrinthArray[0]); x++ {
			if copyOfLabyrinthArray[y][x] == 'X' {
				fmt.Print(color.GreenString("%s", string(labyrinthArray[y][x])))

			} else if copyOfLabyrinthArray[y][x] == '0' {
				fmt.Print(color.RedString("%s", string(copyOfLabyrinthArray[y][x])))
				countZeros++
			} else if copyOfLabyrinthArray[y][x] == '#' {
				fmt.Print(color.BlueString("%s", string(copyOfLabyrinthArray[y][x])))
			} else {
				fmt.Print(string(string(copyOfLabyrinthArray[y][x])))

			}
		}
		println()
	}
	return countZeros

}

func giveVirus(initialY, initialX, counter int) {
	// Increment counter and check for recursion limit
	counter++
	if counter > 10000 {
		return
	}

	// Check if initialY and initialX are within the borders of the array
	if initialY < 0 || initialY >= len(copyOfLabyrinthArray) || initialX < 0 || initialX >= len(copyOfLabyrinthArray[initialY]) {
		return
	}

	// Check if current cell is a wall or already infected
	if copyOfLabyrinthArray[initialY][initialX] == 'X' || copyOfLabyrinthArray[initialY][initialX] == '#' {
		return
	}

	// Infect this cell
	copyOfLabyrinthArray[initialY][initialX] = '#'

	// Recursive calls for adjacent cells
	giveVirus(initialY+1, initialX, counter)
	giveVirus(initialY-1, initialX, counter)
	giveVirus(initialY, initialX+1, counter)
	giveVirus(initialY, initialX-1, counter)

	giveVirus(initialY+1, initialX-1, counter)
	giveVirus(initialY-1, initialX+1, counter)
	giveVirus(initialY-1, initialX-1, counter)
	giveVirus(initialY+1, initialX+1, counter)
}
