package main

import "fmt"

func area() float64 {
	//fmt.Println("")
	return 0.0
}

type Shape interface {
	area() float64
	perimeter() float64
}

type Rectangle struct {
	length, breadth float64
}

func (r Rectangle) area() float64 {
	return r.length * r.breadth
}

func (r Rectangle) perimeter() float64 {
	return 2 * (r.length + r.breadth)
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * 3.14 * c.radius
}

func TotalArea(s ...Shape) float64 {
	totalArea := 0.0
	for _, sh := range s {
		totalArea += sh.area()
	}
	return totalArea
}

func main() {
	r := Rectangle{10.0, 20.0}
	c := Circle{25.0}
	ta := TotalArea(r, c)
	fmt.Println(ta)
	area()
}
