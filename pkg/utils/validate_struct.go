package utils

import (
	"fmt"
	"golang.org/x/exp/slices"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	checkFunctions = []func(tag reflect.StructTag, v reflect.Value, i int, fieldName string) *string{
		checkMin,
		checkMax,
		checkMinLength,
		checkMaxLength,
		checkPattern,
		checkDatePattern,
		checkEnum,
	}
)

// ValidateStruct validates that a struct meets all validation requirements defined by struct tags.
// It recursively validates nested structs and applies various validation rules including:
// required fields, min/max values, min/max lengths, pattern matching, and enum validation.
// Parameters:
//   - s: The struct to validate
//
// Returns:
//   - []string: A slice of error messages for validation failures, or nil if validation passes
func ValidateStruct(s interface{}) []string {
	var errorList []string
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag
		addedError := false

		if required, ok := tag.Lookup("required"); ok && required == "true" && isEmptyValue(v.Field(i)) {
			errorList = append(errorList, fmt.Sprintf("field %s is mandatory", field.Name))
			addedError = true
		}

		if field.Type.Kind() == reflect.Struct && !addedError {
			subErrorList := ValidateStruct(v.Field(i).Interface())
			if subErrorList != nil {
				errorList = append(errorList, subErrorList...)
				addedError = true
			}
		}

		for _, checkFunction := range checkFunctions {
			if !addedError {
				if errMsg := checkFunction(tag, v, i, field.Name); errMsg != nil {
					errorList = append(errorList, *errMsg)
					addedError = true
				}
			}
		}
	}

	return errorList
}

// checkMin validates that a numeric field value is greater than or equal to the minimum value
// specified in the struct tag. This validation is applied when a "min" tag is present.
// Parameters:
//   - tag: The struct tags for the field
//   - v: The reflect.Value of the struct being validated
//   - i: The index of the field being validated
//   - fieldName: The name of the field being validated
//
// Returns:
//   - *string: An error message if validation fails, nil otherwise
func checkMin(tag reflect.StructTag, v reflect.Value, i int, fieldName string) *string {
	if min, ok := tag.Lookup("min"); ok {
		if value, err := strconv.Atoi(min); err == nil {
			if v.Field(i).Int() < int64(value) {
				errMsg := fmt.Sprintf("the value of the field %s must be equal or over than %d", fieldName, value)
				return &errMsg
			}
		}
	}
	return nil
}

// checkMax validates that a numeric field value is less than or equal to the maximum value
// specified in the struct tag. This validation is applied when a "max" tag is present.
// Parameters:
//   - tag: The struct tags for the field
//   - v: The reflect.Value of the struct being validated
//   - i: The index of the field being validated
//   - fieldName: The name of the field being validated
//
// Returns:
//   - *string: An error message if validation fails, nil otherwise
func checkMax(tag reflect.StructTag, v reflect.Value, i int, fieldName string) *string {
	if max, ok := tag.Lookup("max"); ok {
		if value, err := strconv.Atoi(max); err == nil {
			if v.Field(i).Int() > int64(value) {
				errMsg := fmt.Sprintf("the value of the field %s must be equal or lower than %d", fieldName, value)
				return &errMsg
			}
		}
	}
	return nil
}

// checkMinLength validates that a string or slice field has a length greater than or equal to
// the minimum length specified in the struct tag. This validation is applied when a "minLength" tag is present.
// Parameters:
//   - tag: The struct tags for the field
//   - v: The reflect.Value of the struct being validated
//   - i: The index of the field being validated
//   - fieldName: The name of the field being validated
//
// Returns:
//   - *string: An error message if validation fails, nil otherwise
func checkMinLength(tag reflect.StructTag, v reflect.Value, i int, fieldName string) *string {
	if minLength, ok := tag.Lookup("minLength"); ok {
		if value, err := strconv.Atoi(minLength); err == nil {
			if v.Field(i).Type().Kind() == reflect.String && len(v.Field(i).String()) < value {
				errMsg := fmt.Sprintf("%s is too short. Min %d characters", fieldName, value)
				return &errMsg
			}
			if v.Field(i).Type().Kind() == reflect.Slice && v.Field(i).Len() < value {
				errMsg := fmt.Sprintf("%s is too short. Min %d entries", fieldName, value)
				return &errMsg
			}
		}
	}
	return nil
}

