package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckLogLevel(t *testing.T) {
	tests := []struct {
		level   string
		want    bool
		message string
	}{
		{level: "error", want: true},
		{level: "warning", want: true},
		{level: "info", want: true},
		{level: "debug", want: true},
		{level: "invalid", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.level, func(t *testing.T) {
			result := CheckLogLevel(tt.level)

			require.Equal(t, result, tt.want)
		})
	}
}

func TestAnalyseLog(t *testing.T) {
	tests := []struct {
		line            string
		level           string
		expectedDate    string
		expectedMessage string
	}{
		{
			line:            "2025-01-01 12:00:00 - error - some error\n",
			level:           "error",
			expectedDate:    "2025-01-01 12:00:00",
			expectedMessage: "some error",
		},
		{
			line:            "2025-01-01 12:00:00 - warning - some warning",
			level:           "warning",
			expectedDate:    "2025-01-01 12:00:00",
			expectedMessage: "some warning",
		},
		{
			line:            "2025-01-01 12:00:00 - info - some info",
			level:           "info",
			expectedDate:    "2025-01-01 12:00:00",
			expectedMessage: "some info",
		},
		{
			line:            "2025-01-01 12:00:00 - info - some info",
			level:           "warning",
			expectedDate:    "",
			expectedMessage: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.level, func(t *testing.T) {
			date, message := AnalyseLog(tt.line, tt.level)

			require.Equal(t, date, tt.expectedDate)
			require.Equal(t, message, tt.expectedMessage)
		})
	}
}
