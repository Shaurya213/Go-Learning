package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("defer will execute in reverse order")
	fmt.Println("Hello")
}
