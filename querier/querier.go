package querier

import (
	"strconv"
	"strings"
)

type Querier struct{}

func (q *Querier) Query(stmt string) {

}

// Query type 2 for query to set alias with symbol
// Example: glob is I, prok is V, etc.
func (q *Querier) isQueryType1(arr []string) bool {
	isIdx := -1

	for i, s := range arr {
		if s == "is" {
			isIdx = i
			break
		}
	}

	return isIdx == (len(arr)-1)-1
}

// Get values from query type 1
func (q *Querier) getQueryType1Values(arr []string) (string, string) {
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

// Query type 2 for query to set credit value by alias nums
// Example: glob glob Silver is 34 Credits, glob prok Gold is 57800 Credits
func (q *Querier) isQueryType2(arr []string) bool {
	return arr[len(arr)-1] == "Credits"
}

// Get values from query type 2
func (q *Querier) getQueryType2Values(arr []string) (string, string, int) {
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

// Query type 3 for query to ask how munch is alias nums
// Example: how much is pish tegj glob glob ?
func (q *Querier) isQueryType3(arr []string) bool {
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

// Get values from query type 3
func (q *Querier) getQueryType3Values(arr []string) string {
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

// Query type 4 for query to ask how many credits by alias nums and credit type
// Example: how many Credits is glob prok Silver ?
func (q *Querier) isQueryType4(arr []string) bool {
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

// Get values from query type 4
func (q *Querier) getQueryType4Values(arr []string) (string, string) {
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

func NewQuerier() *Querier {
	return &Querier{}
}
