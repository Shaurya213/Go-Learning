package main

import "sync"

func sum1 (ch chan int, w* sync.WaitGroup) {
	defer w.Done()
	ch <- 10 + 20	
}

func main() {
	var w sync.WaitGroup

	ch := make(chan int,2)
	
	w.Add(2)

	
}