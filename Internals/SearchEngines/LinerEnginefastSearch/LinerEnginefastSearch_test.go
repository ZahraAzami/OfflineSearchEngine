package LinerEnginefastSearch

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func TestLinerEngineFastSearch_AddData(t *testing.T) {

	var tests = []struct {
		name     string
		data     string
		IdDoc    int
		expected []TermInfo
	}{
		{name: "test1", data: " the golang is quick  Golang compilation Golang and execution speed;", IdDoc: 1,
			expected: []TermInfo{
				{"golang", 1, 3}, {"is", 1, 1}, {"quick", 1, 1},
				{"compilation", 1, 1}, {"execution", 1, 1}, {"speed", 1, 1},
			}},
		{name: "test2", data: " Golang the Golang is Golang memory Golang safety Golang;", IdDoc: 2,
			expected: []TermInfo{
				{term: "golang", docID: 2, termFrq: 5}, {term: "is", docID: 2, termFrq: 1}, {term: "memory", docID: 2, termFrq: 1}, {term: "safety", docID: 2, termFrq: 1},
			}},
		{name: "test3", data: " the Golang is interfaces Golang enable Loosely couPled Golang systems;", IdDoc: 30,
			expected: []TermInfo{
				{"golang", 30, 3}, {"is", 30, 1}, {"interfaces", 30, 1},
				{"enable", 30, 1}, {"loosely", 30, 1}, {"coupled", 30, 1}, {"systems", 30, 1},
			}},
		{name: "test4", data: " the Golang is interfaces Golang enable Loosely Golang couPled Golang systems Golang;", IdDoc: 310,
			expected: []TermInfo{

				{"golang", 310, 5}, {"is", 310, 1}, {"interfaces", 310, 1},
				{"enable", 310, 1}, {"loosely", 310, 1}, {"coupled", 310, 1}, {"systems", 310, 1},
			}},
	}
	for _, tt := range tests {
		ret := CreateDumbEngine(10000)
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
func TestLinerEngineFastSearch_Search(t *testing.T) {
	var tests = []struct {
		name       string
		data       string
		IdDoc      int
		dataSearch string
		expected   []SearchRes
	}{
		{name: "test1", data: " the golang is quick  Golang compilation Golang and execution speed;", IdDoc: 1, dataSearch: "golang",
			expected: []SearchRes{{1, 3}}},
		{name: "test2", data: "  Golang the Golang is is Golang memory Golang safety Golang;", IdDoc: 2, dataSearch: "is",
			expected: []SearchRes{{1, 1}, {2, 2}}},
		{name: "test3", data: " Golang the Golang is Golang memory Golang safety Golang;", IdDoc: 3, dataSearch: "safety",
			expected: []SearchRes{{2, 1}, {3, 1}}},
		{name: "test4", data: " the golang is quick  Golang compilation Golang and execution speed;", IdDoc: 4, dataSearch: "golang",
			expected: []SearchRes{{1, 3}, {2, 5}, {3, 5}, {4, 3}}},
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
