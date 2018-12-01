package querier

import (
	"strings"
	"testing"
)

func TestQueryType1(t *testing.T) {
	type query struct {
		stmt   string
		isType bool
	}

	queries := []query{
		query{"glob is I", true},
		query{"pish is X", true},
		query{"glob glob Silver is 34 Credits", false},
		query{"pish pish Iron is 3910 Credits", false},
		query{"how much is pish tegj glob glob ?", false},
		query{"how many Credits is glob prok Silver ?", false},
		query{"how many Credits is glob prok Iron ?", false},
		query{"how much wood could a woodchuck chuck if a woodchuck could chuck wood ?", false},
	}

	querier := NewQuerier()

	for _, q := range queries {
		res := querier.IsType1(strings.Split(q.stmt, " "))
		if q.isType != res {
			t.Errorf("TestQueryType1 failed, expected: '%t', got: '%t'", q.isType, res)
		}
	}
}

func TestQueryType2(t *testing.T) {
	type query struct {
		stmt   string
		isType bool
	}

	queries := []query{
		query{"glob is I", false},
		query{"pish is X", false},
		query{"glob glob Silver is 34 Credits", true},
		query{"pish pish Iron is 3910 Credits", true},
		query{"how much is pish tegj glob glob ?", false},
		query{"how many Credits is glob prok Silver ?", false},
		query{"how many Credits is glob prok Iron ?", false},
		query{"how much wood could a woodchuck chuck if a woodchuck could chuck wood ?", false},
	}

	querier := NewQuerier()

	for _, q := range queries {
		res := querier.IsType2(strings.Split(q.stmt, " "))
		if q.isType != res {
			t.Errorf("TestQueryType2 failed, expected: '%t', got: '%t'", q.isType, res)
		}
	}
}

func TestQueryType3(t *testing.T) {
	type query struct {
		stmt   string
		isType bool
	}

	queries := []query{
		query{"glob is I", false},
		query{"pish is X", false},
		query{"glob glob Silver is 34 Credits", false},
		query{"pish pish Iron is 3910 Credits", false},
		query{"how much is pish tegj glob glob ?", true},
		query{"how many Credits is glob prok Silver ?", false},
		query{"how many Credits is glob prok Iron ?", false},
		query{"how much wood could a woodchuck chuck if a woodchuck could chuck wood ?", false},
	}

	querier := NewQuerier()

	for _, q := range queries {
		res := querier.IsType3(strings.Split(q.stmt, " "))
		if q.isType != res {
			t.Errorf("TestQueryType3 failed, expected: '%t', got: '%t'", q.isType, res)
		}
	}
}

func TestQueryType4(t *testing.T) {
	type query struct {
		stmt   string
		isType bool
	}

	queries := []query{
		query{"glob is I", false},
		query{"pish is X", false},
		query{"glob glob Silver is 34 Credits", false},
		query{"pish pish Iron is 3910 Credits", false},
		query{"how much is pish tegj glob glob ?", false},
		query{"how many Credits is glob prok Silver ?", true},
		query{"how many Credits is glob prok Iron ?", true},
		query{"how much wood could a woodchuck chuck if a woodchuck could chuck wood ?", false},
	}

	querier := NewQuerier()

	for _, q := range queries {
		res := querier.IsType4(strings.Split(q.stmt, " "))
		if q.isType != res {
			t.Errorf("TestQueryType4 failed, expected: '%t', got: '%t'", q.isType, res)
		}
	}
}

func TestGetQueryType1Values(t *testing.T) {
	type queryType1Value struct {
		stmt  string
		alias string
		num   string
	}

	tests := []queryType1Value{
		queryType1Value{"glob is I", "glob", "I"},
		queryType1Value{"prok is V", "prok", "V"},
		queryType1Value{"pish is X", "pish", "X"},
		queryType1Value{"tegj is L", "tegj", "L"},
	}

	querier := NewQuerier()

	for _, tt := range tests {
		alias, num := querier.GetType1Values(strings.Split(tt.stmt, " "))

		if alias != tt.alias || num != tt.num {
			t.Errorf("TestGetQueryType1Values failed, expected: '%s %s', got: '%s %s'",
				tt.alias, tt.num, alias, num)
		}
	}
}

func TestGetQueryType2Values(t *testing.T) {
	type queryType2Value struct {
		stmt       string
		aliases    string
		creditType string
		total      int
	}

	tests := []queryType2Value{
		queryType2Value{"glob glob Silver is 34 Credits", "glob glob", "Silver", 34},
		queryType2Value{"glob prok Gold is 57800 Credits", "glob prok", "Gold", 57800},
		queryType2Value{"pish pish Iron is 3910 Credits", "pish pish", "Iron", 3910},
	}

	querier := NewQuerier()

	for _, tt := range tests {
		aliases, creditType, total := querier.GetType2Values(strings.Split(tt.stmt, " "))

		if aliases != tt.aliases || creditType != tt.creditType || total != tt.total {
			t.Errorf("TestGetQueryType2Values failed, expected: '%s %s %d', got: '%s %s %d'",
				tt.aliases, tt.creditType, tt.total, aliases, creditType, total)
		}
	}
}

func TestGetQueryType3Values(t *testing.T) {
	type queryType3Value struct {
		stmt  string
		alias string
	}

	tests := []queryType3Value{
		queryType3Value{"how much is tegj pish ?", "tegj pish"},
		queryType3Value{"how much is pish tegj glob glob ?", "pish tegj glob glob"},
	}

	querier := NewQuerier()

	for _, tt := range tests {
		alias := querier.GetType3Values(strings.Split(tt.stmt, " "))

		if alias != tt.alias {
			t.Errorf("TestGetQueryType3Values failed, expected: '%s', got: '%s'", tt.alias, alias)
		}
	}
}

func TestGetQueryType4Values(t *testing.T) {
	type queryType4Value struct {
		stmt       string
		alias      string
		creditType string
	}

	tests := []queryType4Value{
		queryType4Value{"how many Credits is glob prok Silver ?", "glob prok", "Silver"},
		queryType4Value{"how many Credits is pish prok Gold ?", "pish prok", "Gold"},
		queryType4Value{"how many Credits is glob pish Iron ?", "glob pish", "Iron"},
	}

	querier := NewQuerier()

	for _, tt := range tests {
		alias, creditType := querier.GetType4Values(strings.Split(tt.stmt, " "))

		if alias != tt.alias || creditType != tt.creditType {
			t.Errorf("TestGetQueryType3Values failed, expected: '%s %s', got: '%s %s'",
				tt.alias, tt.creditType, alias, creditType)
		}
	}
}
