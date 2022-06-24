// Golang program to illustrate interfaces and their usage
// Santiago Garcia Arango

// Including main package
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("\n********** EXAMPLE 1 (create structs and a slice with an interface for them) **********")
	fmt.Println("\n--> [OPERATION: create 'c1' as 'circle'] c1 := circle{6}")
	fmt.Println("\n--> [OPERATION: create 's1' as 'square'] s1 := square{side: 3}")
	fmt.Println("\n--> [OPERATION: create 'r1' as 'rectangle'] r1 := rectangle{width: 2, height: 3}")
	c1 := circle{6}
	s1 := square{side: 3}
	r1 := rectangle{width: 2, height: 3}
	fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "c1", c1, c1)
	fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "s1", s1, s1)
	fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "r1", r1, r1)

	fmt.Println("\n--> [OPERATION: create 'shapes' as slice of 'geometric' interface]shapes := []geometric{c1, s1, r1}")
	shapes := []geometric{c1, s1, r1}
	fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "shapes", shapes, shapes)

	fmt.Println("\n--> [OPERATION: show each shape and its area from the 'shapes' struct]")
	for _, shape := range shapes {
		fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "shape", shape, shape)
		fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "shape.get_area()", shape.get_area(), shape.get_area())
	}

}

// INTERFACE THAT IMPLIES THAT "GEOMETRIC" MUST HAVE "GET_AREA()" METHOD ON THEM
type geometric interface {
	get_area() float64
}

// CIRCLE STRUCT AND METHODS
type circle struct {
	radius float64
}

func (c circle) get_area() float64 {
	return math.Pi * c.radius * c.radius
}

// SQUARE STRUCT AND METHODS
type square struct {
	side float64
}

func (s square) get_area() float64 {
	return s.side * s.side
}

// RECTANGLE STRUCT AND METHODS
type rectangle struct {
	width  float64
	height float64
}

func (r rectangle) get_area() float64 {
	return r.height * r.width
}
