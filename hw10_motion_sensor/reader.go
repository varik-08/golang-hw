package main

import "fmt"

type Reader struct {
	Sensor         *Sensor
	resultsChannel chan float64
	doneChannel    *chan bool
}

func NewReader(sensor *Sensor, doneChannel *chan bool) *Reader {
	resultsChannel := make(chan float64)

	return &Reader{
		resultsChannel: resultsChannel,
		Sensor:         sensor,
		doneChannel:    doneChannel,
	}
}

func (reader *Reader) Read() {
	defer close(reader.resultsChannel)

	for {
		select {
		case <-*reader.doneChannel:
			return
		default:
			sum := 0

			for i := 0; i < 10; i++ {
				value := <-reader.Sensor.DataChannel

				sum += value
			}

			average := float64(sum) / 10

			reader.resultsChannel <- average
		}
	}
}

func (reader *Reader) Print() {
	for {
		select {
		case <-*reader.doneChannel:
			return
		default:
			average := <-reader.resultsChannel
			fmt.Printf("Average: %.2f\n", average)
		}
	}
}
