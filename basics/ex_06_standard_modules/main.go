// Golang program to illustrate standard modules examples
// Santiago Garcia Arango

// Including main package
package main

import (
	"fmt"
	"strings"
)

func main() {
	greeting_from_santi := "Hello from Santi!"

	fmt.Println("\n********** String is: **********")
	fmt.Println(greeting_from_santi)

	does_string_contain_given_word(greeting_from_santi, "Hello")
	does_string_contain_given_word(greeting_from_santi, "Santi")
	does_string_contain_given_word(greeting_from_santi, "Moni")

	fmt.Println("\n********** Replacing \"Moni\" for \"Santi\"**********")
	greeting_from_santi = strings.ReplaceAll(greeting_from_santi, "Santi", "Moni")

	fmt.Println("\n********** String is: **********")
	fmt.Println(greeting_from_santi)

	does_string_contain_given_word(greeting_from_santi, "Moni")

}

func does_string_contain_given_word(input_string string, word_to_validate string) {
	fmt.Printf("\n********** Does string contain %s? **********\n", word_to_validate)
	fmt.Println(strings.Contains(input_string, word_to_validate))
}
