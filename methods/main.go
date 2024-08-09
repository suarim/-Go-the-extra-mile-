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
	suarim.GetStatus()
	suarim.newEmail("kbb@jgjg.com")
	fmt.Print(suarim.Email)
}

type User struct {
	Name   string
	Age    int
	Email  string
	Status bool
}

func (u User) GetStatus() {
	fmt.Print(u.Status)
}

func (u *User) newEmail(val string) {
	u.Email = val
	fmt.Println(u.Email)
}
