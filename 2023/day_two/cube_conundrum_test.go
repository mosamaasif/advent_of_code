package main

import (
	"Advent_of_Code/utils"
	"testing"
)

func TestSumPartOne(t *testing.T) {
	lines, err := utils.ReadFileByLine("test.txt")
	if err != nil {
		t.Errorf("%s", err.Error())
		return
	}
	maxBallsMap := map[string]int { "red": 12, "green": 13, "blue": 14 }

	get := CalcSumPartOne(&lines, &maxBallsMap)
	want := 8

	if get != want {
		t.Errorf("Calculated Sum: %d does not match required sum: %d", get, want)
	}
}