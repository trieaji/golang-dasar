package main

import "fmt"

func main() {
	counter := 0
	name := "laksa"

	increment := func() {
		// name = "bayu"
		name := "bayu" //pendeklarasian ulang
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()

	fmt.Println(counter)
	fmt.Println(name)
}