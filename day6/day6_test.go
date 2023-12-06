package day6

import (
	"testing"
)

var testInput1 = `Time:      7  15   30
Distance:  9  40  200
`

var testInput2 = `Time:      71530
Distance:  940200
`

func TestCaseDay6Part1(t *testing.T) {
	input := testInput1
	want := 288

	sum := Execute(input, part1Solution)

	got := sum

	if got != want {
		t.Errorf("FindMultipliedSolutions() == %d, want %d", got, want)
	}
}

func TestCaseDay6Part2(t *testing.T) {
	input := testInput2
	want := 71503

	sum := Execute(input, part1Solution)

	got := sum

	if got != want {
		t.Errorf("FindSeedLocations() == %d, want %d", got, want)
	}
}

func TestCaseDay6Part2_2(t *testing.T) {
	input := testInput2
	want := 71503

	sum := Execute(input, part2Solution)

	got := sum

	if got != want {
		t.Errorf("FindSeedLocations() == %d, want %d", got, want)
	}
}
