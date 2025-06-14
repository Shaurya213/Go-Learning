package main

import (
	"fmt"
)

func main() {
	a := make(map[string]string)
	var b map[string]string

	fmt.Println(a == nil)
	fmt.Println(b == nil)
}
