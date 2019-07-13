package converter

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestStringToInt(t *testing.T) {
	for input, expected := range string2IntData() {
		actual := StrToInt(input)
		assert.Equal(t, actual, expected)
	}
}

func TestStringToFloat(t *testing.T) {
	for input, expected := range string2FloatData() {
		actual := StrToFloat(input)
		assert.Equal(t, actual, expected)
	}
}

func string2IntData() map[string]int {
	x := make(map[string]int, 0)
	x["0"] = 0
	x["0.01"] = 0
	x["0.99"] = 0
	x["1.00"] = 0
	x["1"] = 1
	x["1.25"] = 0
	x["2"] = 2
	x["10.00"] = 0
	x["10"] = 10
	x["non-integer"] = 0
	return x
}

func string2FloatData() map[string]float64 {
	x := make(map[string]float64, 0)
	x["0"] = 0.00
	x["0.01"] = 0.01
	x["0.99"] = 0.99
	x["1.00"] = 1.00
	x["1"] = 1.00
	x["1.25"] = 1.25
	x["10.00"] = 10.00
	x["non-float"] = 0
	return x
}
