package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountWords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]int
	}{
		{
			input:    "Hello, World! Hello, World!",
			expected: map[string]int{"Hello": 2, "World": 2},
		},
		{
			input:    "This is a test.",
			expected: map[string]int{"This": 1, "is": 1, "a": 1, "test": 1},
		},
		{
			input:    "This is a test, with some punctuation.",
			expected: map[string]int{"This": 1, "is": 1, "a": 1, "test": 1, "with": 1, "some": 1, "punctuation": 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := CountWords(tt.input)
			require.Equal(t, tt.expected, result)
		})
	}
}
