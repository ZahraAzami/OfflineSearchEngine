package LinearSearchEngineWithPosting

import (
	"OfflineSearchEngine/Internals/LinguisticModule"
	"OfflineSearchEngine/Internals/SearchEngines/Models"
	"bufio"
	"fmt"
	"sort"
	"strings"
)

type TermPosting struct {
	Term        string
	PostingList []Models.SearchRes
}

type SearchRes struct {
	DocId   int
	TermFrq int
}

func (t TermPosting) String() string {
	return fmt.Sprintf("{Term : %s, Postinglist : %v}", t.Term, t.PostingList)
}

type DumbEngine struct {
	data []TermPosting
}

func CreateDumbEngine(capacity int) DumbEngine {
	return DumbEngine{data: make([]TermPosting, 0, capacity)}
}
func (de *DumbEngine) findID(index int, docId int) int {
	for j := range de.data[index].PostingList {
		if docId == de.data[index].PostingList[j].DocId {
			return j
		}
	}
	return -1
}
func (de *DumbEngine) FindTerm(term string) int {
	for i := range de.data {
		if term == de.data[i].Term {
			return i
		}
	}
	return -1
}
func (de *DumbEngine) AddData(sd *bufio.Scanner, ID int) {
	info := TermPosting{}
	p := Models.SearchRes{}
	subTerm := ""
	for sd.Scan() {
		subTerm = sd.Text()
		subTerm = Linguistics.LinModule(subTerm)
		if subTerm == "" {
			continue
		}
		Out1 := de.FindTerm(subTerm)
		if Out1 == -1 {
			p = Models.SearchRes{
				DocId:   ID,
				TermFrq: 1,
			}
			info = TermPosting{
				Term: subTerm,
				PostingList: []Models.SearchRes{
					p,
				},
			}
			de.data = append(de.data, info)
		} else {
			out2 := de.findID(Out1, ID)
			if out2 == -1 {
				p = Models.SearchRes{
					DocId:   ID,
					TermFrq: 1,
				}
				de.data[Out1].PostingList = append(de.data[Out1].PostingList, p)
			} else {
				de.data[Out1].PostingList[out2].TermFrq++
			}
		}
		sort.SliceStable(de.data, func(i, j int) bool {
			return de.data[i].Term < de.data[j].Term
		})
	}
}
func (de *DumbEngine) Search(data string) []Models.SearchRes {
	SubData := strings.ToLower(data)
	result := make([]Models.SearchRes, 0)
	for i := range de.data {
		if SubData == de.data[i].Term {
			result = append(result, de.data[i].PostingList...)
		}
	}
	return result
}
