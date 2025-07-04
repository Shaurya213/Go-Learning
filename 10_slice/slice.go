package main

import "fmt"

func main() {
	numbers := make([]int, 2, 5)
	numbers = append(numbers, 1)
	numbers = append(numbers, 2)
	numbers = append(numbers, 3)
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))
	slice1 := numbers[2:4]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
}
