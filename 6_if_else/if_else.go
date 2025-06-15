package main

import "fmt"

func main() {
	role := "admin"
	permissions := "true"
	if role == "admin" && permissions == "true" {
		fmt.Println("You have admin access")
	} else {
		fmt.Println("Sorry you dont have permissions")
	}
}
