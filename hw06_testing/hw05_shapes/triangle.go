package hw05_shapes

import "math"

type Triangle struct {
	a, b, c float64
}

func (t *Triangle) A() float64 {
	return t.a
}

func (t *Triangle) SetA(a float64) {
	t.a = a
}

func (t *Triangle) B() float64 {
	return t.b
}

func (t *Triangle) SetB(b float64) {
	t.b = b
}

func (t *Triangle) C() float64 {
	return t.c
}

func (t *Triangle) SetC(c float64) {
	t.c = c
}

func (t *Triangle) Area() float64 {
	s := (t.a + t.b + t.c) / 2
	return math.Sqrt(s * (s - t.a) * (s - t.b) * (s - t.c))
}
