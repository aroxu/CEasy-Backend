package utils

import "strings"

//GetLocationFromSJ returns location info of CBS Send Location
func GetLocationFromSJ(rawSJ string) string {
	return analyzeStringInBetween(rawSJ, "[", "]")
}

func analyzeStringInBetween(str string, start string, end string) (result string) {
	s := strings.Index(str, start)
	if s == -1 {
		return
	}
	s += len(start)
	e := strings.Index(str, end)
	if e == -1 {
		return
	}
	return str[s:e]
}
