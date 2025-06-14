package main

import "fmt"

func main() {

	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("Its integer")
		default:
			fmt.Println("Its not an interger", t)
		}
	}
	whoAmI("golang")
}
