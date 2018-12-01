package converter

import (
	"errors"
	"testing"

	"github.com/ramadani/intergalactic/numeral/roman"
)

func TestAddSomeUnitsAndGetErrorWhenAliasUnitIsExist(t *testing.T) {
	type aliasUnitTest struct {
		alias string
		symb  string
		err   error
	}

	tests := []aliasUnitTest{
		aliasUnitTest{"glob", "I", nil},
		aliasUnitTest{"prok", "V", nil},
		aliasUnitTest{"pish", "X", nil},
		aliasUnitTest{"tegj", "X", nil},
		aliasUnitTest{"glob", "I", errors.New("glob unit is exists")},
		aliasUnitTest{"pish", "X", errors.New("pish unit is exists")},
	}

	converter := NewConverter(roman.NewRoman())

	for _, tt := range tests {
		err := converter.AddUnit(tt.alias, tt.symb)

		if err != nil && err.Error() != tt.err.Error() {
			t.Errorf("TestAddSomeUnitsAndGetErrorWhenAliasUnitIsExist failed, expected: '%s', got: '%s'", tt.err.Error(), err.Error())
		}
	}
}

func TestAddSomeUnitsAndGetNumBasedOnNumeralEngine(t *testing.T) {
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
			t.Errorf("TestAddSomeUnitsAndGetNumBasedOnNumeralEngine failed, expected: '%d', got: '%d'", tt.out, res)
		}
	}
}

func TestAddSomeUnitsAndGetErrorWhenAliasUnitIsNotExists(t *testing.T) {
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
			t.Errorf("TestAddSomeUnitsAndGetErrorWhenAliasUnitIsNotExists failed, expected: '%s', got: '%s'",
				tt.err.Error(), err.Error())
		}
	}
}
