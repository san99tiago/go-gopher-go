// Golang program to illustrate Pointers usage
// Santiago Garcia Arango

// PENDIENTE POR TERMINARRRR

// Including main package
package main

import (
	"fmt"
)

func main() {
	fmt.Println("\n---------- TEST 01 ----------")
	test_01()

	fmt.Println("\n---------- TEST 02 ----------")
	test_02()
}

func test_01() {
	var name = "Santiago"
	var point_to_name = &name

	fmt.Println("\n***** BEFORE CHANGE *****")
	fmt.Println("--> name is: ", name)
	fmt.Println("--> *point_to_name is: ", *point_to_name)
	fmt.Println("--> point_to_name is: ", point_to_name)
	fmt.Println("--> &point_to_name is: ", &point_to_name)

	fmt.Println("\n***** APPLYING CHANGE *****")
	fmt.Println("--> Change is: [name = \"Santi\"]")
	name = "Santi"

	fmt.Println("\n***** AFTER CHANGE *****")
	fmt.Println("--> name is: ", name)
	fmt.Println("--> *point_to_name is: ", *point_to_name)
	fmt.Println("--> point_to_name is: ", point_to_name)
	fmt.Println("--> &point_to_name is: ", &point_to_name)
}

func test_02() {
	var name = "Santiago"
	var point_to_name = &name

	fmt.Println("\n***** BEFORE CHANGE *****")
	fmt.Println("--> name is: ", name)
	fmt.Println("--> *point_to_name is: ", *point_to_name)
	fmt.Println("--> point_to_name is: ", point_to_name)
	fmt.Println("--> &point_to_name is: ", &point_to_name)

	fmt.Println("\n***** APPLYING CHANGE *****")
	fmt.Println("--> Change is: [*point_to_name = \"Santi\"]")
	*point_to_name = "Santi"

	fmt.Println("\n***** AFTER CHANGE *****")
	fmt.Println("--> name is: ", name)
	fmt.Println("--> *point_to_name is: ", *point_to_name)
	fmt.Println("--> point_to_name is: ", point_to_name)
	fmt.Println("--> &point_to_name is: ", &point_to_name)

}
