package converter

import (
	"errors"
	"testing"

	"github.com/ramadani/intergalactic/numeral/roman"
)

func TestAddSomeUnitsAndGetNum(t *testing.T) {
	converter := NewConverter(roman.NewRoman())
	converter.AddUnit("glob", "I")
	converter.AddUnit("prok", "V")
	converter.AddUnit("pish", "X")
	converter.AddUnit("tegj", "L")

	type testUnitNum struct {
		in  string
		out int
		err error
	}

	tests := []testUnitNum{
		testUnitNum{"pish pish prok", 25, nil},
		testUnitNum{"pish tegj glob glob", 42, nil},
	}

	for _, tt := range tests {
		res, _ := converter.GetNum(tt.in)

		if res != tt.out {
			t.Errorf("TestAddSomeUnitsAndGetNum failed, expected: '%d', got: '%d'", tt.out, res)
		}
	}
}

func TestAddSomeUnitsAndReturnNotFound(t *testing.T) {
	converter := NewConverter(roman.NewRoman())
	converter.AddUnit("glob", "I")
	converter.AddUnit("prok", "V")
	converter.AddUnit("pish", "X")
	converter.AddUnit("tegj", "L")

	type testUnitNum struct {
		in  string
		out int
		err error
	}

	tests := []testUnitNum{
		testUnitNum{"good great prok", 25, errors.New("Unit not found")},
		testUnitNum{"pish great nice glob", 42, errors.New("Unit not found")},
	}

	for _, tt := range tests {
		_, err := converter.GetNum(tt.in)

		if err.Error() != tt.err.Error() {
			t.Errorf("TestAddSomeUnitsAndReturnNotFound failed, expected: '%s', got: '%s'",
				tt.err.Error(), err.Error())
		}
	}
}
