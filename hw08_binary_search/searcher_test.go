package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name   string
		slice  []int
		target int
		want   int
	}{
		{
			name:   "array 1",
			slice:  []int{1, 3, 5, 7, 9},
			target: 5,
			want:   2,
		},
		{
			name:   "array 2",
			slice:  []int{1, 3, 5, 7, 9, 10, 200, 300},
			target: 200,
			want:   6,
		},
		{
			name:   "target not found",
			slice:  []int{1, 3, 5, 7, 9},
			target: 4,
			want:   -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BinarySearch(tt.slice, tt.target)

			require.Equal(t, got, tt.want)
		})
	}
}
