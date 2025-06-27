package main

import "fmt"

func main() {
	p := new(int)   // allocates int, returns pointer
	fmt.Println(*p) // prints 0 (zero value)

	*p = 42
	fmt.Println(*p) // prints 42
}
