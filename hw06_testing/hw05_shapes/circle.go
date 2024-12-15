package hw05

type Circle struct {
	radius float64
}

func (c *Circle) Radius() float64 {
	return c.radius
}

func (c *Circle) SetRadius(radius float64) {
	c.radius = radius
}

func (c *Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}
