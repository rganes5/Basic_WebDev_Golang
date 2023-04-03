package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://lco.dev"

func main() {
	fmt.Println("This is a web request demonstration")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type : %T", response)
	defer response.Body.Close() //callers responsibility to close the connection.

	databyte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databyte)
	fmt.Println(content)
}
