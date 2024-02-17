package main

import "fmt"

// One interface that not implemente any methods is implemented
// by all structs
// because all implements all yours 0 methods
type Any interface {
}

// for example
// this implements Any interface
type Point struct {
	X int
	Y int
}

type Vector struct {
	A Point
	B Point
}

func acceptAnyType(obj Any) {
	fmt.Println("obj:", obj)

	// type switch to handle multiple
	switch v := obj.(type) {
	case Vector:
		fmt.Println("is a Vector", v.A, v.B)

	case Point:
		fmt.Println("is a Point", v.X, v.Y)
	}

}

func main() {
	pt := Point{10, 5}
	acceptAnyType(pt)

	vt := Vector{
		Point{1, 5},
		Point{5, 6},
	}

	acceptAnyType(vt)
}
