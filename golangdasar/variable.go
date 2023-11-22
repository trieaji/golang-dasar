package main

import "fmt"

func main() {
	var name string

	name = "Afgan"
	fmt.Println(name)

	var friendName = "Rossa"
	fmt.Println(friendName)

	//Deklarasi awal tanpa menuliskan "var"
	city := "Jakarta"
	fmt.Println(city)

	city = "Surabaya"
	fmt.Println(city)

	//Multiple variable
	var (
		firstName = "Laksa"
		lastName = "Bayu"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}
