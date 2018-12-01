package querier

import (
	"strconv"
	"strings"
)

// Querier for querying from the input
type Querier struct{}

// IsType1 for check a query to set alias with symbol
// Example: glob is I, prok is V, etc.
func (q *Querier) IsType1(arr []string) bool {
	isIdx := -1

	for i, s := range arr {
		if s == "is" {
			isIdx = i
			break
		}
	}

	return isIdx == (len(arr)-1)-1
}

// GetType1Values to get values from query type 1
func (q *Querier) GetType1Values(arr []string) (string, string) {
	var aliases []string
	num := arr[len(arr)-1]

	for _, s := range arr {
		if s == "is" {
			break
		}
		aliases = append(aliases, s)
	}

	return strings.Join(aliases, " "), num
}

// IsType2 for check a query to set credit value by alias nums
// Example: glob glob Silver is 34 Credits, glob prok Gold is 57800 Credits
func (q *Querier) IsType2(arr []string) bool {
	return arr[len(arr)-1] == "Credits"
}

// GetType2Values to get values from query type 2
func (q *Querier) GetType2Values(arr []string) (string, string, int) {
	var aliases []string

	for _, s := range arr {
		if s == "is" {
			break
		}
		aliases = append(aliases, s)
	}

	alias := aliases[:len(aliases)-1]
	creditType := aliases[len(aliases)-1]
	sTotal := arr[len(arr)-2]
	total, _ := strconv.Atoi(sTotal)

	return strings.Join(alias, " "), creditType, total
}

// IsType3 for check a query to ask how munch is alias nums
// Example: how much is pish tegj glob glob ?
func (q *Querier) IsType3(arr []string) bool {
	howIdx := -1
	muchIdx := -1
	isIdx := -1

	for i, s := range arr {
		if s == "how" {
			howIdx = i
		} else if s == "much" {
			muchIdx = i
		} else if s == "is" {
			isIdx = i
		}
	}

	return howIdx == 0 && muchIdx == 1 && isIdx == 2
}

// GetType3Values to get values from query type 3
func (q *Querier) GetType3Values(arr []string) string {
	var aliases []string

	arr = arr[:len(arr)-1]

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == "is" {
			break
		}
		aliases = append([]string{arr[i]}, aliases...)
	}

	return strings.Join(aliases, " ")
}

// IsType4 for check a query to ask how many credits by alias nums and credit type
// Example: how many Credits is glob prok Silver ?
func (q *Querier) IsType4(arr []string) bool {
	howIdx := -1
	manyIdx := -1
	creditsIdx := -1

	for i, s := range arr {
		if s == "how" {
			howIdx = i
		} else if s == "many" {
			manyIdx = i
		} else if s == "Credits" {
			creditsIdx = i
		}
	}

	return howIdx == 0 && manyIdx == 1 && creditsIdx == 2
}

// GetType4Values to get values from query type 4
func (q *Querier) GetType4Values(arr []string) (string, string) {
	var aliases []string

	arr = arr[:len(arr)-1]

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == "is" {
			break
		}
		aliases = append([]string{arr[i]}, aliases...)
	}

	alias := aliases[:len(aliases)-1]
	creditType := aliases[len(aliases)-1]

	return strings.Join(alias, " "), creditType
}

// NewQuerier to make instance of Querier
func NewQuerier() *Querier {
	return &Querier{}
}
