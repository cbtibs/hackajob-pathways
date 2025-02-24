package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{2, 4, 5, 4, 2},
			expected: 5,
		},
		{
			input:    []int{7},
			expected: 7,
		},
		{
			input:    []int{1, 1},
			expected: 0,
		},
	}

	for _, test := range tests {
		output := run(test.input)
		assert.Equal(t, test.expected, output, "For input %v, expected %d but got %d", test.input, test.expected, output)
	}
}
