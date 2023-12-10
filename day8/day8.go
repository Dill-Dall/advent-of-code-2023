package day8

import (
	"bufio"
	"embed"
	"fmt"
	"log"
	"regexp"

	"strings"
)

var (
	//go:embed puzzleinput
	puzzleinput embed.FS

	isPartTwo = false
	iterators = make(map[string]*Node)
	nodes     = make(map[string]*Node)
)

type Node struct {
	left  *Node
	right *Node
	value string
}

func RunPart_1() {
	data, err := puzzleinput.ReadFile("puzzleinput")
	if err != nil {
		log.Fatal(err)
	}

	input := string(data)

	result := Execute(input, false)

	fmt.Printf("Day 8_1 answer is: %d\n", result)

}

func RunPart_2() {
	data, err := puzzleinput.ReadFile("puzzleinput")
	if err != nil {
		log.Fatal(err)
	}

	input := string(data)

	result := Execute(input, true)

	fmt.Printf("Day 8_2 answer is: %d\n", result)

}

func Execute(input string, thisIsPartTwo bool) int {
	isPartTwo = thisIsPartTwo
	scanner := bufio.NewScanner(strings.NewReader(input))
	counter := 0
	var navigater string
	for scanner.Scan() {
		line := scanner.Text()
		if counter == 0 {
			navigater = line
		} else if counter > 1 {
			node := parseNodes(line)
			if node.value == "AAA" && !isPartTwo {
				iterators[node.value] = node
			} else if isPartTwo && strings.HasSuffix(node.value, "A") {
				iterators[node.value[0:2]] = node
			}
		}

		counter++
	}
	returnValue := getNumberOfIterationsForZ(navigater)

	return returnValue

}

func getNumberOfIterationsForZ(navigater string) int {
	navigateCounter := 0

	zFounds := make(map[string]int)

outer:
	for i := 0; i < len(navigater); i++ {
		navigateCounter++
		direction := navigater[i]
		for key := range iterators {
			_, exists := zFounds[key]
			if !exists {
				if direction == 'L' {
					iterators[key] = iterators[key].left
				} else {
					iterators[key] = iterators[key].right
				}

				if isPartTwo && iterators[key].value[2] == 'Z' {
					zFounds[key] = navigateCounter
					if len(zFounds) == len(iterators) {
						break outer
					}
				}
				if !isPartTwo && iterators[key] != nil && iterators[key].value == "ZZZ" {
					return navigateCounter
				}
			}
		}

		if i == len(navigater)-1 {
			i = -1
		}
	}
	return extractLcm(zFounds)
}

func extractLcm(zFounds map[string]int) int {
	values := make([]int, 0, len(zFounds))
	for _, value := range zFounds {
		values = append(values, value)
	}

	result := values[0]
	for i := 1; i < len(values); i++ {
		result = lcm(result, values[i])
	}
	return result
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

var nodesRxp = regexp.MustCompile(`([A-Z0-9]+) = \(([A-Z0-9]+), ([A-Z0-9]+)\)`)

func parseNodes(input string) *Node {

	matches := nodesRxp.FindAllStringSubmatch(input, -1)

	var mainNode Node

	var leftNode = &Node{value: matches[0][2]}
	var rightNode = &Node{value: matches[0][3]}
	if nodes[leftNode.value] == nil {
		nodes[leftNode.value] = leftNode
	} else {
		leftNode = nodes[leftNode.value]
	}
	if nodes[rightNode.value] == nil {
		nodes[rightNode.value] = rightNode
	} else {
		rightNode = nodes[rightNode.value]
	}
	if nodes[matches[0][1]] != nil {
		mainNode := nodes[matches[0][1]]
		mainNode.left = leftNode
		mainNode.right = rightNode
	} else {
		mainNode = Node{
			value: matches[0][1],
			left:  leftNode,
			right: rightNode,
		}
	}

	nodes[mainNode.value] = &mainNode

	return &mainNode
}
