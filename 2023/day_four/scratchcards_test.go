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

	parseData := parseData(&lines)
	calculatedSum := calcSumPartOne(&parseData)
	requiredSum := 13

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

	parseData := parseData(&lines)
	calculatedSum := calcSumPartTwo(&parseData)
	requiredSum := 30

	if calculatedSum != requiredSum {
		t.Errorf("Calculated Sum: %d is not equal to required Sum: %d", calculatedSum, requiredSum)
	}
}