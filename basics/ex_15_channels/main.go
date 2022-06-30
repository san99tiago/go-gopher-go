// Golang program to illustrate channels
// Santiago Garcia Arango

// Including main package
package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	fmt.Println("\n----- EX 01: SINGLE ITEM CHANNEL -----")
	ex_1_single_item_channel()
	fmt.Println("\n----- EX 02: BUFFERED CHANNEL STATIC -----")
	ex_2_buffered_channel_static()
	fmt.Println("\n----- EX 03: BUFFERED CHANNEL DYNAMIC -----")
	ex_3_buffered_channel_dynamic()

}

func ex_1_single_item_channel() {
	ch := make(chan int)
	fmt.Printf("--> Channel type is: [%T], capacity is: [%v], and value is: [%v]\n", ch, cap(ch), ch)
	wg.Add(2)

	// Receiver go routine
	go func() {
		i := <-ch // Receive data from channel
		fmt.Println(i)
		wg.Done()
	}()

	// Sender go routine
	go func() {
		i := 50
		ch <- i // Send data to channel
		i = 33  // This won't affect this program at all
		wg.Done()
	}()
	wg.Wait()

}

func ex_2_buffered_channel_static() {
	ch := make(chan int, 10)
	fmt.Printf("--> Channel type is: [%T], capacity is: [%v], and value is: [%v]\n", ch, cap(ch), ch)
	wg.Add(2)

	// Receiver go routine
	go func() {
		var item int
		for i := 0; i < cap(ch); i++ {
			item = <-ch // Receive data from channel
			fmt.Println(item)
		}
		wg.Done()
	}()

	// Sender go routine
	go func() {
		var item int
		for i := 0; i < cap(ch); i++ {
			item = i * 10
			ch <- item // Send data to channel
		}
		wg.Done()
	}()
	wg.Wait()

}

func ex_3_buffered_channel_dynamic() {
	ch := make(chan int, 100)
	fmt.Printf("--> Channel type is: [%T], capacity is: [%v], and value is: [%v]\n", ch, cap(ch), ch)
	wg.Add(2)

	// Receiver go routine
	go func() {
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}()

	// Sender go routine
	go func() {
		ch <- 1 // Send random data 0 to channel
		ch <- 9 // Send random data 1 to channel
		ch <- 0 // Send random data 2 to channel
		ch <- 5 // Send random data 3 to channel
		ch <- 6 // Send random data 4 to channel

		close(ch) // Mandatory for dynamic channel receivers (can not use channel after)

		wg.Done()
	}()
	wg.Wait()

}
