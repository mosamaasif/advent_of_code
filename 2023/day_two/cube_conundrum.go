package main

import (
	"Advent_of_Code/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func strToInt(s *string) int {
	num, err := strconv.Atoi(*s)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	return num
}

func islineSplitRune(r rune) bool {
	return r == ':' || r == ';' || r == ','
}

func CalcSumPartOne(lines *[]string, maxBallsMap *map[string]int) int {
	sum := 0
	for _, line := range *lines {
		lineSlice := strings.FieldsFunc(line, islineSplitRune)
		gameId := strToInt(&strings.Split(lineSlice[0], " ")[1])
		gamePassed := true
		for idx := 1; idx < len(lineSlice); idx++ {
			trimmedLineSlice := strings.Trim(lineSlice[idx], " ")
			ballsSlice := strings.Split(trimmedLineSlice, " ")
			color := ballsSlice[1]
			
			if count := strToInt(&ballsSlice[0]); count >= 0 && count > (*maxBallsMap)[color] {
				gamePassed = false
				break
			}
		}
		if gamePassed && gameId >= 0 {
			sum += gameId
		}
	}

	return sum
}

func main() {
	if len(os.Args) < 5 {
		fmt.Println("Provide a filename and max balls count per color via cmd line args")
		return
	}

	fileName := os.Args[1]
	lines, err := utils.ReadFileByLine(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	maxR, maxG, maxB := strToInt(&os.Args[2]), strToInt(&os.Args[3]), strToInt(&os.Args[4])

	sumOne, err := utils.ExecuteAndLogTime(CalcSumPartOne, &lines, &map[string]int{ "red": maxR, "green": maxG, "blue": maxB })
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Sum Part 1: %d\n", sumOne)
}