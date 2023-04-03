package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghj231ghnj"

func main() {
	fmt.Println("welcome to handling urls in golang")
	fmt.Println(myurl)

	//parse
	result, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Scheme)
	fmt.Println(result.Path)
	fmt.Println(result.Host)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The qparams is of type : %T", qparams)
	fmt.Println("")
	fmt.Println(qparams["coursename"])
	fmt.Println(qparams["paymentid"])

	for _, v := range qparams {
		fmt.Println("Param is: ", v)
	}

	//construct a url
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)

}
