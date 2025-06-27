package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Welcome to web werb methods")
	PerformGetRequest()

}

func PerformGetRequest() {
	const url = "http://localhost:8080/get"

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status Code : ", response.StatusCode)
	fmt.Println("Content length  : ", response.ContentLength)

	content, _ := io.ReadAll(response.Body)
	fmt.Println("Response Body : ", string(content))

}
