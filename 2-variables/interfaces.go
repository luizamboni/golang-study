package main

import "fmt"

type Point struct {
	X int
	Y int
}

type Geometry interface {
	Position() Point
}

type Round interface {
	getRadius() int
}

type Circle struct {
	position Point
	Radius   int
}

// Circle implements methods of Geometry interface
// then it is a Geometry "type"
func (s Circle) Position() Point {
	return s.position
}

// Circle implements too getRadius() of
// Round interface
// then, it is a Round "type" too
func (s Circle) getRadius() int {
	return s.Radius
}

// this function require a object that implement a Geometry interface
func printPosition(shape Geometry) {
	fmt.Println("Position is", shape.Position())
}

// this function require a object that implement a Round interface
func printRadius(round Round) {
	fmt.Println("Radius is", round.getRadius())
}

func main() {
	s1 := Circle{
		Point{5, 5},
		10,
	}

	printPosition(s1)

	printRadius(s1)
}
