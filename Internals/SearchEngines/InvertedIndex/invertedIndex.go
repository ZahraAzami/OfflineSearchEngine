package InvertedIndex

//import (
//	"OfflineSearchEngine/Internals/LinguisticModule"
//	"OfflineSearchEngine/Internals/SearchEngines/Models"
//	"bufio"
//	"strings"
//)
//
//type DumbEngine struct {
//	data map[string][]Models.SearchRes
//}
//type SearchRes struct {
//	docId   int
//	termFrq int
//}
//
//func CreateInvertedIndex(cap int) DumbEngine {
//	return DumbEngine{data: make(map[string][]Models.SearchRes, cap)}
//}
//func (de *DumbEngine) FindId(term string, docId int) int {
//	for key, value := range de.data {
//		for i := range value {
//			if value[i].docId == docId && key == term {
//				return i
//			}
//		}
//	}
//	return -1
//}
//func (de *DumbEngine) AddData(sd *bufio.Scanner, ID int) {
//	p := Models.SearchRes{}
//	subTerm := ""
//	for sd.Scan() {
//		subTerm = sd.Text()
//		subTerm = Linguistics.LinModule(subTerm)
//		if subTerm == "" {
//			continue
//		}
//		_, Ok := de.data[subTerm]
//		if Ok == false {
//			p = Models.SearchRes{
//				DocId:   ID,
//				TermFrq: 1,
//			}
//			de.data[subTerm] = append(de.data[subTerm], p)
//		} else {
//			Out2 := de.FindId(subTerm, ID)
//			if Out2 == -1 {
//				p1 := SearchRes{
//					docId:   ID,
//					termFrq: 1,
//				}
//				de.data[subTerm] = append(de.data[subTerm], p1)
//			} else {
//				de.data[subTerm][Out2].TermFrq++
//			}
//		}
//	}
//}
//func (de *DumbEngine) Search(data string) []SearchRes {
//	SubData := strings.ToLower(data)
//	Search, _ := de.data[SubData]
//	result := Search
//	return result
//}
