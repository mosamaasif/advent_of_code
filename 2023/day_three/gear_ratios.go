package main

import (
	"Advent_of_Code/utils"
	"fmt"
	"os"
	"unicode"
)

func isAdjacentToSymbol(numStIdx int, numEndIdx int, lineNum int, lines *[]string) bool {
	startLine, endLine := max(0, lineNum - 1), min(len(*lines) - 1, lineNum + 1)
	linesToCheck := (*lines)[startLine:endLine + 1]
	// will loop 2 time min (2 lines), and 3 time max (3 lines)
	for _, line := range linesToCheck {
		startRuneIdx, endRuneIdx := max(0, numStIdx - 1), min(len(line) - 1, numEndIdx + 1)
		lineSlice := line[startRuneIdx:endRuneIdx + 1]
		// will loop for digits in num + 2, basically to check surrounding areas of the number
		for _, r := range lineSlice {
			if r != '.' && !unicode.IsDigit(r) {
				return true
			}
		}
	}
	return false
}

func CalcSumPartOne(lines *[]string) int {
	sum := 0
	for lineNum, line := range *lines {
		// loops until end of line
		for lineIdx := 0; lineIdx < len(line); lineIdx++ {
			numStIdx := lineIdx
			// this loops for the length of a number once a digit is found (basically finds the number)
			for lineIdx < len(line) && unicode.IsDigit(rune(line[lineIdx])) {
				lineIdx++
			}
			numEndIdx := lineIdx - 1
			// if no digit/number found and is not adjacent to a symbol
			if numEndIdx - numStIdx < 0 || !isAdjacentToSymbol(numStIdx, numEndIdx, lineNum, lines) {
				continue
			}
			numStr := line[numStIdx:numEndIdx + 1]
			// else adds number to sum
			if number, err := utils.StrToInt(&numStr); err != nil {
				fmt.Println(err.Error())
			} else if number >= 0 {
				sum += number
			}
		}
	}
	return sum
}

func CalcSumPartTwo() int {
	sum := 0
	//TODO: Add code here
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

	sumOne, err := utils.ExecuteAndLogTime(CalcSumPartOne, &lines)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Sum Part 1: %d\n", sumOne)

	// sumTwo, err := utils.ExecuteAndLogTime(CalcSumPartTwo, &lines)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// fmt.Printf("Sum Part 2: %d\n", sumTwo)
}