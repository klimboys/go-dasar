package main

import "fmt"

func main() {
	const firstName string = "Sony"
	const lastName = "MelliFriady"
	const age = 17

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(age)

	const (
		namaDepan    string = "Sony"
		namaBelakang        = "MelliFriady"
		umur         int    = 17
	)

	fmt.Println(namaDepan)
	fmt.Println(namaBelakang)
	fmt.Println(umur)
}
