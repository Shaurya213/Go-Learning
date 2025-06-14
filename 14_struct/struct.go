package main

import "fmt"

func main() {
	type person struct {
		name string
		age  int
	}
	p1 := person{}
	p1.name = "han"
	p1.age = 18
	fmt.Println(p1)

}
