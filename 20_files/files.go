package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Println("Welcome to the file handling in Go")
	content := "This needs to go in a file - LearnCodeOnline.in"

	file, err := os.Create("./mylcogofile.txt")

	if err != nil {
		panic(err)
	}
	//io package for write
	lenght, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}
	fmt.Println("lenght is ->", lenght)

	defer file.Close()

	readFiles("./mylcogofile.txt")

}

func readFiles(filename string) {
	databyte, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("Text inside the file is -> ", string(databyte))
}
