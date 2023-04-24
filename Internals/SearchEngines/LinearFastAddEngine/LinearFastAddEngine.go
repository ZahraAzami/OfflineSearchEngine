package LinearFastAddEngine

import (
	"OfflineSearchEngine/Internals/LinguisticModule"
	"OfflineSearchEngine/Internals/SearchEngines/Models"
	"bufio"
	"fmt"
	"strings"
)

type TermInfo struct {
	Term  string
	DocID int
}

func (t TermInfo) String() string {
	return fmt.Sprintf("{Term : %s, DocId: %d}", t.Term, t.DocID)
}

type LinearFastAddEngine struct {
	data []TermInfo
}
type Search struct {
	data []Models.SearchRes
}

func CreateLinearFastAddEngine(capacity int) LinearFastAddEngine {
	return LinearFastAddEngine{data: make([]TermInfo, 0, capacity)}
}
func (le *LinearFastAddEngine) AddData(sd *bufio.Scanner, ID int) {
	for sd.Scan() {
		subTerm := sd.Text()
		subTerm = Linguistics.LinModule(subTerm)
		if subTerm == "" {
			continue
		}
		info := TermInfo{
			Term:  subTerm,
			DocID: ID,
		}
		le.data = append(le.data, info)
	}
}
func (S *Search) findIndex(index int) int {
	for i := range S.data {
		if S.data[i].DocId == index {
			return i
		}
	}
	return -1
}
func (le *LinearFastAddEngine) Search(data string) []Models.SearchRes {
	SubData := strings.ToLower(data)
	res := Models.SearchRes{}
	result := Search{}
	for i := range le.data {
		if SubData == le.data[i].Term {
			Out := result.findIndex(le.data[i].DocID)
			if Out == -1 {
				res = Models.SearchRes{
					DocId:   le.data[i].DocID,
					TermFrq: 1,
				}
				result.data = append(result.data, res)
			} else {
				result.data[Out].TermFrq++
			}
		}
	}
	return result.data
}

//o (log n )--->
//
//func (de *DumbEngine) SearchWord1(data string) []SearchRes {
//	SubData := strings.ToLower(data)
//	countNum := make([]int, 0)
//	numData := 0
//	result := make([]SearchRes, 0)
//	for i := 0; i < len(de.data); i++ {
//		if de.data[i].term == SubData {
//			countNum = append(countNum, de.data[i].docID)
//		}
//	}
//	sort.Ints(countNum)
//	res := SearchRes{}
//	for i := 0; i < len(countNum); i++ {
//		if i == len(countNum)-1 {
//			numData++
//			res = SearchRes{
//				docId:   countNum[i],
//				termFrq: numData,
//			}
//			result = append(result, res)
//			break
//		}
//		if countNum[i] != countNum[i+1] {
//			numData++
//			res = SearchRes{
//				docId:   countNum[i],
//				termFrq: numData,
//			}
//			result = append(result, res)
//			numData = 0
//			continue
//		}
//		numData++
//	}
//	return result
//}
