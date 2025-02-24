package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	var tests = []struct {
		N        int
		M        int
		expected string
	}{
		// Standard Cases
		{1, 5, "1,2,Fizz,4,Buzz"},
		{10, 15, "Buzz,11,Fizz,13,14,FizzBuzz"},
		{3, 3, "Fizz"},
		{5, 5, "Buzz"},
		{15, 15, "FizzBuzz"},

		// Edge Cases
		{1, 1, "1"},
		{0, 0, "FizzBuzz"},
		{-3, -3, "Fizz"},
		{-5, -5, "Buzz"},
		{-15, -15, "FizzBuzz"},
		{-5, 5, "Buzz,-4,Fizz,-2,-1,FizzBuzz,1,2,Fizz,4,Buzz"},

		// Invalid Cases
		{5, 1, "Error: N must be less than or equal to M"},
	}

	for _, test := range tests {
		output := run(test.N, test.M)
		assert.Equal(t, test.expected, output, "For input (%d, %d), expected %s but got %s", test.N, test.M, test.expected, output)
	}
}
