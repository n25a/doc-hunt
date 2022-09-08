// Package parser used for parse input parameter
package parser

import "strings"

// ExportData used for export data from string to array of string
func ExportData(str string) []string {
	for strings.Contains(str, "  ") {
		str = strings.ReplaceAll(str, " ", "") // remove spaces
	}

	str = str[1 : len(str)-2] // remove brackets

	result := strings.Split(str, ",")

	return result
}
