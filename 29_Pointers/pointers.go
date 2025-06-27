package main

import "fmt"

func Nonptr(i int) {
	i = 0
}

func Ptr(i *int) {
	*i = 0
}
func main() {
	var i int = 1
	Nonptr(i)
	fmt.Println("before nonptr", i)
	Ptr(&i)
	fmt.Println("after ptr", i)
}
