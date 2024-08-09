package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world!")
	suarim := User{
		Name:   "Suarim",
		Age:    21,
		Email:  "kjk",
		Status: true,
	}
	fmt.Println(suarim)
	fmt.Println(suarim.Name)
}

type User struct {
	Name   string
	Age    int
	Email  string
	Status bool
}
