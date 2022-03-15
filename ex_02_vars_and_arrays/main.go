// Golang program to illustrate variable usage
// Santiago Garcia Arango

// Including main package
package main

import (
	"fmt"
)


func main() {
	// VARIABLES
	var var_1 string  // Declaration without initialization
	var_1 = "Santi"
	var var_2 = "Moni" // Declaration with type inference
	var_3 := "Manuelita" // Short declaration

	// ARRAYS AND SLICES
	var arr_1 = make([]string, 3)
	arr_1[0] = "Python"
	arr_1[1] = "Java"
	arr_1[2] = "JavaScript"

	fmt.Println("*****Showing variables*****")
	fmt.Println(var_1)
	fmt.Println(var_2)
	fmt.Println(var_3)

	fmt.Println("*****Showing arrays*****")
	fmt.Println(arr_1)
}