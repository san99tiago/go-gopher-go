// Golang program to illustrate mutable/immutable data types and behavior
// Santiago Garcia Arango

// Including main package
package main

import "fmt"

func main() {
	fmt.Printf("\n********** EXAMPLE 1 (behavior with numbers: initial number doesn't change) **********\n")
	var number_1 int = 65
	fmt.Println("\tnumber_1 value is: ", number_1)

	fmt.Println("--> here we do: number_2 := number_1")
	number_2 := number_1
	fmt.Println("--> here we do: number_2 = 100 (without modifying number_1)")
	number_2 = 100

	fmt.Println("\tnumber_1 value is: ", number_1)
	fmt.Println("\tnumber_2 value is: ", number_2)
	fmt.Println("--> RESULT: number_1 value didn't change!")

	fmt.Printf("\n********** EXAMPLE 2 (behavior with arrays: initial array doesn't change) **********\n")
	var array_1 [3]int = [3]int{1, 2, 3}
	fmt.Println("\tarray_1 values are: ", array_1)

	fmt.Println("--> here we do: array_2 := array_1")
	array_2 := array_1
	fmt.Println("--> here we do: array_2[0] = 99 (without modifying array_1)")
	array_2[0] = 99

	fmt.Println("\tarray_1 values are: ", array_1)
	fmt.Println("\tarray_2 values are: ", array_2)
	fmt.Println("--> RESULT: array_1 values didn't change!")

	fmt.Printf("\n********** EXAMPLE 3 (interesting behavior with slices: both slices change) **********\n")
	var slice_1 []int = []int{1, 2, 3}
	fmt.Println("\tslice_1 values are: ", slice_1)

	fmt.Println("--> here we do: slice_2 := slice_1")
	slice_2 := slice_1
	fmt.Println("--> here we do: slice_2[0] = 99 (without modifying slice_1)")
	slice_2[0] = 99

	fmt.Println("\tslice_1 values are: ", slice_1)
	fmt.Println("\tslice_2 values are: ", slice_2)
	fmt.Println("--> RESULT: slice_1 values changed, because slice_2 had a pointer to slice_1")

	fmt.Printf("\n********** EXAMPLE 4 (interesting behavior with maps: both maps change) **********\n")
	var map_1 map[string]string = map[string]string{
		"user":      "san99tiago",
		"name":      "Santiago",
		"last_name": "Garcia",
		"alias":     "santi",
	}
	fmt.Println("\tmap_1 values are: ", map_1)

	fmt.Println("--> here we do: map_2 := map_1")
	map_2 := map_1
	fmt.Println("--> here we do: map_2[\"user\"]] = \"santiago.garcia\" (without modifying map_1)")
	map_2["user"] = "santiago.garcia"

	fmt.Println("\tmap_1 values are: ", map_1)
	fmt.Println("\tmap_2 values are: ", map_2)
	fmt.Println("--> RESULT: map_1 values changed, because map_2 had a pointer to map_1")
}
