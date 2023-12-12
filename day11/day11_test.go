package day11

import (
	"testing"
)

/*
..F7.
.FJ|.
SJ.L7
|F--J
LJ...
*/
var testInput = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

var testInput2 = `FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJIF7FJ-
L---JF-JLJIIIIFJLJJ7
|F|F-JF---7IIIL7L|7|
|FFJF7L7F-JF7IIL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L`

func TestCasePart1(t *testing.T) {
	input := testInput
	want := 374

	sum := Execute(input, false)

	got := sum

	if got != want {
		t.Errorf("Part1() == %d, want %d", got, want)
	}
}

func TestCasePart2(t *testing.T) {
	input := testInput
	want := 8410

	sum := Execute(input, true)

	got := sum

	if got != want {
		t.Errorf("Part2() == %d, want %d", got, want)
	}
}
