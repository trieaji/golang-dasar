package main

import "fmt"

func random() interface{} {
	return 100
}

func main() {
	var result interface{} = random()
	//mudah terkena panic
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	//cara aman supaya tidak mudah terkena panic
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("unknown type")
	}
}
