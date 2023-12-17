package main

import (
	"Advent_of_Code/utils"
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
)

func mapRange(oldValue int, oldMin int, oldMax int, newMin int, newMax int) int {
	newValue := 0
	oldRange := oldMax - oldMin
	if oldRange == 0 {
    	newValue = newMin
	} else
	{
    	newRange := newMax - newMin  
    	newValue = (((oldValue - oldMin) * newRange) / oldRange) + newMin
	}

	return newValue
}

func parseData(lines *[]string) ([][]int, error) {
	data := [][]int{}
	seeds, _ := utils.StrArrToIntArr(strings.FieldsFunc((*lines)[0], func(r rune) bool { return r == ':' || r == ' ' })[1:])
	data = append(data, seeds)

	// skipping empty line and heading line hence why i starts at 3
	mapIdx := 0
	for i := 1; i < len(*lines); i++ {
		line := (*lines)[i]
		if line == "" {
			data = append(data, []int{})
			mapIdx++
			continue
		}
		if !regexp.MustCompile(`\d`).MatchString(line) {
			continue
		}
		if intArr, err := utils.StrArrToIntArr(strings.Split(line, " ")); err != nil {
			return nil, err
		} else {
			data[mapIdx] = append(data[mapIdx], intArr...)
		}
	}

	return data, nil
}

func findClosestLocationForSeeds(seeds *[]int, mapsData *[][]int) int {
	closestLoc := math.MaxInt
	for _, seed := range *seeds {
		sourceVal := seed
		for mapIdx := 0; mapIdx < len(*mapsData); mapIdx++ {
			mapData := (*mapsData)[mapIdx]
			for i := 0; i < len(mapData); i += 3 {
				destSt := mapData[i]
				sourceSt := mapData[i + 1]
				mapSize := mapData[i + 2]

				if sourceVal >= sourceSt && sourceVal <= sourceSt + mapSize {
					sourceVal = mapRange(sourceVal, sourceSt, sourceSt + mapSize, destSt, destSt + mapSize)
					break
				}
			}
		}
		closestLoc = min(closestLoc, sourceVal)
	}
	return closestLoc
}

func findClosestLocationPartOne(mapsData *[][]int) int {
	mappings := (*mapsData)[1:]
	return findClosestLocationForSeeds(&(*mapsData)[0], &mappings)
}

func findClosestLocationPartTwo(mapsData *[][]int) int {
	closestLoc := math.MaxInt
	seedData := (*mapsData)[0]
	mappings := (*mapsData)[1:]
	for seedIdx := 0; seedIdx < len(seedData); seedIdx += 2 {
		seedSt := seedData[seedIdx]
		seedsRange := seedData[seedIdx + 1]
		seeds := []int{}
		for seedVal := seedSt; seedVal < seedSt + seedsRange; seedVal++ {
			seeds = append(seeds, seedVal)
		}
		closestLoc = min(closestLoc, findClosestLocationForSeeds(&seeds, &mappings))
	}
	return closestLoc
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
	parsedData, ok := parseResult.([][]int)
	if !ok {
		fmt.Println("Failed to parse data")
		return
	}

	closestLocOne, err := utils.ExecuteAndLogTime(findClosestLocationPartOne, &parsedData)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Location Part 1: %d\n", closestLocOne)

	closestLocTwo, err := utils.ExecuteAndLogTime(findClosestLocationPartTwo, &parsedData)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Location Part 2: %d\n", closestLocTwo)
}