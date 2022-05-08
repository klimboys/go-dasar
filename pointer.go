package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {

	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1

	var address3 *Address = &address1

	address2.City = "Bandung"

	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	var address4 *Address = new(Address)
	address4.City = "DKI JAKARTA"
	fmt.Println(address4)

	var alamat = Address{
		City:     "Tasik Malaya",
		Province: "Jawa Barat",
		Country:  "",
	}

	var alamatPointer *Address = &alamat

	ChangeCountryToIndonesia(alamatPointer)
	fmt.Println(alamat)
}
