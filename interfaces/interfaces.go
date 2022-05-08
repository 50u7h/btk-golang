package interfaces

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type rectangle struct {
	height, width float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func school(s shape) {
	fmt.Println(s.area())
}

func Demo17() {
	r := rectangle{width: 10, height: 4}
	school(r)

	c := circle{radius: 10}
	school(c)
}
