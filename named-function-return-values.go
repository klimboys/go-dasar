package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Sony"
	middleName = "Melli"
	lastName = "Friady"

	return
}

func main() {
	firstName, middleName, lastName := getCompleteName()
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}
