package main

import "fmt"

func main() {

	// Check types of writing a for loop!

	x := 0
	for {
		fmt.Println("Do stuff", x)
		x += 3
		if x > 25 {
			break
		}
	}

	for x := 0; x < 25; x += 3 {
		fmt.Println("Do stuff", x)
	}
}
