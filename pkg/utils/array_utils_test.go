package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ContainsKey_Contains(t *testing.T) {
	// --- Given
	testMap := map[string][]string{
		"toto": nil,
		"titi": nil,
	}
	key := "toto"

	// --- When
	result := ContainsKey(testMap, key)

	// --- Then
	assert.Equal(t, result, true)
}

func Test_ContainsKey_NotContains(t *testing.T) {
	// --- Given
	testMap := map[string][]string{
		"toto": nil,
		"titi": nil,
	}
	key := "tata"

	// --- When
	result := ContainsKey(testMap, key)

	// --- Then
	assert.Equal(t, result, false)
}
