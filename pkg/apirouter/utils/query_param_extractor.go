package utils

import (
	"strings"
)

// QueryParamExtractor - Returns the value of the query param given in parameter, otherwise return nil.
func QueryParamExtractor(queryParams map[string][]string, paramName string) *string {
	if value, ok := queryParams[paramName]; ok {
		queryParamValue := strings.Join(value, ",")
		return &queryParamValue
	}
	return nil
}
