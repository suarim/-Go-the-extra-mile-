package main

import (
	"fmt"
	"net/url"
)

var myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	fmt.Println("Hello, world!")
	fmt.Println(myurl)
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())
	fmt.Println(result.Query())
	x := result.Query()
	fmt.Println(x["coursename"][0])
	for _, val := range x {
		fmt.Println(val[0])
	}
	partsofurls := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tut",
		RawPath: "user=abc",
	}
	newurl := partsofurls.String()
	fmt.Println(newurl)
}
