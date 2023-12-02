package day2

import (
	"io"
	"log"
	"os"
	"testing"
)

func TestDay2PartTwoInitiaTests(t *testing.T) {
	input := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`
	want := 8

	input2 := `12 red cubes, 13 green cubes, 14 blue cubes`
	got := calculatePossibleStoneSets(input, input2)
	if got != want {
		t.Errorf("findPossibleAmountOfStones () == %d, want %d", got, want)
	}
}

func TestDay2Execute(t *testing.T) {

	file, err := os.Open("puzzleinput")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	input := string(bytes)
	input2 := `12 red cubes, 13 green cubes, 14 blue cubes`
	want := 2541
	got := calculatePossibleStoneSets(input, input2)

	if got != want {
		t.Errorf("findPossibleAmountOfStones() == %d, want %d", got, want)
	}
}

func TestDay2Partt2(t *testing.T) {

	input := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`
	want := 2286

	got := findMinimumAmountOfPossibleStones(input)

	if got != want {
		t.Errorf("findMinimumAmountOfStones() == %d, want %d", got, want)
	}
}

func TestDay2Part2(t *testing.T) {

	file, err := os.Open("puzzleinput")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	input := string(bytes)
	want := 66016
	got := findMinimumAmountOfPossibleStones(input)

	if got != want {
		t.Errorf("findPossibleAmountOfStones() == %d, want %d", got, want)
	}
}
