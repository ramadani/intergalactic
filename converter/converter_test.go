package converter

import (
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
	}

	tests := []testUnitNum{
		testUnitNum{"pish pish prok", 25},
		testUnitNum{"pish tegj glob glob", 42},
	}

	for _, tt := range tests {
		res, _ := converter.GetNum(tt.in)

		if res != tt.out {
			t.Errorf("TestAddSomeUnitsAndGetNum failed, expected: '%d', got: '%d'", tt.out, res)
		}
	}
}
