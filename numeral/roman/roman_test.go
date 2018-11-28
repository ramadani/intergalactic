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
		romanTests{"MCMXLIV", 1944},
		romanTests{"MCMXXXIV", 1934},
		romanTests{"MCMIII", 1903},
	}

	roman := NewRoman()

	for _, tt := range tests {
		num, _ := roman.ToNumber(tt.in)

		expected := tt.out
		actual := num

		if actual != expected {
			t.Errorf("TestToNumber failed, expected: '%d', got: '%d'", expected, actual)
		}
	}
}
