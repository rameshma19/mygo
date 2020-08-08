package main

import "fmt"

func main() {

	// // var colors map[string]int

	// colors := make(map[string]int32)
	// colors["anan"] = 3

	// delete(colors, "anan")

	ageMap := map[string]int{
		"anan":  3,
		"aarav": 17,
		"hun":   18,
		"self":  19,
	}

	printMap(ageMap)
}

func printMap(a map[string]int) {
	for nm, age := range a {
		fmt.Println("Name: ", nm, " age: ", age)
	}
}
