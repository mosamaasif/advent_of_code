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

	parseData, err := parseDataPartOne(&lines)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	calculatedMoe := findMarginOfErrorPartOne(&parseData)
	requiredMoe := 288

	if calculatedMoe != requiredMoe {
		t.Errorf("Calculated Margin of Error: %d is not equal to required Margin of Error: %d", calculatedMoe, requiredMoe)
	}
}

func TestPartTwo(t *testing.T) {
	lines, err := utils.ReadFileByLine("test_input_part_two.txt")
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	parseData, err := parseDataPartTwo(&lines)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	calculatedMoe := findMarginOfErrorPartTwo(&parseData)
	requiredMoe := 71503

	if calculatedMoe != requiredMoe {
		t.Errorf("Calculated Margin of Error: %d is not equal to required Margin of Error: %d", calculatedMoe, requiredMoe)
	}
}