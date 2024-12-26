package main

import (
	"fmt"
	"sync"
)

type Reader struct {
	Sensor         *Sensor
	resultsChannel chan float64
	wgDone         *sync.WaitGroup
}

func NewReader(sensor *Sensor, wgDone *sync.WaitGroup) *Reader {
	resultsChannel := make(chan float64)

	return &Reader{
		resultsChannel: resultsChannel,
		Sensor:         sensor,
		wgDone:         wgDone,
	}
}

func (reader *Reader) Read() {
	defer close(reader.resultsChannel)

	index := 0
	sum := 0

	for value := range reader.Sensor.DataChannel {
		sum += value

		if index == 9 {
			average := float64(sum) / 10
			reader.resultsChannel <- average

			index = 0
			sum = 0
		} else {
			index++
		}
	}
}

func (reader *Reader) Print() {
	for average := range reader.resultsChannel {
		fmt.Printf("Average: %.2f\n", average)
	}

	reader.wgDone.Done()
}
