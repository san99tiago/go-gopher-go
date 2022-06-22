// Golang program to illustrate a console inputs from user and a cool output
// Santiago Garcia Arango

// Including main package
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	// Input gathering from the user
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Please enter the year that you were born:\n")
	scanner.Scan()
	input_string := scanner.Text()

	// Convert input sting number to int64 type (to be able to do math operations)
	var input_int, _ = strconv.ParseInt(input_string, 10, 64)

	// Find approximately age of the user based on current year
	t := time.Now()
	var current_year int64 = int64(t.Year())
	var age int64 = current_year - input_int

	fmt.Println("Hey my friend, nice to meet you!")
	if age < 0 || age > 120 {
		fmt.Println("I am pretty sure that the year you entered is not the real one!")
	} else {
		fmt.Printf("I am guessing that you are close to %d years old :)\n", age)
	}
}
