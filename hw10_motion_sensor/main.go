package main

import "time"

func main() {
	doneChannel := make(chan bool)
	go func() {
		defer close(doneChannel)

		time.Sleep(1 * time.Minute)

		doneChannel <- true
	}()
	sensor := NewSensor(&doneChannel)
	reader := NewReader(sensor, &doneChannel)

	go sensor.Enable()
	go reader.Read()
	go reader.Print()

	<-doneChannel
}
