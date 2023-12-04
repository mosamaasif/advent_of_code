package main

import (
	"Advent_of_Code/utils"
	"fmt"
	"strconv"
)

const FILE_NAME string = "input.txt"
// this is to find and map words to numbers
var NUMBER_MAP 	map[string]int = map[string]int {
	"one"	: 	1,
	"two"	: 	2,
	"three"	: 	3,
	"four"	: 	4,
	"five"	: 	5,
	"six"	: 	6,
	"seven"	: 	7,
	"eight"	: 	8,
	"nine"	: 	9,
}

// finds ascii digits fron byte values
func isDigit(x byte) bool {
	return x >= '0' && x <= '9'
}

func updateDigits(firstDigit *int, lastDigit *int, val int) {
	if *firstDigit == 0 {
		*firstDigit = val
	}
	*lastDigit = val
}

// uses sliding window technique
func CalcSum(lines *[][]byte) int {
	sum := 0
	maxWinLen := 5 // window can't exceed this length if number is in words
	for _, line := range *lines {
		firstDigit, lastDigit := 0, 0
		lineLen := len(line)
		for winSt, winEnd := 0, 0; winSt < lineLen; winSt++ {
			winEnd = winSt
			
			// if starting index contains a number, just use that as is and continue
			if isDigit(line[winSt]) {
				num, _ := strconv.Atoi(string(line[winSt]))
				updateDigits(&firstDigit, &lastDigit, num)
				continue
			}

			// else use sliding window to find the word if possible
			for winEnd < lineLen && winEnd - winSt < maxWinLen {
				if val, ok := NUMBER_MAP[string(line[winSt:winEnd + 1])]; ok {
					updateDigits(&firstDigit, &lastDigit, val)
					break
				}
				winEnd++
			}
		}
		sum += firstDigit * 10 + lastDigit
	}

	return sum
}

func main() {
	lines, err := utils.ReadFileByLineBytes(FILE_NAME)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	sum, err := utils.ExecuteAndLogTime(CalcSum, &lines)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Sum: %d\n", sum)
}

