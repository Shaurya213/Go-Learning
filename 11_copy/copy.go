package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))

	neededNum := numbers[:len(numbers)-1]
	//Using make to get the same array size as neededNum
	//in make we can specify the length and capacity and if capacity in not defined then it will be same as lenght
	numberscopy := make([]int, len(neededNum))
	copy(numberscopy, neededNum)

	fmt.Printf("numbersCopy = %v\n", numberscopy)
	fmt.Printf("length = %d\n", len(numberscopy))
	fmt.Printf("capacity = %d\n", cap(numberscopy))

}
