package pkg

import (
	"strconv"
)

func StringToInt32(number string) (int32, error) {
	numberInt, err := strconv.Atoi(number)
	if err != nil {
		return 0, err
	}
	return int32(numberInt), nil
}
