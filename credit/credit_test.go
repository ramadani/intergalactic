package credit

import (
	"errors"
	"testing"
)

func TestAddSomeTypeAndGetErrorWhenTypeIsExist(t *testing.T) {
	type creditTypeTests struct {
		name   string
		amount float64
		err    error
	}

	tests := []creditTypeTests{
		creditTypeTests{"Silver", 17, nil},
		creditTypeTests{"Gold", 14450, nil},
		creditTypeTests{"Iron", 195.5, nil},
		creditTypeTests{"Silver", 15, errors.New("Silver credit is exists")},
		creditTypeTests{"Iron", 200, errors.New("Iron credit is exists")},
	}

	credit := NewCredit()

	for _, tt := range tests {
		err := credit.AddType(tt.name, tt.amount)

		if err != nil && err.Error() != tt.err.Error() {
			t.Errorf("TestAddSomeTypeAndGetErrorWhenTypeIsExist failed, expected: '%s', got: '%s'", tt.err.Error(), err.Error())
		}
	}
}

func TestAddSomeTypeAndGetErrorWhenTypeIsNotExist(t *testing.T) {
	type creditType struct {
		name   string
		amount float64
	}

	type creditTypeTests struct {
		name string
		err  error
	}

	creditTypes := []creditType{
		creditType{"Silver", 17},
		creditType{"Gold", 14450},
		creditType{"Iron", 195.5},
	}

	tests := []creditTypeTests{
		creditTypeTests{"Silver", nil},
		creditTypeTests{"Gold", nil},
		creditTypeTests{"Iron", nil},
		creditTypeTests{"Great", errors.New("Great credit type is not found")},
		creditTypeTests{"Good", errors.New("Good credit type is not found")},
	}

	credit := NewCredit()

	for _, ct := range creditTypes {
		credit.AddType(ct.name, ct.amount)
	}

	for _, tt := range tests {
		_, err := credit.GetType(tt.name)

		if err != nil && err.Error() != tt.err.Error() {
			t.Errorf("TestAddSomeTypeAndGetErrorWhenTypeIsNotExist failed, expected: '%s', got: '%s'", tt.err.Error(), err.Error())
		}
	}
}

func TestAddSomeTypeAndCalculateItWithNumber(t *testing.T) {
	type creditType struct {
		name   string
		amount float64
	}

	type creditTypeTests struct {
		name string
		num  int
		res  float64
	}

	creditTypes := []creditType{
		creditType{"Silver", 17},
		creditType{"Gold", 14450},
		creditType{"Iron", 195.5},
	}

	tests := []creditTypeTests{
		creditTypeTests{"Silver", 2, 34},
		creditTypeTests{"Gold", 4, 57800},
		creditTypeTests{"Iron", 20, 3910},
	}

	credit := NewCredit()

	for _, ct := range creditTypes {
		credit.AddType(ct.name, ct.amount)
	}

	for _, tt := range tests {
		cTp, _ := credit.GetType(tt.name)
		res := float64(tt.num) * cTp.Amount

		if res != tt.res {
			t.Errorf("TestAddSomeTypeAndCalculateItWithNumber failed, expected: '%f', got: '%f'", tt.res, res)
		}
	}
}
