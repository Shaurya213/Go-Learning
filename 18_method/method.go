package main

import "fmt"

type User struct {
	Name     string
	IsActive bool
}

func main() {
	Shaurya := User{Name: " Shaurya", IsActive: true}
	Shaurya.GetStatus()
}

func (u User) GetStatus() {
	fmt.Println(" Is user active: ", u.IsActive)
}
