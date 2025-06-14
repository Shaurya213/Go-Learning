package main

import "fmt"

//for-> only for loop is available in golan so we implemment the while loop using for loop
func main() {
	/*while loop
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i = i + 1
	}*/

	/* infinite loop
	for{
		Println("1")
	}
	*/

	for i := 0; i <= 3; i++ {
		if i == 2 {
			continue
		}
		fmt.Println(i)
	}
	//for printing the number ranging from 0 to specific range
	for i := range 3 {
		if i == 2 {
			continue
		}
		fmt.Println(i)
	}
}
