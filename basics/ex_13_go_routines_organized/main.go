// Golang program to illustrate go routines organized (with RWMutex)
// Santiago Garcia Arango

// Including main package
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{} // wait for groups of go-routines to complete
var counter = 0
var m = sync.RWMutex{} // protect data access

func main() {
	runtime.GOMAXPROCS(100) // Tune the max threads (warning, too many could slow down)

	for i := 0; i < 10; i++ {
		wg.Add(3)

		m.RLock()
		go sayHello()
		m.RLock()
		go sayGoodbye()

		m.Lock()
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	fmt.Printf("Hello my friends, this is counter: #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func sayGoodbye() {
	fmt.Printf("Goodbye my friends, this is counter: #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}
