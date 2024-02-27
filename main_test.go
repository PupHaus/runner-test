package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFibonacciSequence(t *testing.T) {
	// Define test cases in a table
	testCases := []struct {
		name     string
		start    int
		expected []int
	}{
		{
			name:     "Starting from 8",
			start:    8,
			expected: []int{8, 13, 21, 34, 55},
		},
		{
			name:     "Starting from 10 (not in sequence)",
			start:    10,
			expected: []int{13, 21, 34, 55, 89},
		},
		{
			name:     "Starting from 0",
			start:    0,
			expected: []int{1, 1, 2, 3, 5},
		},
		{
			name:     "Starting from 1",
			start:    1,
			expected: []int{1, 1, 2, 3, 5},
		},
		{
			name:     "Starting from a larger number",
			start:    100,
			expected: []int{144, 233, 377, 610, 987},
		},
	}

	// Iterate through test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := fibonacciSequence(tc.start)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func BenchmarkFibonacciSequence0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Running the fibonacciSequence function for a constant value
		fibonacciSequence(0)
	}
}

func BenchmarkFibonacciSequence10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Running the fibonacciSequence function for a constant value
		fibonacciSequence(10)
	}
}

func BenchmarkFibonacciSequence100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Running the fibonacciSequence function for a constant value
		fibonacciSequence(100)
	}
}
