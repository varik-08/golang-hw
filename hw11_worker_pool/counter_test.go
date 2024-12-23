package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCounter_Inc(t *testing.T) {
	counter := NewCounter()
	wg := sync.WaitGroup{}
	expectedCount := 1000
	wg.Add(expectedCount)

	for i := 0; i < expectedCount; i++ {
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}

	wg.Wait()

	require.Equal(t, counter.value, expectedCount)
}
