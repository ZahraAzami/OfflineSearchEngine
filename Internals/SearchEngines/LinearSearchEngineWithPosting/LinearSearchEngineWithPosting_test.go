package LinearSearchEngineWithPosting

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func TestLinearSearchEngineWithPosting_AddData(t *testing.T) {

	var tests = []struct {
		name     string
		data     string
		IdDoc    int
		expected []TermPosting
	}{
		{name: "test1", data: " the golang is quick  Golang compilation Golang and execution speed;", IdDoc: 1,
			expected: []TermPosting{
				{"compilation", []Posting{{1, 1}}}, {"execution", []Posting{{1, 1}}},
				{"golang", []Posting{{1, 3}}}, {"is", []Posting{{1, 1}}},
				{"quick", []Posting{{1, 1}}}, {"speed", []Posting{{1, 1}}},
			}},
		{name: "test2", data: "  Golang the Golang is is Golang memory Golang safety Golang;", IdDoc: 2,
			expected: []TermPosting{
				{"compilation", []Posting{{1, 1}}}, {"execution", []Posting{{1, 1}}},
				{"golang", []Posting{{1, 3}, {2, 5}}}, {"is", []Posting{{1, 1}, {2, 2}}},
				{"memory", []Posting{{2, 1}}}, {"quick", []Posting{{1, 1}}}, {"safety", []Posting{{2, 1}}}, {"speed", []Posting{{1, 1}}},
			}},
		{name: "test3", data: " Golang the Golang is Golang memory Golang safety Golang;", IdDoc: 3,
			expected: []TermPosting{
				{"compilation", []Posting{{1, 1}}}, {"execution", []Posting{{1, 1}}},
				{"golang", []Posting{{1, 3}, {2, 5}, {3, 5}}}, {"is", []Posting{{1, 1}, {2, 2}, {3, 1}}},
				{"memory", []Posting{{2, 1}, {3, 1}}}, {"quick", []Posting{{1, 1}}}, {"safety", []Posting{{2, 1}, {3, 1}}}, {"speed", []Posting{{1, 1}}},
			}},
	}
	ret := CreateDumbEngine(100000)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(tt.data))
			scanner.Split(bufio.ScanWords)
			ret.AddData(scanner, tt.IdDoc)
			if !reflect.DeepEqual(tt.expected, ret.data) {
				t.Errorf("\ngot : %s\nbut expected : %s\n", ret.data, tt.expected)
			}

		})
	}
}
func TestLinearSearchEngineWithPosting_Search(t *testing.T) {
	var tests = []struct {
		name       string
		data       string
		IdDoc      int
		dataSearch string
		expected   []Posting
	}{
		{name: "test1", data: " the golang is quick  Golang compilation Golang and execution speed;", IdDoc: 1, dataSearch: "golang",
			expected: []Posting{{1, 3}}},
		{name: "test2", data: "  Golang the Golang is is Golang memory Golang safety Golang;", IdDoc: 2, dataSearch: "is",
			expected: []Posting{{1, 1}, {2, 2}}},
		{name: "test3", data: " Golang the Golang is Golang memory Golang safety Golang;", IdDoc: 3, dataSearch: "safety",
			expected: []Posting{{2, 1}, {3, 1}}},
		{name: "test4", data: " the golang is quick  Golang compilation Golang and execution speed;", IdDoc: 4, dataSearch: "golang",
			expected: []Posting{{1, 3}, {2, 5}, {3, 5}, {4, 3}}},
	}
	ret := CreateDumbEngine(100000)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(tt.data))
			scanner.Split(bufio.ScanWords)
			ret.AddData(scanner, tt.IdDoc)
			res := ret.Search(tt.dataSearch)
			if !reflect.DeepEqual(tt.expected, res) {
				t.Errorf("\ngot : %v\nbut expected : %v\n", res, tt.expected)
			}

		})
	}
}
