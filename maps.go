package main

import "fmt"

func main() {
	grades := make(map[string]float32)

	grades["TaHa"] = 42
	grades["Jess"] = 54
	grades["Jesus"] = 100

	fmt.Println(grades)

	TaHasGrade := grades["TaHa"]
	fmt.Println(TaHasGrade)

	delete(grades, "Jess")
	fmt.Println(grades)

	//Iterate through a map
	for k, v := range grades {
		fmt.Println(k, v)
	}
}
