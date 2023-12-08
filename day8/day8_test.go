package day8

import (
	"testing"
)

var testInput1 = `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`

var testInput2 = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`

func TestCaseDay8Part1(t *testing.T) {
	input := testInput1
	want := 6

	sum := Execute(input, false)

	got := sum

	if got != want {
		t.Errorf("Day8() == %d, want %d", got, want)
	}
}

func TestCaseDay8Part2(t *testing.T) {
	input := testInput2
	want := 6

	sum := Execute(input, true)

	got := sum

	if got != want {
		t.Errorf("Day8() == %d, want %d", got, want)
	}
}
