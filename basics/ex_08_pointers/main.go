// Golang program to illustrate pointers usage and examples
// Santiago Garcia Arango

// Including main package
package main

import "fmt"

// EXTRA NOTES:
// & --> Means pointer operator
// * --> Means dereference operator

func main() {
	fmt.Printf("\n********** EXAMPLE 1 **********\n")
	fmt.Println("\n--> [OPERATION 'assign x'] x := 99")
	x := 99
	fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "x", x, x)
	fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "&x", &x, &x)
	fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "*&x", *&x, *&x)

	fmt.Println("\n--> [OPERATION 'assign y equals to x'] y := x")
	y := x
	fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "y", y, y)
	fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "&y", &y, &y)
	fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "*&y", *&y, *&y)

	fmt.Println("\n--> [OPERATION 'assign z equals to pointer of x'] z := &x")
	z := &x
	fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "z", z, z)
	fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "&z", &z, &z)
	fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "*&z", *&z, *&z)

	fmt.Println("\n--> [OPERATION 'modify z value to modify x (dereference)'] *z = 11")
	*z = 11
	fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "x", x, x)
	fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "&x", &x, &x)
	fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "*&x", *&x, *&x)
	fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "z", z, z)
	fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "&z", &z, &z)
	fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "*&z", *&z, *&z)

	fmt.Printf("\n********** EXAMPLE 2 **********\n")
	fmt.Println("\n--> [OPERATION] initial_string_1 := \"Well done is better than well said\"")
	fmt.Println("--> [OPERATION] initial_string_2 := \"Well done is better than well said\"")
	initial_string_1 := "Well done is better than well said"
	initial_string_2 := "Well done is better than well said"

	fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "initial_string_1", initial_string_1, initial_string_1)
	fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "initial_string_2", initial_string_2, initial_string_2)

	fmt.Println("\n--> [OPERATION] change_value_1(&initial_string_1)")
	fmt.Println("--> [OPERATION] change_value_2(initial_string_2)")
	change_value_1(&initial_string_1)
	change_value_2(initial_string_2)

	fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "initial_string_1", initial_string_1, initial_string_1)
	fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "initial_string_2", initial_string_2, initial_string_2)
}

func change_value_1(cool_string *string) {
	*cool_string = "Discipline, wil sooner or later, defeat intelligence"
}

func change_value_2(cool_string string) {
	cool_string = "Discipline, wil sooner or later, defeat intelligence"
}
