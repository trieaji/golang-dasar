package main

import "fmt"

func main() {
	counter := 9

	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}

	//kode program for dengan statement
	for defend := 8; defend <= 10; defend++ {
		fmt.Println("Perulangan ke", defend)
	}

	//kode program for range
	names := []string{"laksa", "bayu", "trie"}
	for index, name := range names{
		fmt.Println("index", index, "=", name)
	}

	//iteration
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}