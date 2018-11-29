package roman

import (
	"errors"
	"testing"
)

type romanTests struct {
	in  string
	out int
	err error
}

func TestToNumber(t *testing.T) {
	tests := []romanTests{
		romanTests{"XXXIX", 39, nil},
		romanTests{"MVI", 1006, nil},
		romanTests{"MMVI", 2006, nil},
	}

	roman := NewRoman()

	for _, tt := range tests {
		num, err := roman.ToNumber(tt.in)

		if err != tt.err {
			t.Error(err)
		} else {
			expected := tt.out
			actual := num
			if actual != expected {
				t.Errorf("TestToNumber failed, expected: '%d', got: '%d'", expected, actual)
			}
		}
	}
}

func TestSmallValueSymbolSubtracted(t *testing.T) {
	tests := []romanTests{
		romanTests{"XLV", 45, nil},
		romanTests{"LXXXIX", 89, nil},
		romanTests{"MCMXLIV", 1944, nil},
	}

	roman := NewRoman()

	for _, tt := range tests {
		num, err := roman.ToNumber(tt.in)

		if err != tt.err {
			t.Error(err)
		} else {
			expected := tt.out
			actual := num
			if actual != expected {
				t.Errorf("TestSmallValueSymbolSubtracted failed, expected: '%d', got: '%d'", expected, actual)
			}
		}
	}
}

func TestCanBeRepeatedUntilThreeTimesForParticularSymbols(t *testing.T) {
	tests := []romanTests{
		romanTests{"III", 3, nil},
		romanTests{"XXXIX", 39, nil},
		romanTests{"CCCC", 0, errors.New("Cannot be repeat more then 3 times for this numerals: C")},
		romanTests{"CXXXX", 0, errors.New("Cannot be repeat more then 3 times for this numerals: X")},
		romanTests{"MMMM", 0, errors.New("Cannot be repeat more then 3 times for this numerals: M")},
	}

	roman := NewRoman()

	for _, tt := range tests {
		_, err := roman.ToNumber(tt.in)

		if err != nil {
			expected := tt.err.Error()
			actual := err.Error()
			if actual != expected {
				t.Errorf("TestCanBeRepeatedUntilThreeTimesForParticularSymbols failed, expected: '%s', got: '%s'", expected, actual)
			}
		}
	}
}

func TestCanNotBeRepeatedForParticularSymbols(t *testing.T) {
	tests := []romanTests{
		romanTests{"VVI", 0, errors.New("Cannot be repeat for this numerals: V")},
		romanTests{"LLX", 0, errors.New("Cannot be repeat for this numerals: L")},
		romanTests{"CDD", 0, errors.New("Cannot be repeat for this numerals: D")},
		romanTests{"CLLVV", 0, errors.New("Cannot be repeat for this numerals: L")},
	}

	roman := NewRoman()

	for _, tt := range tests {
		_, err := roman.ToNumber(tt.in)

		if err != nil {
			expected := tt.err.Error()
			actual := err.Error()
			if actual != expected {
				t.Errorf("TestCanBeRepeatedForParticularSymbols failed, expected: '%s', got: '%s'", expected, actual)
			}
		}
	}
}
