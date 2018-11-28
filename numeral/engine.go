package numeral

// Engine the contract for numeral convention
type Engine interface {
	ToNumber(numeral string) (int, error)
}
