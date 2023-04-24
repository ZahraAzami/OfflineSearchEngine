package Linguistics

import "strings"

func ToLowerData(SubTerm string) string {
	return strings.ToLower(SubTerm)
}
func PunctuationRemove(SubTerm string) string {
	CorrectiveWords := []string{".", ";", "!", "?", ":", "[", "]", "{", "}", ",", "(", ")"}
	for i := range CorrectiveWords {
		SubTerm = strings.TrimSuffix(SubTerm, CorrectiveWords[i])
	}
	return SubTerm
}
func CheckStopWork(subTerm string) string {

	if subTerm == "the" || subTerm == "on" || subTerm == "a" || subTerm == "an" || subTerm == "of" || subTerm == "and" {
		return ""
	}
	return subTerm
}
func LinModule(subTerm string) string {
	subTerm = ToLowerData(subTerm)
	subTerm = PunctuationRemove(subTerm)
	subTerm = CheckStopWork(subTerm)
	if subTerm == "" {
		return ""
	}
	return subTerm
}
