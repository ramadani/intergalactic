package roman

import (
	"testing"
)

type romanTests struct {
	in  string
	out int
}

func TestToNumber(t *testing.T) {
	tests := []romanTests{
		romanTests{"XXXIX", 39},
		romanTests{"MVI", 1006},
		romanTests{"MMVI", 2006},
	}

	roman := NewRoman()

	for _, tt := range tests {
		num, err := roman.ToNumber(tt.in)
		if err != nil {
			t.Error(err)
		}

		expected := tt.out
		actual := num
		if actual != expected {
			t.Errorf("TestToNumber failed, expected: '%d', got: '%d'", expected, actual)
		}
	}
}

func TestSmallValueSymbolSubtracted(t *testing.T) {
	tests := []romanTests{
		romanTests{"XLV", 45},
		romanTests{"LXXXIX", 89},
		romanTests{"MCMXLIV", 1944},
	}

	roman := NewRoman()

	for _, tt := range tests {
		num, err := roman.ToNumber(tt.in)
		if err != nil {
			t.Error(err)
		}

		expected := tt.out
		actual := num
		if actual != expected {
			t.Errorf("TestToNumber failed, expected: '%d', got: '%d'", expected, actual)
		}
	}
}
