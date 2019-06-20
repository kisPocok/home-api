package converter

import "strconv"

func StrToInt(s string) int {
	if i, err := strconv.Atoi(s); err == nil {
		return i
	}
	return 0
}

func StrToFloat(s string) float64 {
	if f, err := strconv.ParseFloat(s, 64); err == nil {
		return f
	}
	return 0
}
