package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://golang.org"

func main() {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type %T\n", response)

	defer response.Body.Close() //Caller responibility to close the the response body

	databytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("The content of the page is ->", string(databytes))

}
