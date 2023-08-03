package main

import (
	"fmt"
	"math"
)

func main() {
	var r geometry = rectangle{width: 3, height: 4}
	fmt.Println(r.area())
	fmt.Println(r.perimeter())

}

type geometry interface {
	area() float64
	perimeter() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	height, width float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius

}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}
