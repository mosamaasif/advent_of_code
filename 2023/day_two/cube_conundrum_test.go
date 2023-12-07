package main

import (
	"Advent_of_Code/utils"
	"testing"
)

func TestPartOne(t *testing.T) {
	lines, err := utils.ReadFileByLine("test_part_one.txt")
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	maxBallsMap := map[string]int { "red": 12, "green": 13, "blue": 14 }

	calculatedSum := CalcSumPartOne(&lines, &maxBallsMap)
	requiredSum := 8

	if calculatedSum != requiredSum {
		t.Errorf("Calculated Sum: %d is not equal to required sum: %d", calculatedSum, requiredSum)
	}
}