package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type Sensor struct {
	DataChannel chan int
	doneChannel *chan bool
}

func NewSensor(doneChannel *chan bool) *Sensor {
	dataChannel := make(chan int, 10)

	return &Sensor{
		DataChannel: dataChannel,
		doneChannel: doneChannel,
	}
}

func (sensor *Sensor) Enable() {
	defer close(sensor.DataChannel)

	for {
		select {
		case <-*sensor.doneChannel:
			return
		default:
			random, err := rand.Int(rand.Reader, big.NewInt(100))
			if err != nil {
				fmt.Println("Ошибка генерации рандомного числа: ", err)
				continue
			}

			sensor.DataChannel <- int(random.Int64())
		}
	}
}
