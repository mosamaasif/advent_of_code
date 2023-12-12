package main

import (
	"Advent_of_Code/utils"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

//--------------------------- CONSTS ---------------------------//
const FILE_NAME string = "input.txt"
// this is to find and map words to numbers
var NUMBER_MAP 	map[string]int = map[string]int {
	"zero"	:	0,
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
//--------------------------------------------------------------//


//--------------------------- HELPERS ---------------------------//
func updateDigits(firstDigit *int, lastDigit *int, val int) {
	if *firstDigit == -1 {
		*firstDigit = val
	}
	*lastDigit = val
}
//---------------------------------------------------------------//


//--------------------------- DRIVER FUNCS ---------------------------//
func calcSumPartOne(lines *[]string) int {
	sum := 0
	for _, line := range *lines {
		firstDigit, lastDigit := -1, -1
		for _, r := range line {
			if unicode.IsDigit(r) {
				num, _ := strconv.Atoi(string(r))
				updateDigits(&firstDigit, &lastDigit, num)
			}
		}
		sum += firstDigit * 10 + lastDigit
	}

	return sum
}

func calcSumPartTwo(lines *[]string) int {
	sum := 0
	maxWinLen := 5 // window can't exceed this length if number is in words
	for _, line := range *lines {
		firstDigit, lastDigit := -1, -1
		lineLen := len(line)
		for winSt, winEnd := 0, 0; winSt < lineLen; winSt++ {
			winEnd = winSt
			
			// if starting index contains a number, just use that as is and continue
			if unicode.IsDigit(rune(line[winSt])) {
				numStr := string(line[winSt])
				if num, err := utils.StrToInt(&numStr); err != nil {
					fmt.Println(err.Error())
				} else {
					updateDigits(&firstDigit, &lastDigit, num)
				}
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

	sumOne, err := utils.ExecuteAndLogTime(calcSumPartOne, &lines)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Sum Part 1: %d\n", sumOne)

	sumTwo, err := utils.ExecuteAndLogTime(calcSumPartTwo, &lines)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Sum Part 2: %d\n", sumTwo)
}
