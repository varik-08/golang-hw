package hw05

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculateArea(t *testing.T) {
	tests := []struct {
		name         string
		shape        Shape
		expectedArea float64
	}{
		{
			name:         "circle",
			shape:        &Circle{radius: 5},
			expectedArea: 78.5,
		},
		{
			name:         "rectangle",
			shape:        &Rectangle{width: 10, height: 8},
			expectedArea: 80,
		},
		{
			name:         "triangle",
			shape:        &Triangle{a: 3, b: 4, c: 5},
			expectedArea: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			area, err := CalculateArea(tt.shape)

			require.Nil(t, err)

			require.Equal(t, tt.expectedArea, area)
		})
	}
}
