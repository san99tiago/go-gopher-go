// Golang program to illustrate a slice
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
	
	var new_names = append(names, "Hans")
	fmt.Println("NAMES: ", new_names)

}

func fill_names(names []string)  {
	names[0] = "Santi"
	names[1] = "Moni"
	names[2] = "Elkin"
	names[3] = "Yesid"
	names[4] = "David"
}
