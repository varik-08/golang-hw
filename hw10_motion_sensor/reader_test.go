package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReader_Read(t *testing.T) {
	wgDone := sync.WaitGroup{}
	mockSensor := NewSensor()
	reader := NewReader(mockSensor, &wgDone)

	go reader.Read()

	for i := 0; i < 10; i++ {
		mockSensor.DataChannel <- i + 1
	}

	expectedAverage := 5.5
	average := -1.0

	for tempAverage := range reader.resultsChannel {
		average = tempAverage
		break
	}

	require.Equal(t, expectedAverage, average)
}
