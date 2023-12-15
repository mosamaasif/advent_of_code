package utils

import (
	"strconv"
)

type Point struct {
	X, Y int
}

func StrToInt(s *string) (int, error) {
	num, err := strconv.Atoi(*s)
	if err != nil {
		return -1, err
	}
	return num, nil
}

func StrArrToIntArr(strs []string) ([]int, error) {
	numbers := make([]int, len(strs))
	for idx, numStr := range strs {
		if number, err := StrToInt(&numStr); err != nil {
			return nil, err
		} else {
			numbers[idx] = number
		}
	}
	return numbers, nil
}