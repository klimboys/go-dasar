package main

import "fmt"

func main() {
	name := "Sony"

	switch name {
	case "Sony":
		fmt.Println("Hello Sony")
	case "Joko":
		fmt.Println("Hello Joko")

	default:
		fmt.Println("Kenalan dong")
	}

	//	short expression

	//switch length := len(name); length > 5 {
	//case true:
	//	fmt.Println("Nama terlalu panjang")
	//case false:
	//	fmt.Println("Nama sudah benar")
	//}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama Sudah benar")
	}

}
