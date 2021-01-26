package main // Do not change this.

import "reflect"

type shape interface {
	area() int
}

type square struct { //Square struct
	side int
}

type rectangle struct { //rectangle struct
	width, height int
}

func (s square) area() int { // Square area method.
	return s.side * s.side
}

func (r rectangle) area() int { // Rectangle area method.
	return r.width * r.height
}

// Your Solution Here

func myType(s shape) string {
	return reflect.TypeOf(s).String()
}
