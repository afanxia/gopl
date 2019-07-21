package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// Squarer
	go func() {
		for b := range naturals {
			squares <- b * b
		}
		close(squares)
	}()

	// Printer (in main goroutine)
	for c := range squares {
		fmt.Println(c)
	}
}
