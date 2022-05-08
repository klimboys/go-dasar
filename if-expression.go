package main

import "fmt"

func main() {
	name := "Sony"

	if name == "Sony" {
		fmt.Println("Hello Sony")
	} else if name == "Joko" {
		fmt.Println("Hello Joki")
	} else {
		fmt.Println("Hi, Kenalan DOng")
	}

	//	short statement
	if length := len(name); length > 5 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
