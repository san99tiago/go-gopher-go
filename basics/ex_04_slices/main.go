// Golang program to illustrate a slice
// Santiago Garcia Arango

// Including main package
package main

import (
	"fmt"
)

func main() {
	// Create array...
	var all_names [10]string = [10]string{
		"Santiago",
		"Monica",
		"Elkin",
		"Yesid",
		"Medina",
		"Upegui",
		"Manu",
		"Hans",
		"Manuelita",
		"Juanjo"}
	fmt.Printf("All names are: %q, ARRAY capacity is: %d, and type is: %T\n", all_names, cap(all_names), all_names)

	// Create slices from array...
	var names_couple []string = all_names[0:2]
	var names_au_eu []string = all_names[2:6]
	var names_6_fantastic []string = append(all_names[0:2], all_names[6:]...)

	fmt.Printf("Couple's names are: %q, SLICE capacity is: %d, and type is: %T\n", names_couple, cap(names_couple), names_couple)
	fmt.Printf("AU-EU's names are: %q, SLICE capacity is: %d, and type is: %T\n", names_au_eu, cap(names_au_eu), names_au_eu)
	fmt.Printf("6-Fantastic's names are: %q, SLICE capacity is: %d, and type is: %T\n", names_6_fantastic, cap(names_6_fantastic), names_6_fantastic)

	// Create slice from scratch way 1 (without capacity)...
	var names_family_no_capacity []string = []string{"Dad", "Mom", "Sister"}
	fmt.Printf("Family's no-cap names are: %q, SLICE capacity is: %d, and type is: %T\n", names_family_no_capacity, cap(names_family_no_capacity), names_family_no_capacity)

	// Create slice from scratch way 2 (with capacity)...
	var names_family = make([]string, 0)
	fmt.Printf("Family's with-cap v1 names are: %q, SLICE capacity is: %d, and type is: %T\n", names_family, cap(names_family), names_family)
	names_family = append(names_family, "Dad")
	fmt.Printf("Family's with-cap v2 names are: %q, SLICE capacity is: %d, and type is: %T\n", names_family, cap(names_family), names_family)
	names_family = append(names_family, []string{"Mom", "Sister"}...)
	fmt.Printf("Family's with-cap v3 names are: %q, SLICE capacity is: %d, and type is: %T\n", names_family, cap(names_family), names_family)

}
