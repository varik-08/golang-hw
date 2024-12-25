package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"
	"time"
)

type Sensor struct {
	DataChannel chan int
}

func NewSensor() *Sensor {
	dataChannel := make(chan int, 10)

	return &Sensor{
		DataChannel: dataChannel,
	}
}

func (sensor *Sensor) Enable() {
	defer close(sensor.DataChannel)
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		timer := time.NewTimer(5 * time.Second)
		defer timer.Stop()

		for {
			select {
			case <-timer.C:
				wg.Done()

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
	}()

	wg.Wait()
}
