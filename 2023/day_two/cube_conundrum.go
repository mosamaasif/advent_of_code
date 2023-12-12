package main

import (
	"Advent_of_Code/utils"
	"fmt"
	"os"
	"regexp"
)

// --------------------------- HELPERS ---------------------------//
// parses each line and stores a map of max count for r,g,b balls for each game
func parseData(lines *[]string) []map[string]int {
	data := make([]map[string]int, len(*lines))
	for i, line := range *lines {
		data[i] = map[string]int{"red": 0, "green": 0, "blue": 0}
		parsedLine := regexp.MustCompile("[,;:]\\s").Split(line, -1)
		for j := 1; j < len(parsedLine); j++ {
			ballSlice := regexp.MustCompile("\\s").Split(parsedLine[j], -1)
			color := ballSlice[1]
			if count, err := utils.StrToInt(&ballSlice[0]); err != nil {
				fmt.Println(err.Error())
			} else if count > data[i][color] {
				data[i][color] = count
			}
		}
	}
	return data
}
//---------------------------------------------------------------//


// --------------------------- DRIVER FUNCS ---------------------------//
func calcSumPartOne(data *[]map[string]int, redLimit int, greenLimit int, blueLimit int) int {
	sum := 0
	for idx, colorMap := range *data {
		if colorMap["red"] <= redLimit && colorMap["green"] <= greenLimit && colorMap["blue"] <= blueLimit {
			sum += idx + 1
		}
	}
	return sum
}

func calcSumPartTwo(data *[]map[string]int) int {
	sum := 0
	for _, colorMap := range *data {
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

	// parsing cmd line args
	lines, err := utils.ReadFileByLine(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	maxR, err := utils.StrToInt(&os.Args[2])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	maxG, err := utils.StrToInt(&os.Args[3])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	maxB, err := utils.StrToInt(&os.Args[4])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	
	// performing sum calcs
	parseResult, err := utils.ExecuteAndLogTime(parseData, &lines)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	parsedData, ok := parseResult.([]map[string]int)
	if !ok {
		fmt.Println("Failed to parse data")
		return
	}

	sumOne, err := utils.ExecuteAndLogTime(calcSumPartOne, &parsedData, maxR, maxG, maxB)
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
