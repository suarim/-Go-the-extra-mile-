package main

import (
	"fmt"
	"time"
)

func main() {
	// Your code goes here
	fmt.Println("My Maths in Go")

	presenttime := time.Now()
	fmt.Println("Present time is: ", presenttime)

	fmt.Println(presenttime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2024, time.August, 6, 8, 9, 0, 0, time.UTC)
	fmt.Println("Created date is: ", createdDate.Format("Monday"))

}
