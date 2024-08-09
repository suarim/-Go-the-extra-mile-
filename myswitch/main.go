package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello, world!")
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(6) + 1
	switch 3 {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
		fallthrough
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	case 6:
		fmt.Println("Six")
	default:
		fmt.Println("Unknown number")
	}
	fmt.Println(randNum)
}
