// Golang program to illustrate a simple loop with a cyclic message
// Santiago Garcia Arango

// Including main package
package main

import (
	"fmt"
	"time"
)

var (
	seconds_to_wait int = 1
)

func main() {
	for {
		fmt.Println("Hello my amazing friends!!!")
		time.Sleep(time.Second * time.Duration(seconds_to_wait))
	}
}