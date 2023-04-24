package InvertedIndex

//import (
//	"bufio"
//	"reflect"
//	"strings"
//	"testing"
//)
//
//func TestInvertedIndex_AddData(t *testing.T) {
//	var tests = []struct {
//		name     string
//		data     string
//		IdDoc    int
//		expected DumbEngine
//	}{
//		{name: "test1", data: " the Golang is quick   Golang compilation Golang and execution speed;", IdDoc: 1,
//			expected: DumbEngine{
//				map[string][]SearchRes{"compilation": {{1, 1}}, "execution": {{1, 1}}, "golang": {{1, 3}}, "is": {{1, 1}}, "quick": {{1, 1}}, "speed": {{1, 1}}},
//			}},
//
//		{name: "test2", data: " Golang the Golang is Golang memory Golang safety Golang;", IdDoc: 2,
//			expected: DumbEngine{
//				map[string][]SearchRes{"compilation": {{1, 1}}, "execution": {{1, 1}}, "golang": {{1, 3}, {2, 5}}, "is": {{1, 1}, {2, 1}},
//					"memory": {{2, 1}}, "quick": {{1, 1}}, "safety": {{2, 1}}, "speed": {{1, 1}}},
//			}},
//		{name: "test3", data: " Golang the is Golang memory Golang safety Golang;", IdDoc: 35,
//			expected: DumbEngine{
//				map[string][]SearchRes{"compilation": {{1, 1}}, "execution": {{1, 1}}, "golang": {{1, 3}, {2, 5}, {35, 4}}, "is": {{1, 1}, {2, 1}, {35, 1}},
//					"memory": {{2, 1}, {35, 1}}, "quick": {{1, 1}}, "safety": {{2, 1}, {35, 1}}, "speed": {{1, 1}}},
//			}},
//	}
//	ret := CreateInvertedIndex(10000)
//	for _, tt := range tests {
//
//		t.Run(tt.name, func(t *testing.T) {
//			scanner := bufio.NewScanner(strings.NewReader(tt.data))
//			scanner.Split(bufio.ScanWords)
//			ret.AddData(scanner, tt.IdDoc)
//			if !reflect.DeepEqual(tt.expected, ret) {
//				t.Errorf("\ngot:{%v}\n but expected:%v\n", ret, tt.expected)
//			}
//		})
//	}
//}
//func TestLinearFastAddEngine_Search(t *testing.T) {
//	var tests = []struct {
//		name       string
//		data       string
//		IdDoc      int
//		dataSearch string
//		expected   []SearchRes
//	}{
//		{name: "test1", data: " the golang is quick  Golang compilation Golang and execution speed;", IdDoc: 1, dataSearch: "golang",
//			expected: []SearchRes{{1, 3}}},
//		{name: "test2", data: "  Golang the Golang is is Golang memory Golang safety Golang;", IdDoc: 2, dataSearch: "is",
//			expected: []Posting{{1, 1}, {2, 2}}},
//		{name: "test3", data: " Golang the Golang is Golang memory Golang safety Golang;", IdDoc: 3, dataSearch: "safety",
//			expected: []Posting{{2, 1}, {3, 1}}},
//		{name: "test4", data: " the golang is quick  Golang compilation Golang and execution speed;", IdDoc: 4, dataSearch: "golang",
//			expected: []Posting{{1, 3}, {2, 5}, {3, 5}, {4, 3}}},
//	}
//	ret := CreateInvertedIndex(100000)
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			scanner := bufio.NewScanner(strings.NewReader(tt.data))
//			scanner.Split(bufio.ScanWords)
//			ret.AddData(scanner, tt.IdDoc)
//			res, _ := ret.Search(tt.dataSearch)
//			if !reflect.DeepEqual(tt.expected, res) {
//				t.Errorf("\ngot : %v\nbut expected : %v\n", res, tt.expected)
//			}
//
//		})
//	}
//}
