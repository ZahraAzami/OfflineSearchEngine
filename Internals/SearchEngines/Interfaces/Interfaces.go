package Interfaces

import (
	"OfflineSearchEngine/Internals/SearchEngines/Models"
	"bufio"
)

type ISearchEngine interface {
	AddData(sd *bufio.Scanner, ID int)
	Search(data string) []Models.SearchRes
}
