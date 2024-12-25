package main

import "sync"

func main() {
	wgDone := sync.WaitGroup{}
	wgDone.Add(1)

	sensor := NewSensor()
	reader := NewReader(sensor, &wgDone)

	go sensor.Enable()
	go reader.Read()
	go reader.Print()

	wgDone.Wait()
}
