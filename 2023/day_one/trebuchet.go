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

// uses sliding window technique
func calcSum(lines *[][]byte) int {
	sum := 0
	maxWinLen := 5 // window can't exceed this length if number is in words
	for _, line := range *lines {
		firstDigit, lastDigit := 0, 0
		winSt, winEnd, lenLine := 0, 0, len(line)
		for winSt < lenLine {
			// if starting index contains a number, just use that as is and continue
			if isDigit(line[winSt]) {
				num, _ := strconv.Atoi(string(line[winSt]))
				if firstDigit == 0 {
					firstDigit = num
				}
				lastDigit = num
			// else use sliding window to find the word if possible
			} else {
				for winEnd < lenLine && winEnd - winSt < maxWinLen {
					if val, ok := NUMBER_MAP[string(line[winSt:winEnd + 1])]; ok {
						if firstDigit == 0 {
							firstDigit = val
						}
						lastDigit = val
						break
					}
					winEnd++
				}
			}
			// move window forward and reset size
			winSt++
			winEnd = winSt
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

	sum := calcSum(&lines)
	fmt.Printf("Sum of numbers in the file: %d\n", sum)
}

