package main

import (
	"Advent_of_Code/utils"
	"testing"
)

func TestPartOne(t *testing.T) {
	lines, err := utils.ReadFileByLine("test_input_part_one.txt")
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	calculatedSum := CalcSumPartOne(&lines, 12, 13, 14)
	requiredSum := 8

	if calculatedSum != requiredSum {
		t.Errorf("Calculated Sum: %d is not equal to required sum: %d", calculatedSum, requiredSum)
	}
}

func TestPartTwo(t *testing.T) {
	lines, err := utils.ReadFileByLine("test_input_part_two.txt")
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	calculatedSum := CalcSumPartTwo(&lines)
	requiredSum := 2286

	if calculatedSum != requiredSum {
		t.Errorf("Calculated Sum: %d is not equal to required sum: %d", calculatedSum, requiredSum)
	}
}
