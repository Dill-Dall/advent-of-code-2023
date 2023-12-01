package day1

import (
	"io"
	"log"
	"os"
	"testing"
)

func TestDay1PartOneInitial(t *testing.T) {
	input := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`
	want := 142
	got := calculatedCalibration(input, false)
	if got != want {
		t.Errorf("CalibrationValue() == %d, want %d", got, want)
	}
}

func TestDay1PartOneTestPuzzle(t *testing.T) {
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
	want := 55712
	got := calculatedCalibration(input, false)
	if got != want {
		t.Errorf("CalibrationValue() == %d, want %d", got, want)
	}
}

func TestDay1PartTwoInitial1(t *testing.T) {
	input := `two1nine`
	want := 29
	got := calculatedCalibration(input, true)
	if got != want {
		t.Errorf("CalibrationValue() == %d, want %d", got, want)
	}
}

func TestDay1PartTwoInitial2(t *testing.T) {
	input := `eightwothree`
	want := 83
	got := calculatedCalibration(input, true)
	if got != want {
		t.Errorf("CalibrationValue() == %d, want %d", got, want)
	}
}

func TestDay1PartTwoInitial3(t *testing.T) {
	input := `abcone2threexyz`
	want := 13
	got := calculatedCalibration(input, true)
	if got != want {
		t.Errorf("CalibrationValue() == %d, want %d", got, want)
	}
}

func TestDay1PartTwoInitial4(t *testing.T) {
	input := `xtwone3four`
	want := 24
	got := calculatedCalibration(input, true)
	if got != want {
		t.Errorf("CalibrationValue() == %d, want %d", got, want)
	}
}

func TestDay1PartTwoInitial5(t *testing.T) {
	input := `4nineeightseven2`
	want := 42
	got := calculatedCalibration(input, true)
	if got != want {
		t.Errorf("CalibrationValue() == %d, want %d", got, want)
	}
}

func TestDay1PartTwoInitial6(t *testing.T) {
	input := `zoneight234`
	want := 14
	got := calculatedCalibration(input, true)
	if got != want {
		t.Errorf("CalibrationValue() == %d, want %d", got, want)
	}
}

func TestDay1PartTwoInitial7(t *testing.T) {
	input := `7pqrstsixteen`
	want := 76
	got := calculatedCalibration(input, true)
	if got != want {
		t.Errorf("CalibrationValue() == %d, want %d", got, want)
	}
}

func TestDay1PartTwoInitialAll(t *testing.T) {
	input := `54`
	want := 54
	got := calculatedCalibration(input, true)
	if got != want {
		t.Errorf("CalibrationValue() == %d, want %d", got, want)
	}
}
func TestDay1PartTwoInitial8(t *testing.T) {
	input := `374`
	want := 34
	got := calculatedCalibration(input, true)
	if got != want {
		t.Errorf("CalibrationValue() == %d, want %d", got, want)
	}
}

func TestDay1PartTwoTestingWithOffCases(t *testing.T) {
	input := `eighthree
oneight
sevenine
twone`
	want := 201
	got := calculatedCalibration(input, true)

	if got != want {
		t.Errorf("CalibrationValue() == %d, want %d", got, want)
	}
}

// I hate my life. Why is this test failing?
func TestDay1PartTwoTestPuzzle(t *testing.T) {

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
	want := 0
	got := calculatedCalibration(input, true)

	if got != want {
		t.Errorf("CalibrationValue() == %d, want %d", got, want)
	}
}
