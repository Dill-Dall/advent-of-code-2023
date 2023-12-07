package day7

import (
	"testing"
)

var testInput1 = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

func TestCaseDay7Part1(t *testing.T) {
	input := testInput1
	want := 6440

	sum := Execute(input, false)

	got := sum

	if got != want {
		t.Errorf("Day7() == %d, want %d", got, want)
	}
}

func TestCaseDay7Part2(t *testing.T) {
	input := testInput1
	want := 5905

	sum := Execute(input, true)

	got := sum

	if got != want {
		t.Errorf("Day7() == %d, want %d", got, want)
	}
}
