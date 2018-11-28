package roman

import (
	"strings"
)

// Roman numeral convention
type Roman struct {
	symbols map[string]int
}

// ToNumber convert numeral of roman to number
func (r *Roman) ToNumber(numeral string) (int, error) {
	numArr := strings.Split(numeral, "")
	n := len(numArr)
	i := 0
	num := 0

	for i < n {
		val := r.symbols[numArr[i]]

		if i+1 < n {
			nextVal := r.symbols[numArr[i+1]]

			if val < nextVal {
				val = (nextVal - val)
				i++
			}
		}

		num += val
		i++
	}

	return num, nil
}

// NewRoman to make instance of roman
func NewRoman() *Roman {
	symbols := make(map[string]int)
	symbols["I"] = 1
	symbols["V"] = 5
	symbols["X"] = 10
	symbols["L"] = 50
	symbols["C"] = 100
	symbols["D"] = 500
	symbols["M"] = 1000

	return &Roman{symbols}
}
