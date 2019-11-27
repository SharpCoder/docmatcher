package utils

import (
	"strings"
)

func PascalToWords(token string) []string {
	var result []string
	part := ""

	for i := range token {
		char := string(token[i])
		isNewToken := string(char) == strings.ToUpper(char)
		if isNewToken && len(part) > 0 {
			result = append(result, part)
			part = ""
		}

		part += char
	}

	if len(part) > 0 {
		result = append(result, part)
	}

	return result
}

func Ngram(vector []string) []string {
	var results []string
	for i := range vector {
		if i < len(vector) - 1 {
			results = append(results, vector[i] + vector[i + 1])
		}
	}
	return results
}
