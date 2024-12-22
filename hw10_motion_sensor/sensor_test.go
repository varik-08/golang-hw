package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestSensor_Enable(t *testing.T) {
	doneChannel := make(chan bool)
	go func() {
		defer close(doneChannel)

		time.Sleep(1 * time.Second)

		doneChannel <- true
	}()

	sensor := NewSensor(&doneChannel)

	go sensor.Enable()
	value := -1

	for {
		select {
		case value = <-sensor.DataChannel:
		case <-doneChannel:
			require.NotEqual(t, value, -1)
			return
		}
	}
}
