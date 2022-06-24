// Golang program to illustrate pointers usage and examples
// Santiago Garcia Arango

// Including main package
package main

import "fmt"

// EXTRA NOTES:
// & --> Means to get the pointer
// * --> Means dereference operator

func main() {
	fmt.Printf("\n********** EXAMPLE 1 **********\n")
	fmt.Println("\n--> [OPERATION] x:= 99")
	x := 99
	fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "x", x, x)
	fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "&x", &x, &x)

	fmt.Println("\n--> [OPERATION] y:= x")
	y := &x
	fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "y", y, y)
	fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "&y", &y, &y)

	fmt.Println("\n--> [OPERATION] z:= &x")
	z := &x
	fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "z", z, z)
	fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "&z", &z, &z)

}
