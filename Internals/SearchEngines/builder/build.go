package builder

import (
	"OfflineSearchEngine/Internals/SearchEngines/Interfaces"
	"OfflineSearchEngine/Internals/SearchEngines/LinearFastAddEngine"
)

func NewSearchEngine(str string, cap int) Interfaces.ISearchEngine {
	switch str {
	case "LinearFastAddEngine":
		ret := LinearFastAddEngine.CreateLinearFastAddEngine(cap)
		return &ret
	default:
		return nil
	}
}
