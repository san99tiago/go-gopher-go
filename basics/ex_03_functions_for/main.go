// Golang program to illustrate a function and simple for/forr loops
// Santiago Garcia Arango

// Including main package
package main

import (
	"fmt"
)

func main() {
	fmt.Println("\n********** Starting processing of names **********")

	var names []string = []string{"Santi", "Moni", "Elkin", "Yesid", "David", "Esteban"}
	fmt.Println("NAMES: ", names)

	// Please look at the 2 different "for" approaches on each "find_name_x" function
	find_name_1("Yesid", names)
	find_name_2("Yesid", names)

	fmt.Println("\n********** Starting processing of numbers **********")
	var numbers []int = []int{1, 5, 3, 6, 7, 3, 4, 6}
	fmt.Println("NUMBERS: ", numbers)

	// This functions illustrates double "for" solutions with elements and indexes used
	find_duplicates_1(numbers)
	find_duplicates_2(numbers)
}

func find_name_1(name string, names []string) {
	fmt.Println("\n----- Finding only the name ----")
	for i := 0; i < len(names); i++ {
		if names[i] == name {
			fmt.Printf("Name: %q , was found!\n", name)
			break
		}
	}
}

func find_name_2(name string, names []string) {
	fmt.Println("\n----- Finding name and its index ----")
	for index, item := range names {
		if item == name {
			fmt.Printf("Name: %q , was found at index: %d\n", name, index)
			break
		}
	}
}

func find_duplicates_1(numbers []int) {
	fmt.Println("\n----- Finding duplicates of array ----")
	var duplicates []int
	for i, element_1 := range numbers {
		for j, element_2 := range numbers {
			if element_1 == element_2 && i < j {
				duplicates = append(duplicates, element_2)
			}
		}
	}
	fmt.Println("Duplicates are: ", duplicates)
}

func find_duplicates_2(numbers []int) {
	fmt.Println("\n----- Finding duplicates of array ----")
	var duplicates []int
	for i, element_1 := range numbers {
		for j := i + 1; j < len(numbers); j++ {
			element_2 := numbers[j]
			if element_1 == element_2 {
				duplicates = append(duplicates, element_2)
			}
		}
	}
	fmt.Println("Duplicates are: ", duplicates)
}
