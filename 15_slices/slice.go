package main

import "fmt"

//slice -> dynamic array
//most used construct in go
// + useful methods
func main() {
	//uninitialized slice is nil
	// var nums []int
	// fmt.Println(len(nums))

	var nums = make([]int, 2, 5)
	fmt.Println(nums)
	for i := len(nums) + 1; i <= 5; i++ {
		nums = append(nums, i)
	}
	fmt.Println(nums)
}
