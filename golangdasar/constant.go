package main

import "fmt"

func main() {
	const firstName string = "Laksa"
	const lastName = "Bayu"
	const value = 100

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)

	const (
		country string = "Hongkong"
		city = "Jombang"
	)
	fmt.Println(country)
	fmt.Println(city)
}