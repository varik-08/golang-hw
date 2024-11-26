package main

import "math"

type Triangle struct {
	a, b, c float64
}

func (t *Triangle) Area() float64 {
	s := (t.a + t.b + t.c) / 2
	return math.Sqrt(s * (s - t.a) * (s - t.b) * (s - t.c))
}
