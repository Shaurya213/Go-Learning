package main

import "fmt"

/*func main() {
	// assigning multiple values
	// var a, b, c, d int = 1, 2, 3, 4
	// fmt.Println(a, b, c, d)
	var a, b = 6, "Hello"
	c, d := 7, "World!"
	fmt.Println(a, b, c, d)
}*/

/*func main() {
	var (
		a int
		b int    = 1
		c string = "Hello"
	)
	fmt.Println(a, b, c)
}*/

func main() {
	var i string = "Hello"
	var j int = 15

	fmt.Printf("i has value: %v and type: %T\n", i, i)
	fmt.Printf("j has value: %v and type: %T", j, j)
}
