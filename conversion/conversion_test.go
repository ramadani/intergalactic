package conversion

import (
	"errors"
	"testing"

	"github.com/ramadani/intergalactic/credit"
	"github.com/ramadani/intergalactic/querier"

	"github.com/ramadani/intergalactic/converter"
	"github.com/ramadani/intergalactic/numeral/roman"
)

func TestQueryFromType1(t *testing.T) {
	type queryTest1 struct {
		stmt string
		err  error
	}

	tests := []queryTest1{
		queryTest1{"glob is I", nil},
		queryTest1{"prok is V", nil},
		queryTest1{"pish is X", nil},
		queryTest1{"tegj is L", nil},
		queryTest1{"prok is V", errors.New("prok unit is exists")},
		queryTest1{"pish is X", errors.New("pish unit is exists")},
	}

	converter := converter.NewConverter(roman.NewRoman())
	conversion := NewConversion(querier.NewQuerier(), converter, credit.NewCredit())

	for _, tt := range tests {
		res, err := conversion.Query(tt.stmt)

		if res != "" {
			t.Errorf("TestQueryFromType1 failed, expected: empty, got: '%s'", res)
		} else if err != nil && err.Error() != tt.err.Error() {
			t.Errorf("TestQueryFromType1 failed, expected: '%s', got: '%s'", tt.err.Error(), err.Error())
		}
	}
}

func TestQueryFromType2(t *testing.T) {
	type queryTest2 struct {
		stmt string
		err  error
	}

	tests := []queryTest2{
		queryTest2{"glob glob Silver is 34 Credits", nil},
		queryTest2{"glob prok Gold is 57800 Credits", nil},
		queryTest2{"pish pish Iron is 3910 Credits", nil},
		queryTest2{"pish great Silver is 5000 Credits", errors.New("Unit not found for great")},
		queryTest2{"pish nice good Gold is 7000 Credits", errors.New("Unit not found for nice")},
		queryTest2{"glob prok Gold is 60000 Credits", errors.New("Gold credit is exists")},
	}

	converter := converter.NewConverter(roman.NewRoman())
	conversion := NewConversion(querier.NewQuerier(), converter, credit.NewCredit())

	// init units from query type 1
	conversion.Query("glob is I")
	conversion.Query("prok is V")
	conversion.Query("pish is X")
	conversion.Query("tegj is L")

	for _, tt := range tests {
		res, err := conversion.Query(tt.stmt)

		if res != "" {
			t.Errorf("TestQueryFromType2 failed, expected: empty, got: '%s'", res)
		} else if err != nil && err.Error() != tt.err.Error() {
			t.Errorf("TestQueryFromType2 failed, expected: '%s', got: '%s'", tt.err.Error(), err.Error())
		}
	}
}

func TestQueryFromType3(t *testing.T) {
	type queryTest3 struct {
		stmt string
		res  string
		err  error
	}

	tests := []queryTest3{
		queryTest3{"how much is prok glob glob ?", "prok glob glob is 7", nil},
		queryTest3{"how much is pish tegj glob glob ?", "pish tegj glob glob is 42", nil},
		queryTest3{"how much is tegj prok glob ?", "tegj prok glob is 56", nil},
		queryTest3{"how much is great prok glob ?", "", errors.New("Unit not found for great")},
		queryTest3{"how much is nice prok glob glob ?", "", errors.New("Unit not found for nice")},
	}

	converter := converter.NewConverter(roman.NewRoman())
	conversion := NewConversion(querier.NewQuerier(), converter, credit.NewCredit())

	// init units from query type 1
	conversion.Query("glob is I")
	conversion.Query("prok is V")
	conversion.Query("pish is X")
	conversion.Query("tegj is L")

	for _, tt := range tests {
		res, err := conversion.Query(tt.stmt)

		if err != nil && err.Error() != tt.err.Error() {
			t.Errorf("TestQueryFromType3 failed, expected: '%s', got: '%s'", tt.err.Error(), err.Error())
		} else if res != tt.res {
			t.Errorf("TestQueryFromType3 failed, expected: '%s', got: '%s'", tt.res, res)
		}
	}
}

func TestQueryFromType4(t *testing.T) {
	type queryTest4 struct {
		stmt string
		res  string
		err  error
	}

	tests := []queryTest4{
		queryTest4{"how many Credits is glob prok Silver ?", "glob prok Silver is 68 Credits", nil},
		queryTest4{"how many Credits is glob prok Gold ?", "glob prok Gold is 57800 Credits", nil},
		queryTest4{"how many Credits is glob prok Iron ?", "glob prok Iron is 782 Credits", nil},
		queryTest4{"how many Credits is glob prok Platinum ?", "", errors.New("Platinum credit type is not found")},
	}

	converter := converter.NewConverter(roman.NewRoman())
	conversion := NewConversion(querier.NewQuerier(), converter, credit.NewCredit())

	// init units from query type 1
	conversion.Query("glob is I")
	conversion.Query("prok is V")
	conversion.Query("pish is X")
	conversion.Query("tegj is L")
	// init credit type from query type 2
	conversion.Query("glob glob Silver is 34 Credits")
	conversion.Query("glob prok Gold is 57800 Credits")
	conversion.Query("pish pish Iron is 3910 Credits")

	for _, tt := range tests {
		res, err := conversion.Query(tt.stmt)

		if err != nil && err.Error() != tt.err.Error() {
			t.Errorf("TestQueryFromType3 failed, expected: '%s', got: '%s'", tt.err.Error(), err.Error())
		} else if res != tt.res {
			t.Errorf("TestQueryFromType3 failed, expected: '%s', got: '%s'", tt.res, res)
		}
	}
}

func TestQueryFromUnknown(t *testing.T) {
	type queryTestUnknown struct {
		stmt string
		res  string
	}

	tests := []queryTestUnknown{
		queryTestUnknown{
			"how much wood could a woodchuck chuck if a woodchuck could chuck wood ?",
			"I have no idea what you are talking about",
		},
		queryTestUnknown{
			"how much wood great good is best ever ?",
			"I have no idea what you are talking about",
		},
	}

	converter := converter.NewConverter(roman.NewRoman())
	conversion := NewConversion(querier.NewQuerier(), converter, credit.NewCredit())

	// init units from query type 1
	conversion.Query("glob is I")
	conversion.Query("prok is V")
	conversion.Query("pish is X")
	conversion.Query("tegj is L")
	// init credit type from query type 2
	conversion.Query("glob glob Silver is 34 Credits")
	conversion.Query("glob prok Gold is 57800 Credits")
	conversion.Query("pish pish Iron is 3910 Credits")

	for _, tt := range tests {
		res, _ := conversion.Query(tt.stmt)

		if res != tt.res {
			t.Errorf("TestQueryFromType3 failed, expected: '%s', got: '%s'", tt.res, res)
		}
	}
}
