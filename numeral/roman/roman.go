package roman

import (
	"fmt"
	"strconv"
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
	repeat := 1
	num := 0

	for i < n {
		numKey := numArr[i]
		val := r.symbols[numKey]

		// Check current symbol is equal to prev symbol and check it can be repeat or not
		if i > 0 && numKey == numArr[i-1] && !r.canBeRepeat(numKey) {
			return 0, fmt.Errorf("Can not be repeat for this numerals: %s", numKey)
		}

		if i+1 < n {
			nextNumKey := numArr[i+1]

			if numKey == nextNumKey {
				// Check the current symbol can be repeated or not
				if !r.canBeRepeat(numKey) {
					return 0, fmt.Errorf("Can not be repeat for this numerals: %s", numKey)
				}

				// Increment repeat
				repeat++

				// Check if repeat is more than 3 times
				if repeat > 3 {
					return 0, fmt.Errorf("Can not be repeat more then 3 times for this numerals: %s", numKey)
				}
			} else {
				// Reset the repeat counter
				repeat = 1
			}

			nextVal := r.symbols[nextNumKey]

			if val < nextVal {
				// Check symbol is can be subtracted with the next symbol or not
				if !r.canBeSubtracted(numKey, nextNumKey) {
					return 0, fmt.Errorf("%s can not be subtracted with %s", numKey, nextNumKey)
				}

				val = (nextVal - val)
				i++
			}
		}

		num += val
		i++
	}

	return num, nil
}

// Check the symbol can be repeated or not
func (r *Roman) canBeRepeat(symb string) bool {
	strVal := strconv.Itoa(r.symbols[symb])
	vals := []rune(strVal)

	// reverse number
	for i := 0; i < len(vals)/2; i++ {
		ii := len(vals) - i - 1
		vals[i], vals[ii] = vals[ii], vals[i]
	}

	num, _ := strconv.Atoi(string(vals))

	return num == 1
}

// Check symbol is can be subtracted with the next symbol or not
func (r *Roman) canBeSubtracted(symb, nextSymb string) bool {
	val := r.symbols[symb]
	nextVal := r.symbols[nextSymb]

	return val*5 == nextVal || val*10 == nextVal
}

// NewRoman to make instance of roman numerals
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
