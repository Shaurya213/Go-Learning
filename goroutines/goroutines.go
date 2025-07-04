package main

import (
	"fmt"
	"sync"
)

func addNumbers(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 10 + 20
	ch <- sum
}

func addMoreNumbers(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 5 + 15
	ch <- sum
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 2) // buffered channel

	wg.Add(2)

	go addNumbers(ch, &wg)
	go addMoreNumbers(ch, &wg)

	wg.Wait()

	// Read results from channel
	total1 := <-ch
	total2 := <-ch

	fmt.Println("Sum 1:", total1)
	fmt.Println("Sum 2:", total2)
	fmt.Println("Total:", total1+total2)
}
