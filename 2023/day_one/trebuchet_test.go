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

	calculatedSum := calcSumPartOne(&lines)
	requiredSum := 142

	if calculatedSum != requiredSum {
		t.Errorf("Calculated Sum: %d is not equal to required Sum: %d", calculatedSum, requiredSum)
	}
}

func TestPartTwo(t *testing.T) {
	lines, err := utils.ReadFileByLine("test_input_part_two.txt")
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	calculatedSum := calcSumPartTwo(&lines)
	requiredSum := 281

	if calculatedSum != requiredSum {
		t.Errorf("Calculated Sum: %d is not equal to required Sum: %d", calculatedSum, requiredSum)
	}
}