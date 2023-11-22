package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "Laksa",
		"city": "Yogya",
	}

	//menambahkan data key baru
	person["title"] = "Programmer"

	//merubah data key
	person["city"] = "Jakarta"

	fmt.Println(person)
	fmt.Println(person["city"])
	fmt.Println(person["title"])

	//membuat map baru
	var book map[string]string = make(map[string]string)
	book["author"] = "Askal"
	book["country"] = "Indonesia"

	//delete book
	// delete(book, "ups")

	fmt.Println(book["country"])
}
