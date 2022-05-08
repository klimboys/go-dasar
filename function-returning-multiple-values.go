package main

import "fmt"

func getFullName() (string, string, int) {
	return "Sony", "MelliFriady", 18
}

func main() {

	//firstname, lastName, age := getFullName()
	firstname, _, age := getFullName()
	fmt.Println(firstname, age)

}
