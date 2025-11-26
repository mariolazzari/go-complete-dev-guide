package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	b float64
	h float64
}

type square struct {
	l float64
}

func main() {
	s := square{2}
	t := triangle{2, 3}

	printArea(&s)
	printArea(&t)
}

func printArea(s shape) {
	fmt.Println("Area:", s.getArea())
}

func (t *triangle) getArea() float64 {
	return t.b * t.h / 2
}

func (s *square) getArea() float64 {
	return s.l * s.l
}
