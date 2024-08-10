package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Printf("%T\n", res)
	// fmt.Println(res)
	defer res.Body.Close()
	dbts, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(dbts))
}

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, World!")
// }
