package main

import (
	"Advent_of_Code/utils"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Race struct {
	TimeLimit, RecordDistance float64
}

// --------------------------- HELPERS ---------------------------//
func parseDataPartOne(lines *[]string) ([]Race, error) {
	if len(*lines) < 2 {
		return nil, fmt.Errorf("incomplete data")
	}

	timeLimits, err := utils.StrArrToIntArr(regexp.MustCompile(`\d+`).FindAllString((*lines)[0], -1))
	if err != nil {
		return nil, fmt.Errorf("failed to parse times")
	}
	recordDistances, err := utils.StrArrToIntArr(regexp.MustCompile(`\d+`).FindAllString((*lines)[1], -1))
	if err != nil {
		return nil, fmt.Errorf("failed to parse distances")
	}
	
	if len(timeLimits) != len(recordDistances) {
		return nil, fmt.Errorf("incomplete Data")
	}

	races := []Race{}
	for i := 0; i < len(timeLimits); i++ {
		races = append(races, Race{ TimeLimit: float64(timeLimits[i]), RecordDistance: float64(recordDistances[i]) })
	}

	return races, nil
}

func parseDataPartTwo(lines *[]string) (Race, error) {
	if len(*lines) < 2 {
		return Race{}, fmt.Errorf("incomplete data")
	}

	numStrs := regexp.MustCompile(`\d+`).FindAllString((*lines)[0], -1)
	timeLimit, err := strconv.Atoi(strings.Join(numStrs, ""))
	if err != nil {
		return Race{}, fmt.Errorf("failed to parse times")
	}

	numStrs = regexp.MustCompile(`\d+`).FindAllString((*lines)[1], -1)
	recordDistance, err := strconv.Atoi(strings.Join(numStrs, ""))
	if err != nil {
		return Race{}, fmt.Errorf("failed to parse distances")
	}

	return Race{ TimeLimit: float64(timeLimit), RecordDistance: float64(recordDistance) }, nil
}

func calculateWinCount(t float64, d float64) int {
	determinant := math.Sqrt((t * t) - (4 * d))
	lowerHoldTimeLimit := int(math.Floor(0.5 * (t - determinant))) // smallest hold time that will break record
	upperHoldTimeLimit := int(math.Ceil(0.5 * (t + determinant))) // largest hold time that will break record
	return upperHoldTimeLimit - lowerHoldTimeLimit - 1
}
//---------------------------------------------------------------//


// --------------------------- DRIVER FUNCS ---------------------------//
func findMarginOfErrorPartOne(races *[]Race) int {
	marginOfError := 1
	for _, race := range *races {
		marginOfError *= calculateWinCount(race.TimeLimit, race.RecordDistance)
	}
	return marginOfError
}

func findMarginOfErrorPartTwo(race *Race) int {
	return calculateWinCount(race.TimeLimit, race.RecordDistance)
}
//--------------------------------------------------------------------//


func main() {
	if len(os.Args) < 2 {
		fmt.Println("Provide a filename cmd line args")
		return
	}

	fileName := os.Args[1]
	lines, err := utils.ReadFileByLine(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// PART ONE
	parseResult, err := utils.ExecuteAndLogTime(parseDataPartOne, &lines)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	parsedDataOne, ok := parseResult.([]Race)
	if !ok {
		fmt.Println("Failed to parse data")
		return
	}

	moeOne, err := utils.ExecuteAndLogTime(findMarginOfErrorPartOne, &parsedDataOne)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Margin of Error Part 1: %d\n", moeOne)
	//----------

	//PART TWO
	parseResult, err = utils.ExecuteAndLogTime(parseDataPartTwo, &lines)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	parsedDataTwo, ok := parseResult.(Race)
	if !ok {
		fmt.Println("Failed to parse data")
		return
	}

	moeTwo, err := utils.ExecuteAndLogTime(findMarginOfErrorPartTwo, &parsedDataTwo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Margin of Error Part 2: %d\n", moeTwo)
}