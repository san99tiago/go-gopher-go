// Golang program to illustrate go routines not organized
// Santiago Garcia Arango

// Including main package
package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(3)
		go sayHello()
		go sayGoodbye()
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	fmt.Printf("Hello my friends, this is counter: #%v\n", counter)
	wg.Done()
}

func sayGoodbye() {
	fmt.Printf("Goodbye my friends, this is counter: #%v\n", counter)
	wg.Done()
}

func increment() {
	counter++
	wg.Done()
}
