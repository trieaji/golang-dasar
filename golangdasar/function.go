package main

import "fmt"

func sayHello() {
	fmt.Println("Hello world")
}

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

// function return multiple value
func getFullName() (string,string) {
	return "trie", "aji"
}

// function return named value
func getFull() (nameFirst, nameLast string) {
	nameFirst = "clint"
	nameLast = "vexana"

	return
}

//function as value
func getGoodBye(name string) string {
	return "Good Bye " + name
}

//function as parameter
func sayHelloWithFilter(name string, filter func(string) string) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	fmt.Println("=== name ===")
	fmt.Println(name)
	if name == "Anjing" {
		return "Okkkeee"
	} else {
		return name
	}
}

//function recursive(function yang memanggil dirinya sendiri)
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= 1
	}
	return result
}


func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	sayHello()
	sayHelloTo("Laksa", "Bayu")
	/*
		firstName := "Laksa"
		sayHelloTo(firstName, "Bayu")
	*/
	// first, last := getFullName()
	// fmt.Println(first, last)

	//menghiraukan return valuenya
	first, _ := getFullName()
	fmt.Println(first)
	// fmt.Println(getFullName())

	a,b := getFull()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(getFull())

	sayGoodBye := getGoodBye
	result := sayGoodBye("valir")
	fmt.Println(result)

	//function as parameter
	sayHelloWithFilter("Laksa", spamFilter)

	//function recursive
	loop := factorialLoop(5)
	fmt.Println(loop)
	fmt.Println(5*4*3*2*1)

	recursive := factorialRecursive(5)
	fmt.Println(recursive)
}
