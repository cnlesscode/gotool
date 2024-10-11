package maths

import (
	"fmt"
	"math"
	"strconv"
)

func DecimalPlaces32(num float32, places int) (float32, error) {
	numString := fmt.Sprintf("%."+strconv.Itoa(places)+"f", num)
	numNew, err := strconv.ParseFloat(numString, 32)
	if err != nil {
		return num, err
	}
	return float32(numNew), nil
}

func DecimalPlaces64(num float64, places int) (float64, error) {
	numString := fmt.Sprintf("%."+strconv.Itoa(places)+"f", num)
	numNew, err := strconv.ParseFloat(numString, 64)
	if err != nil {
		return num, err
	}
	return numNew, nil
}

func FloatToString(num float64, places int) string {
	return fmt.Sprintf("%."+strconv.Itoa(places)+"f", num)
}

func Round(num float64) int {
	return int(math.Floor(num + 0.5))
}
