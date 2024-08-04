package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello, world!")

	// Create a slice of strings named fruits
	var fruits = []string{"apple", "banana", "cherry"}

	// Print the type of fruits
	fmt.Printf("type of fruits is %T", fruits)

	// Print the contents of fruits
	fmt.Println("\nfruits = ", fruits)

	// Append "kiwi" to fruits
	fruits = append(fruits, "kiwi")

	// Print the updated fruits
	fmt.Println("fruits = ", fruits)

	// Remove the first element from fruits
	fruits = fruits[1:]

	// Create a new slice x that contains the first 3 elements of fruits
	x := fruits[:3]

	// Print the contents of x
	fmt.Println("x = ", x)

	// Print the updated fruits
	fmt.Println("\nfruits = ", fruits)

	highscore := make([]int, 4)
	highscore[0] = 234
	highscore[1] = 543
	highscore[2] = 123
	highscore[3] = 999
	highscore = append(highscore, 1000, 1500, 2000)
	fmt.Println("highscore = ", highscore)
	fmt.Println("length of highscore = ", len(highscore))

	sort.Ints(highscore)
	fmt.Println("sorted highscore = ", highscore)
	fmt.Println(sort.IntsAreSorted(highscore))
}
