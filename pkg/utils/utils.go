package utils

import "fmt"

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func CheckWithMessage(err error, errorMessage string) bool {
	if err != nil {
		fmt.Printf(errorMessage)
		return true
	}

	return false
}

func Average(values []uint8) uint8 {
	var sum uint8 = 0
	for i := 0; i < len(values); i++ {
		sum += values[i]
	}
	return sum / uint8(len(values))
}