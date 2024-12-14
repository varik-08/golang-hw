package hw03

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateBoard(t *testing.T) {
	tests := []struct {
		name string
		size int
		want [][]rune
	}{
		{
			name: "size 1",
			size: 1,
			want: [][]rune{
				{' '},
			},
		},
		{
			name: "size 2",
			size: 2,
			want: [][]rune{
				{' ', '#'},
				{'#', ' '},
			},
		},
		{
			name: "size 8",
			size: 8,
			want: [][]rune{
				{' ', '#', ' ', '#', ' ', '#', ' ', '#'},
				{'#', ' ', '#', ' ', '#', ' ', '#', ' '},
				{' ', '#', ' ', '#', ' ', '#', ' ', '#'},
				{'#', ' ', '#', ' ', '#', ' ', '#', ' '},
				{' ', '#', ' ', '#', ' ', '#', ' ', '#'},
				{'#', ' ', '#', ' ', '#', ' ', '#', ' '},
				{' ', '#', ' ', '#', ' ', '#', ' ', '#'},
				{'#', ' ', '#', ' ', '#', ' ', '#', ' '},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateBoard(tt.size)

			require.Equal(t, got, tt.want)
		})
	}
}
