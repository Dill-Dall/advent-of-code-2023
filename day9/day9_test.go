package day9

import (
	"testing"
)

var testInput = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

func TestCasePart1(t *testing.T) {
	input := testInput
	want := 114

	sum := Execute(input, false)

	got := sum

	if got != want {
		t.Errorf("Part1() == %d, want %d", got, want)
	}
}

func TestCasePart2(t *testing.T) {
	input := testInput
	want := 2

	sum := Execute(input, true)

	got := sum

	if got != want {
		t.Errorf("Part2() == %d, want %d", got, want)
	}
}
