package main

import (
	"Advent_of_Code/utils"
	"fmt"
	"os"
	"strings"
)

func isLineSplitRune(r rune) bool {
	return r == ':' || r == '|'
}

func findMatchingNumbers(haveNumbers []string, winningNumbers []string) int {
	matchingCount := 0
	for _, haveNumber := range haveNumbers {
		for _, winningNumber := range winningNumbers {
			if haveNumber == winningNumber {
				matchingCount++
			}
		}
	}
	return matchingCount
}

func parseData(lines *[]string) []int {
	matchingCounts := []int{}
	for _, line := range *lines {
		lineSlices := strings.FieldsFunc(line, isLineSplitRune)
		winningNumbers := strings.Fields(strings.Trim(lineSlices[1], " "))
		haveNumbers := strings.Fields(strings.Trim(lineSlices[2], " "))

		matchingCounts = append(matchingCounts, findMatchingNumbers(haveNumbers, winningNumbers))
	}
	return matchingCounts
}

func calcSumPartOne(matchingCounts *[]int) int {
	sum := 0
	for _, matchingCount := range *matchingCounts {
		if matchingCount == 0 {
			continue
		}
		worth := 1
		for i := 1; i < matchingCount; i++ {
			worth *= 2
		}	
		sum += worth
	}
	return sum
}

func calcSumPartTwo(matchingCounts *[]int) int {
	cardsCount := make([]int, len(*matchingCounts))
	for cardIdx, matchingCount := range *matchingCounts {
		cardsCount[cardIdx] += 1
		if(matchingCount == 0) {
			continue
		}
		for i := cardIdx + 1; i < cardIdx + 1 + matchingCount; i++ {
			cardsCount[i] += cardsCount[cardIdx]
		}
	}

	sum := 0
	for _, cardCount := range cardsCount {
		sum += cardCount
	}
	return sum
}

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

	parseResult, err := utils.ExecuteAndLogTime(parseData, &lines)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	parsedData, ok := parseResult.([]int)
	if !ok {
		fmt.Println("Failed to parse data")
		return
	}

	sumOne, err := utils.ExecuteAndLogTime(calcSumPartOne, &parsedData)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Sum Part 1: %d\n", sumOne)

	sumTwo, err := utils.ExecuteAndLogTime(calcSumPartTwo, &parsedData)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Sum Part 2: %d\n", sumTwo)
}