package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHello(hasName HasName) {
	fmt.Println("=== tes ===")
	fmt.Println(hasName)
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func main() {
	var profil Person
	profil.Name = "Kaido"

	sayHello(profil)
}