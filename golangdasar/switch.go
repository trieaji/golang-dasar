package main

import "fmt"

func main() {
	name := "lak"

	switch name {
	case "Laksa" :
		fmt.Println("Hello Laksa")

	case "joko" :
		fmt.Println("Hello joko")
	
	default:
		fmt.Println("salah luur")
	}

	//kode program switch short statement
	switch length := len(name); length > 4 {
	case true:
		fmt.Println("nama benar")
	case false:
		fmt.Println("nama salah")
	}
}