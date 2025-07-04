package main

import "fmt"

func sender(ch chan string) {
	ch <- "Hello from sender!"
}

func receiver(ch chan string) {
	msg := <-ch
	fmt.Println(msg)
}

func main() {
	ch := make(chan string)

	go sender(ch)
	go receiver(ch)

	var input string
	fmt.Scanln(&input)
}
