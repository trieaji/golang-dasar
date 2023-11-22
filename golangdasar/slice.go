package main

import "fmt"

func main() {
	var months = [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}
	var slice = months[4:7]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	// months[5] = "Diubah"
	// fmt.Println(slice)

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Tsukuyomi")
	fmt.Println(slice3)

	//membuat slice dari awal banget
	newSlice := make([]string, 2,5)
	newSlice[0] = "Laksa"
	newSlice[1] = "Bayu"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	//kode program copy slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	//perbedaan kode program array vs slice
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}
	
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
