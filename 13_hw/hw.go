package main

import "fmt"

type Person struct {
	name     string
	surname  string
	phone    int
	email    string
	age      int
	zip      int
	relation string
}

func main() {
	var per1 Person
	per1.name = "shaurya"
	per1.surname = "Choudha"
	per1.phone = 7777777777
	per1.email = "shag@gmail.com"
	per1.age = 21
	per1.zip = 789878
	per1.relation = "student"

	fmt.Println("Name : ", per1.name)
	fmt.Println("Surname : ", per1.surname)
	fmt.Println("age : ", per1.age)
	fmt.Println("Email : ", per1.email)
	fmt.Println("Zip : ", per1.zip)
	fmt.Println("Relation : ", per1.relation)

	var per2 Person
	per2.name = "Harsh"
	per2.surname = "shit"
	per2.phone = 7777000077
	per2.email = "harii@gmail.com"
	per2.age = 29
	per2.zip = 780000
	per2.relation = "kid"

	fmt.Println("Name : ", per2.name)
	fmt.Println("Surname : ", per2.surname)
	fmt.Println("age : ", per2.age)
	fmt.Println("Email : ", per2.email)
	fmt.Println("Zip : ", per2.zip)
	fmt.Println("Relation : ", per2.relation)

	var addressBook []Person
	addressBook = append(addressBook, per1, per2)
	fmt.Println("AddressBook : ", addressBook)

}
