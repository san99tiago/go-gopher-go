// Golang program to illustrate pointers usage and examples
// Santiago Garcia Arango

// Including main package
package main

import (
	"fmt"
)


func main() {
	fmt.Println("***** EXAMPLE 1 (without pointers) *****")
	var names_1 [5]string
	fill_names_1(names_1)
	fmt.Println("NAMES: ", names_1)

	fmt.Println("***** EXAMPLE 2 (with pointers) *****")
	var names_2 [5]string
	var names_2_p = &names_2
	fill_names_2(names_2_p)
	fmt.Println("NAMES: ", names_2)

}

func fill_names_1(names_array [5]string)  {
	names_array[0] = "Santi"
	names_array[1] = "Moni"
	names_array[2] = "Elkin"
	names_array[3] = "Yesid"
	names_array[4] = "David"
}

func fill_names_2(names_array *[5]string)  {
	names_array[0] = "Santi"
	names_array[1] = "Moni"
	names_array[2] = "Elkin"
	names_array[3] = "Yesid"
	names_array[4] = "David"
}