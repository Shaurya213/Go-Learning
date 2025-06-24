package main

import "fmt"

func main() {
	Shaurya := User{Name: " Shaurya", IsActive: true}
	Shaurya.GetStatus()
}

type User struct {
	Name     string
	IsActive bool
}

func (u User) GetStatus() {
	fmt.Println(" Is user active: ", u.IsActive)
}
