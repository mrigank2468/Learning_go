package main

import (
	"fmt"
	"io"
	"net/http"
)

const Url = "https://google.com"

func main() {
	fmt.Println("welcome to handling http request ")
	response, err := http.Get(Url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type: %T\n", response)
	defer response.Body.Close()
	databytes,ererr := io.ReadAll(response.Body)
	if ererr != nil {
		panic(ererr)
	}
	fmt.Println(string(databytes))
}
