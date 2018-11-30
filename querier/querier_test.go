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
		res := querier.isQueryType1(strings.Split(q.stmt, " "))
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
		res := querier.isQueryType2(strings.Split(q.stmt, " "))
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
		res := querier.isQueryType3(strings.Split(q.stmt, " "))
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
		res := querier.isQueryType4(strings.Split(q.stmt, " "))
		if q.isType != res {
			t.Errorf("TestQueryType4 failed, expected: '%t', got: '%t'", q.isType, res)
		}
	}
}
