package main

import "fmt"

type Animal struct {
	Name string
}

func (a Animal) Speak() string {
	return "Speak"
}

type Dog struct {
	Animal
}

func main() {
	d := Dog{
		Animal: Animal{Name: "Dog"},
	}
	d.Speak()
	fmt.Println(d.Speak())
}
