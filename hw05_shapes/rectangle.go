package main

type Rectangle struct {
	width, height float64
}

func (r *Rectangle) Width() float64 {
	return r.width
}

func (r *Rectangle) SetWidth(width float64) {
	r.width = width
}

func (r *Rectangle) Height() float64 {
	return r.height
}

func (r *Rectangle) SetHeight(height float64) {
	r.height = height
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}
