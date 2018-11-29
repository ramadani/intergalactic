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

func TestCanNotBeRepeatedForParticularSymbols(t *testing.T) {
	tests := []romanTests{
		romanTests{"VVI", 0, errors.New("Cannot be repeat for this numerals: V")},
		romanTests{"LLX", 0, errors.New("Cannot be repeat for this numerals: L")},
		romanTests{"CDD", 0, errors.New("Cannot be repeat for this numerals: D")},
	}

	roman := NewRoman()

	for _, tt := range tests {
		_, err := roman.ToNumber(tt.in)

		expected := tt.err.Error()
		actual := err.Error()
		if actual != expected {
			t.Errorf("TestCanBeRepeatedForParticularSymbols failed, expected: '%s', got: '%s'", expected, actual)
		}
	}
}
