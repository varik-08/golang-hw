package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSensor_Enable(t *testing.T) {
	sensor := NewSensor()

	go sensor.Enable()
	value := -1

	for tempValue := range sensor.DataChannel {
		value = tempValue
		break
	}

	require.NotEqual(t, value, -1)
}
