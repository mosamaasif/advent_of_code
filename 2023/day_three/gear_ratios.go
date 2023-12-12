package main

import (
	"Advent_of_Code/utils"
	"fmt"
	"os"
	"unicode"
)

type part struct {
	x int
	y int
	symbol rune
}

// --------------------------- HELPERS ---------------------------//
func adjacentSymbol(numStIdx int, numEndIdx int, lineNum int, lines *[]string) part {
	startLine, endLine := max(0, lineNum - 1), min(len(*lines) - 1, lineNum + 1)
	linesToCheck := (*lines)[startLine:endLine + 1]
	// will loop 2 time min (2 lines), and 3 time max (3 lines)
	for lineIdx, line := range linesToCheck {
		startRuneIdx, endRuneIdx := max(0, numStIdx - 1), min(len(line) - 1, numEndIdx + 1)
		lineSlice := line[startRuneIdx:endRuneIdx + 1]
		// will loop for digits in num + 2, basically to check surrounding areas of the number
		for runeIdx, r := range lineSlice {
			if r != '.' && !unicode.IsDigit(r) {
				return part{startRuneIdx + runeIdx, startLine + lineIdx, r}
			}
		}
	}
	return part{-1, -1, 0}
}

func parseData(lines *[]string) map[part][]int {
	// each key stores data for adjacent part found with symbol and it's position
	// and each value is a list of numbers adjacent to that part
	data := make(map[part][]int)
	for lineNum, line := range *lines {
		// loops until end of line
		for runeIdx := 0; runeIdx < len(line); runeIdx++ {
			numStIdx := runeIdx
			// this loops for the length of a number once a digit is found (basically finds the number)
			for runeIdx < len(line) && unicode.IsDigit(rune(line[runeIdx])) {
				runeIdx++
			}
			numEndIdx := runeIdx - 1
			// if no digit/number found and is not adjacent to a symbol
			p := adjacentSymbol(numStIdx, numEndIdx, lineNum, lines)
			if numEndIdx - numStIdx < 0 || p.symbol == 0 {
				continue
			}
			numStr := line[numStIdx:numEndIdx + 1]
			// else adds number to sum
			if number, err := utils.StrToInt(&numStr); err != nil {
				fmt.Println(err.Error())
			} else if number >= 0 {
				// add this to the list
				data[p] = append(data[p], number)
			}
		}
	}
	return data
}
//---------------------------------------------------------------//


// --------------------------- DRIVER FUNCS ---------------------------//
func calcSumPartOne(data *map[part][]int) int {
	sum := 0
	for _, numbers := range *data {
		for _, number := range numbers {
			sum += number
		}
	}
	return sum
}

func calcSumPartTwo(data *map[part][]int) int {
	sum := 0
	for part, numbers := range *data {
		if part.symbol != '*' || len(numbers) < 2 {
			continue
		}
		ratio := 1
		for _, number := range numbers {
			ratio *= number
		}
		sum += ratio
	}
	return sum
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

	parseResult, err := utils.ExecuteAndLogTime(parseData, &lines)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	parsedData, ok := parseResult.(map[part][]int)
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