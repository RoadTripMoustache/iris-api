// Package utils contains all utils methods for the whole API.
package utils

import "strconv"

// StrToInt - Convert a string to int. If the input value is nil or incorrect, it returns nil.
func StrToInt(s *string) *int {
	if s == nil {
		return nil
	}

	intValue, err := strconv.Atoi(*s)
	if err != nil {
		return nil
	}
	return &intValue
}
