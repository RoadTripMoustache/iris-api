package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// -------------------- //
// ----- StrToInt ----- //
// -------------------- //
func Test_StrToInt_ok(t *testing.T) {
	parameters := []struct {
		input    string
		expected int
	}{
		{
			input:    "1",
			expected: 1,
		},
		{
			input:    "10",
			expected: 10,
		},
	}

	for i := range parameters {
		t.Run(fmt.Sprintf("StrToInt [%v]", parameters[i].input), func(t *testing.T) {
			actual := StrToInt(&parameters[i].input)
			if *actual != parameters[i].expected {
				t.Logf("expected%d: , actual:%d", parameters[i].expected, actual)
				t.Fail()
			}
		})
	}
}

func Test_StrToInt_nok(t *testing.T) {
	parameters := []struct {
		input string
	}{
		{
			input: "a",
		},
		{
			input: " ",
		},
		{
			input: "-",
		},
		{
			input: "",
		},
	}

	for i := range parameters {
		t.Run(fmt.Sprintf("StrToInt [%v]", parameters[i].input), func(t *testing.T) {
			actual := StrToInt(&parameters[i].input)
			if actual != nil {
				t.Logf("Expected: nil, actual: %d", actual)
				t.Fail()
			}
		})
	}
}

func Test_StrToInt_nok_stringNil(t *testing.T) {
	result := StrToInt(nil)

	assert.Nil(t, result)
}
