package LinearFastAddEngine

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func TestLinearFastAddEngine_AddData(t *testing.T) {
	var tests = []struct {
		name     string
		data     string
		IdDoc    int
		expected []TermInfo
	}{
		{name: "test1", data: " the Golang is quick     Golang compilation Golang and execution speed;", IdDoc: 1,
			expected: []TermInfo{
				{"golang", 1}, {"is", 1}, {"quick", 1}, {"golang", 1},
				{"compilation", 1}, {"golang", 1}, {"execution", 1}, {"speed", 1},
			}},
		{name: "test2", data: " Golang the Golang is Golang memory Golang safety Golang;", IdDoc: 2,
			expected: []TermInfo{
				{"golang", 2}, {"golang", 2}, {"is", 2}, {"golang", 2},
				{"memory", 2}, {"golang", 2}, {"safety", 2}, {"golang", 2},
			}},
		{name: "test3", data: " the Golang is interfaces Golang enable Loosely couPled Golang systems;", IdDoc: 30,
			expected: []TermInfo{
				{"golang", 30}, {"is", 30}, {"interfaces", 30}, {"golang", 30},
				{"enable", 30}, {"loosely", 30}, {"coupled", 30}, {"golang", 30}, {"systems", 30},
			}},
		{name: "test4", data: " the Golang is interfaces Golang enable Loosely Golang couPled Golang systems Golang;", IdDoc: 310,
			expected: []TermInfo{
				{"golang", 310}, {"is", 310}, {"interfaces", 310}, {"golang", 310},
				{"enable", 310}, {"loosely", 310}, {"golang", 310}, {"coupled", 310}, {"golang", 310}, {"systems", 310}, {"golang", 310},
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
func TestLinearFastAddEngine_Search(t *testing.T) {
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
