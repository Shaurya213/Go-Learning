package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions")
	result := add(3, 4)
	fmt.Println("result is", result)

	proresult, Mymessage := proAdder(23, 4, 4, 5, 5, 3)
	fmt.Println("result is", proresult, Mymessage)

}
func add(a int, b int) int {
	return a + b
}
func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}

	return total, "Hi you can give 2 values or more like this"
}
