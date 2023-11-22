package golangdasar

import "fmt"

func main() {
	var names[3] string

	names[0] = "Laksa"
	names[1] = "Bayu"
	names[2] = "Trie"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90,80,95,
	}
	fmt.Println(values[1])
	fmt.Println(len(values))
	values[0] = 100
	fmt.Println(values)
	fmt.Println(values[0])
}