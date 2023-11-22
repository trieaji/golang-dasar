package main

import "fmt"

type Address struct {
	city, province, country string
}

//pointer method
type Man struct {
	name string
}

//pointer di function
func changeCountry(address *Address) {
	address.country = "Korea"

}

//pointer method
func (man *Man) married() {
	man.name = "Mr. " + man.name
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	// address2 := address1
	address2 := &address1 //menggunakan pointer

	address2.city = "Jakarta"

	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"} //menggunakan * untuk merubah data yang paling awal

	// var address3 *Address = &Address{"Surabaya", "Jawa TImur", "Indonesia"} //pointer untuk membuat data baru

	var address4 *Address = new(Address) //pointer untuk membuat data baru tapi datanya kosong
	// address4.city = "Bandung"


	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address4)

	// pointer di function
	var alamat = Address{
		city: "semarang",
		province: "jawa tengah",
		country: "",
	}
	var alamatPointer *Address = &alamat
	changeCountry(alamatPointer)
	fmt.Println(alamat)

	//pointer method
	identity := Man{"claude"}
	identity.married()
	fmt.Println(identity.name)
}
