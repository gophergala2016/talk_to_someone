package utils

import "strings"

func StringStartWith(original, substring string) bool {
	if len(substring) > len(original) {
		return false
	}
	str := string(original[0:len(substring)])
	return str == substring
}

func GetIdAndName(text string) (string, string) {
	pos := strings.LastIndex(text, ":")
	return string(text[6:pos]), string(text[pos+1:])
}

func GetId(text string) string {
	var result string
	pos := strings.Index(text, ":")
	pos2 := strings.LastIndex(text, ":")
	if pos == pos2 {
		result = string(text[pos+1:])
		return result
	}
	result = string(text[pos+1 : pos2])
	return result
}
