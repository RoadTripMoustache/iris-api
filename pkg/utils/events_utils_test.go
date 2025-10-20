package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GenerateVoyageDetailSessionId(t *testing.T) {
	// --- When
	result := GenerateVoyageDetailSessionId("toto", "titi")

	// --- Then
	assert.Equal(t, result, "toto-titi")
}
