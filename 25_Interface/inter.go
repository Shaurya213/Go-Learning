package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
	Dog
}

func main() {
	dog := Cat{}
	fmt.Println(dog.Dog.Speak())
	fmt.Println(dog.Speak())
}
