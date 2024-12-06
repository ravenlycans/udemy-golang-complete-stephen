package main

import "fmt"

type triangle struct {
	base, height float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func (t triangle) getArea() float64 {
	return t.base * t.height * 0.5
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func getArea(s shape) float64 {
	return s.getArea()
}

func main() {
	t1 := triangle{base: 10, height: 5}
	s1 := square{sideLength: 10}

	fmt.Println(getArea(t1))
	fmt.Println(getArea(s1))
}
