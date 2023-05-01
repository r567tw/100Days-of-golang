package main

import (
	"fmt"
	"math"
	"reflect"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Length float64
}

func (c *Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func (r *Rectangle) Area() float64 {
	return r.Length * r.Width
}

func getArea(shape Shape) float64 {
	return shape.Area()
}

func main() {
	c := Circle{
		Radius: 3.0,
	}

	r := Rectangle{
		Width:  1.0,
		Length: 2.0,
	}

	fmt.Printf("Circle type:  %s, Rectangle type: %s\n", reflect.TypeOf(c), reflect.TypeOf(r))
	fmt.Printf("Circle area:  %.2f, Rectangle area: %.2f\n", getArea(&c), getArea(&r))

}
