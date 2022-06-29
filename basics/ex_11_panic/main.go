// Golang program to illustrate panic usage
// Santiago Garcia Arango

// Including main package
package main

import (
	"fmt"
	"log"
)

func main() {
	// Comment and uncomment these lines to show panic examples

	// EXAMPLE 1: show built-in panic output (throw error)
	// illustrate_panic_example()

	// EXAMPLE 2: show custom panic output when a condition is given
	// create_panic_based_on_input(0)
	// create_panic_based_on_input(1)

	// EXAMPLE 3: survive panic and keep a program running with "recover()"
	survive_panic()
	fmt.Println("We successfully survived panic with the help of recover!")

}

func illustrate_panic_example() {
	a, b := 1, 0
	ans := a / b
	fmt.Println("Answer is: ", ans)
}

func create_panic_based_on_input(number int) {
	if number == 0 {
		panic("This is bad, you can NOT use 0 as the number")
	} else {
		fmt.Println("Your cool number is ok")
	}
}

func survive_panic() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()
	fmt.Println(divide(10, 0))
}

func divide(a, b int) int {
	return a / b
}
