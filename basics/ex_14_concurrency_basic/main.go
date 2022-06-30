// Golang program to illustrate concurrency basics
// Santiago Garcia Arango

// Including main package
package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1000 * time.Millisecond)
	boom := time.After(10000 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println(".")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
