package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
	performPostRequest()
}

// func performGetRequest() {

// }

func performPostRequest() {
	url := "http://localhost:8000/post"
	fmt.Println("URL:>", url)
	requestBody := strings.NewReader(`
	{
		"userId": 1,
		"id": 1,
		"title": "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
		"body": "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto"
	}
	`)
	response, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println(string(content))
}
