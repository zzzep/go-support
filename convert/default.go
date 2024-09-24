package convert

import (
	"fmt"
	"strconv"
)

func ToInt(a any, defaultValue int) int {
	val := fmt.Sprint(a)
	if val == "" || a == nil {
		return defaultValue
	}
	valInt, err := strconv.Atoi(val)
	if err != nil {
		return defaultValue
	}
	return valInt
}

func ToFloat(a any, defaultValue float64) float64 {
	val := fmt.Sprint(a)
	if val == "" || a == nil {
		return defaultValue
	}
	valInt, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return defaultValue
	}
	return valInt
}

func ToString(a any, defaultValue string) string {
	val := fmt.Sprint(a)
	if val == "" || a == nil {
		return defaultValue
	}
	return val
}
