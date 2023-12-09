package main

import (
	"Advent_of_Code/utils"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// --------------------------- HELPERS ---------------------------//
func strToInt(s *string) int {
	num, err := strconv.Atoi(*s)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	return num
}

// parses each line and stores a map of max count for r,g,b balls for each game
func parseData(lines *[]string) []map[string]int {
	data := make([]map[string]int, len(*lines))
	for i, line := range *lines {
		data[i] = map[string]int{"red": 0, "green": 0, "blue": 0}
		parsedLine := regexp.MustCompile("[,;:]\\s").Split(line, -1)
		for j := 1; j < len(parsedLine); j++ {
			ballSlice := regexp.MustCompile("\\s").Split(parsedLine[j], -1)
			color, count := ballSlice[1], strToInt(&ballSlice[0])
			if count > data[i][color] {
				data[i][color] = count
			}
		}
	}
	return data
}
//---------------------------------------------------------------//


// --------------------------- DRIVER FUNCS ---------------------------//
func CalcSumPartOne(lines *[]string, redLimit int, greenLimit int, blueLimit int) int {
	sum := 0
	parsedData := parseData(lines)
	for idx, colorMap := range parsedData {
		if colorMap["red"] <= redLimit && colorMap["green"] <= greenLimit && colorMap["blue"] <= blueLimit {
			sum += idx + 1
		}
	}
	return sum
}

func CalcSumPartTwo(lines *[]string) int {
	sum := 0
	parsedData := parseData(lines)
	for _, colorMap := range parsedData {
		sum += colorMap["red"] * colorMap["green"] * colorMap["blue"]
	}
	return sum
}
//--------------------------------------------------------------------//

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

	sumOne, err := utils.ExecuteAndLogTime(CalcSumPartOne, &lines, maxR, maxG, maxB)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Sum Part 1: %d\n", sumOne)

	sumTwo, err := utils.ExecuteAndLogTime(CalcSumPartTwo, &lines)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Sum Part 2: %d\n", sumTwo)
}
