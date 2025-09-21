package main

import "fmt"

func main() {
	var age int = 30
	name := "Alice"
	const country = "India"

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Country:", country)

	var notInitialized string
	fmt.Println("Default value of notInitialized:", notInitialized)
}
