package utils

import "strconv"

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