// checkMaxLength validates that a string or slice field has a length less than or equal to
// the maximum length specified in the struct tag. This validation is applied when a "maxLength" tag is present.
// Parameters:
//   - tag: The struct tags for the field
//   - v: The reflect.Value of the struct being validated
//   - i: The index of the field being validated
//   - fieldName: The name of the field being validated
//
// Returns:
//   - *string: An error message if validation fails, nil otherwise
func checkMaxLength(tag reflect.StructTag, v reflect.Value, i int, fieldName string) *string {
	if maxLength, ok := tag.Lookup("maxLength"); ok {
		if value, err := strconv.Atoi(maxLength); err == nil {
			if v.Field(i).Type().Kind() == reflect.String && len(v.Field(i).String()) > value {
				errMsg := fmt.Sprintf("%s is too long. Max %d characters", fieldName, value)
				return &errMsg
			}
			if v.Field(i).Type().Kind() == reflect.Slice && v.Field(i).Len() > value {
				errMsg := fmt.Sprintf("%s is too long. Max %d entries", fieldName, value)
				return &errMsg
			}
		}
	}
	return nil
}

// checkPattern validates that a string field matches a regular expression pattern
// specified in the struct tag. This validation is applied when a "pattern" tag is present.
// Parameters:
//   - tag: The struct tags for the field
//   - v: The reflect.Value of the struct being validated
//   - i: The index of the field being validated
//   - fieldName: The name of the field being validated
//
// Returns:
//   - *string: An error message if validation fails, nil otherwise
func checkPattern(tag reflect.StructTag, v reflect.Value, i int, fieldName string) *string {
	if pattern, ok := tag.Lookup("pattern"); ok {
		if !regexp.MustCompile(pattern).MatchString(v.Field(i).String()) {
			errMsg := fmt.Sprintf("the value of the field %s doesn't match the expected pattern", fieldName)
			return &errMsg
		}
	}
	return nil
}

// checkDatePattern validates that a string field can be parsed as a date using the format
// specified in the struct tag. This validation is applied when a "datePattern" tag is present.
// Parameters:
//   - tag: The struct tags for the field
//   - v: The reflect.Value of the struct being validated
//   - i: The index of the field being validated
//   - fieldName: The name of the field being validated
//
// Returns:
//   - *string: An error message if validation fails, nil otherwise
func checkDatePattern(tag reflect.StructTag, v reflect.Value, i int, fieldName string) *string {
	if datePattern, ok := tag.Lookup("datePattern"); ok {
		if _, err := time.Parse(datePattern, v.Field(i).String()); err != nil {
			errMsg := fmt.Sprintf("the value of the field %s doesn't match the expected pattern", fieldName)
			return &errMsg
		}
	}
	return nil
}

// checkEnum validates that a string or slice of strings contains only values that are
// included in the semicolon-separated list specified in the "enum" struct tag.
// Parameters:
//   - tag: The struct tags for the field
//   - v: The reflect.Value of the struct being validated
//   - i: The index of the field being validated
//   - fieldName: The name of the field being validated
//
// Returns:
//   - *string: An error message if validation fails, nil otherwise
func checkEnum(tag reflect.StructTag, v reflect.Value, i int, fieldName string) *string {
	if enum, ok := tag.Lookup("enum"); ok {
		if v.Field(i).Type().Kind() == reflect.Slice {
			for j := 0; j < v.Field(i).Len(); j++ {
				if !slices.Contains(strings.Split(enum, ";"), v.Field(i).Index(j).String()) {
					errMsg := fmt.Sprintf("field %s contains value which are not allowed - %s", fieldName, v.Field(i).Index(j).String())
					return &errMsg
				}
			}
		}
		if v.Field(i).Type().Kind() == reflect.String && !slices.Contains(strings.Split(enum, ";"), v.Field(i).String()) {
			errMsg := fmt.Sprintf("field %s contains value which are not allowed - %s", fieldName, v.Field(i).String())
			return &errMsg
		}
	}
	return nil
}

// isEmptyValue determines whether a reflect.Value should be considered empty.
// This is used by ValidateStruct to check if required fields have values.
// Different types have different criteria for emptiness:
// - Strings: empty if length is 0
// - Booleans: empty if false
// - Pointers, interfaces, slices: empty if nil
// Parameters:
//   - v: The reflect.Value to check
//
// Returns:
//   - bool: true if the value is considered empty, false otherwise
func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Ptr, reflect.Interface, reflect.Slice:
		return v.IsNil()
	}
	return false
}
