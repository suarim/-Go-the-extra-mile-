package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world!")
	languages := make(map[string]string)
	languages["en"] = "English"
	languages["fr"] = "French"
	languages["es"] = "Spanish"
	fmt.Println("languages = ", languages)
	fmt.Println(languages["en"])
	delete(languages, "fr")
	fmt.Println("languages = ", languages)

	for key, value := range languages {
		fmt.Printf("key = %s, value = %s\n", key, value)
	}
}
