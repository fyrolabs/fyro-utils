package utils

import (
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

func TestRandNumeric(t *testing.T) {
	result := RandNumeric(10)
	assert.Equal(t, 10, len(result))

	for _, r := range result {
		assert.True(t, unicode.IsDigit(r))
	}
}
