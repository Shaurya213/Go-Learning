package main

import "fmt"

//const name = "golang" this is corrent
//var k = "kilo" this is also correct
//name:="game" cant declare the variable like this by := method outside the main function
func main() {
	// const name = "golang"
	// //this will give an error
	// //name = "javascript"
	// fmt.Println(name)
	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port, host)
	/*Alternnative way when we want to declare multiple constants

	 */

}
