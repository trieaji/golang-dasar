package main

import "fmt"

func main() {
	type NoKTP string

	var ktp NoKTP = "101010101010"
	fmt.Println(ktp)

	var i = 10
	i++
	fmt.Println(i)
}