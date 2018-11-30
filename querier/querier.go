package querier

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
		}
	}

	return isIdx == (len(arr)-1)-1
}

// Query type 2 for query to set credit value by alias nums
// Example: glob glob Silver is 34 Credits, glob prok Gold is 57800 Credits
func (q *Querier) isQueryType2(arr []string) bool {
	return arr[len(arr)-1] == "Credits"
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

func NewQuerier() *Querier {
	return &Querier{}
}
