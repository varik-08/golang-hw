package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestReader_Read(t *testing.T) {
	doneChannel := make(chan bool)
	go func() {
		defer close(doneChannel)

		time.Sleep(1 * time.Second)

		doneChannel <- true
	}()

	mockSensor := NewSensor(&doneChannel)
	reader := NewReader(mockSensor, &doneChannel)

	go reader.Read()

	for i := 0; i < 10; i++ {
		mockSensor.DataChannel <- i + 1
	}

	for {
		select {
		case average := <-reader.resultsChannel:
			require.Equal(t, average, 5.5)
		case <-doneChannel:
			return
		}
	}
}
