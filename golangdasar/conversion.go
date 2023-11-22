package main

import "fmt"

func main() {
	var iniNilai32 int32 = 100000
	var iniNilai64 int64 = int64(iniNilai32)
	var iniNilai8 int8 = int8(iniNilai32)

	fmt.Println(iniNilai32)
	fmt.Println(iniNilai64)
	fmt.Println(iniNilai8)

	var name = "Laksa"
	var conv byte = name[2]
	var eString string = string(conv)

	fmt.Println(name)
	fmt.Println(eString)
}