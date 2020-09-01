package utils

import "strings"

func ParseStringInBetween(origin string, start string, end string) (result string) {
	modifiedStart := strings.Index(origin, start)
	if modifiedStart == -1 {
		return
	}
	modifiedStart += len(start)
	modifiedEnd := strings.Index(origin, end)
	if modifiedEnd == -1 {
		return
	}
	return origin[modifiedStart:modifiedEnd]
}
