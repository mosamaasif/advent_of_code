package main

import (
	"Advent_of_Code/utils"
	"testing"
)

func TestOnePartOne(t *testing.T) {
	lines, err := utils.ReadFileByLine("test_input_one_part_one.txt")
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	calculatedSum := CalcSumPartOne(&lines)
	requiredSum := 4361

	if calculatedSum != requiredSum {
		t.Errorf("Calculated Sum: %d is not equal to required Sum: %d", calculatedSum, requiredSum)
	}
}

func TestTwoPartOne(t *testing.T) {
	lines, err := utils.ReadFileByLine("test_input_two_part_one.txt")
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	calculatedSum := CalcSumPartOne(&lines)
	requiredSum := 925

	if calculatedSum != requiredSum {
		t.Errorf("Calculated Sum: %d is not equal to required Sum: %d", calculatedSum, requiredSum)
	}
}