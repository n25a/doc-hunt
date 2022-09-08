package parser

import "strings"

func ExportData(str string) []string {
	for strings.Contains(str, "  ") {
		str = strings.ReplaceAll(str, " ", "") // remove spaces
	}

	str = str[1 : len(str)-2] // remove brackets

	result := strings.Split(str, ",")

	return result
}
