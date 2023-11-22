package main

import "fmt"

func main() {
	var name = "joko"

	if name == "laksa" {
		fmt.Println("Hello Laksa")
	} else if name == "joko" {
		fmt.Println("Hello joko")
	} else {
		fmt.Println("iiih salah")
	}

	//kode program if short statement
	if length := len(name); length > 3 {
		fmt.Println("nama benar")
	} else {
		fmt.Println("nama salah")
	}
}
