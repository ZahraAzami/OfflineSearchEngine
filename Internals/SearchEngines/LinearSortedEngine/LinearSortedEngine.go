package LinearSortedEngine

import (
	"OfflineSearchEngine/Internals/LinguisticModule"
	"bufio"
	"fmt"
	"sort"
	"strings"
)

type TermInfo struct {
	term    string
	docID   int
	termFrq int
}
type SearchRes struct {
	docId   int
	termFrq int
}

func (t TermInfo) String() string {
	return fmt.Sprintf("{Term : %s, DocId: %d, termFrq: %d}", t.term, t.docID, t.termFrq)
}

type DumbEngine struct {
	data []TermInfo
}

func CreateDumbEngine(capacity int) DumbEngine {
	return DumbEngine{data: make([]TermInfo, 0, capacity)}
}
func (de *DumbEngine) find(term string, docId int) int {
	for i := range de.data {
		if (term == de.data[i].term) && (docId == de.data[i].docID) {
			return i
		}
	}
	return -1
}
func (de *DumbEngine) AddData(sd *bufio.Scanner, ID int) {
	info := TermInfo{}
	subTerm := ""
	for sd.Scan() {
		subTerm = sd.Text()
		subTerm = Linguistics.LinModule(subTerm)
		if subTerm == "" {
			continue
		}
		Out := de.find(subTerm, ID)
		if Out == -1 {
			info = TermInfo{
				term:    subTerm,
				docID:   ID,
				termFrq: 1,
			}
			de.data = append(de.data, info)
		} else {
			de.data[Out].termFrq++
		}
		sort.SliceStable(de.data, func(i, j int) bool {
			return de.data[i].term < de.data[j].term
		})
	}
}
func (de *DumbEngine) Search(data string) []SearchRes {
	SubData := strings.ToLower(data)
	result := make([]SearchRes, 0)
	for i := range de.data {
		if SubData == de.data[i].term {
			info := SearchRes{
				docId:   de.data[i].docID,
				termFrq: de.data[i].termFrq,
			}
			result = append(result, info)
		}
	}
	return result
}
