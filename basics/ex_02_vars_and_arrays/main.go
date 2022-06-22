// Golang program to illustrate variable usage
// Santiago Garcia Arango

// Including main package
package main

import (
	"fmt"
)

// Declare "stand-alone" global variables
var cool_constant int = 88

// Declare "grouped" global variables (useful for organizing)
var (
	name      string = "Santiago"
	last_name string = "Garcia"
	age       int    = 22
	is_cool   bool   = true
)

func main() {
	// ------------ VARIABLES ------------
	// Declaration without direct initialization...
	var var_1 string
	var_1 = "Moni"

	// Declaration with type inference...
	var var_2 = "Manuelita"

	// Declaration without type inference...
	var var_3 float32 = 99.4

	// Short declaration
	var_4 := 99.4

	// ARRAYS
	var arr_1 [4]string
	arr_1[0] = "Python"
	arr_1[1] = "Java"
	arr_1[2] = "JavaScript"
	arr_1[3] = "Go"

	var arr_2 [5]float32

	arr_3 := [8]int{0, 1, 1, 2, 3, 5, 8, 13}

	fmt.Println("\n********* Showing local variables********* ")
	fmt.Printf("%s variable value is: %v, and is of type: %T \n", "var_1", var_1, var_1)
	fmt.Printf("%s variable value is: %v, and is of type: %T \n", "var_2", var_2, var_2)
	fmt.Printf("%s variable value is: %v, and is of type: %T \n", "var_3", var_3, var_3)
	fmt.Printf("%s variable value is: %v, and is of type: %T \n", "var_4", var_4, var_4)

	fmt.Println("\n********* Showing global variables********* ")
	fmt.Printf("%s variable value is: %v, and is of type: %T \n", "cool_constant", cool_constant, cool_constant)
	fmt.Printf("%s variable value is: %v, and is of type: %T \n", "name", name, name)
	fmt.Printf("%s variable value is: %v, and is of type: %T \n", "last_name", last_name, last_name)
	fmt.Printf("%s variable value is: %v, and is of type: %T \n", "age", age, age)
	fmt.Printf("%s variable value is: %v, and is of type: %T \n", "is_cool", is_cool, is_cool)

	fmt.Println("\n********* Converting int to float32 variables ********* ")
	var my_new_cool_constant float32
	my_new_cool_constant = float32(cool_constant)
	var my_new_age float32
	my_new_age = float32(age)

	fmt.Println("\n********* Showing converted variables********* ")
	fmt.Printf("%s variable value is: %v, and is of type: %T \n", "my_new_cool_constant", my_new_cool_constant, my_new_cool_constant)
	fmt.Printf("%s variable value is: %v, and is of type: %T \n", "my_new_age", my_new_age, my_new_age)

	fmt.Println("\n********* Showing arrays********* ")
	fmt.Printf("%s variable value is: %v, is of type: %T, and has a total of %v elements \n", "arr_1", arr_1, arr_1, len(arr_1))
	fmt.Printf("%s variable value is: %v, is of type: %T, and has a total of %v elements \n", "arr_2", arr_2, arr_2, len(arr_2))
	fmt.Printf("%s variable value is: %v, is of type: %T, and has a total of %v elements \n", "arr_3", arr_3, arr_3, len(arr_3))
}
