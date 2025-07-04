package main

import "fmt"

type Speaker interface {
	Speak()
}

type Animal struct {
	Name string
}

func (a Animal) Eat() {
	fmt.Println(a.Name, "is eating.")
}

type Dog struct {
	Animal
	Breed string
}

func (d Dog) Speak() {
	fmt.Println(d.Name, "barks.")
}

func main() {
	d := Dog{
		Animal: Animal{Name: "Buddy"},
		Breed:  "Golden Retriever",
	}

	d.Eat()
	d.Speak()

	// interface usage

	var s Speaker = d
	s.Speak()
}
