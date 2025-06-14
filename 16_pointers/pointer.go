package main

import "fmt"

func changeval(x *int) {
	*x = 20
}
func sameval(x int) int {
	x = 30
	return x
}
func main() {
	x := 0
	fmt.Println("before change", x)
	sameval(x)
	fmt.Println("without pointer change:", x)
	changeval(&x)
	fmt.Println("After pointer change", &x)
}
