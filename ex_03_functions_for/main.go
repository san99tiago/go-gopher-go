// Golang program to illustrate a function and simple for/forr loops
// Santiago Garcia Arango

// Including main package
package main

import (
	"fmt"
)


func main() {
	var names = make([]string, 5)
	fill_names(names)

	fmt.Println("NAMES: ", names)
	find_name_1("Yesid", names)
	find_name_2("Yesid", names)
}

func fill_names(names []string)  {
	names[0] = "Santi"
	names[1] = "Moni"
	names[2] = "Elkin"
	names[3] = "Yesid"
	names[4] = "David"
}

func find_name_1(name string, names []string)  {
	for i := 0; i < len(names); i++ {
		if names[i] == name {
			fmt.Println("Name '", name, "' was found!")
		}
	}
}

func find_name_2(name string, names []string)  {
	for index, item := range names {
		if item == name {
			fmt.Println("Name '", name, "' was found at index", index, "!")
		}
	}
}
