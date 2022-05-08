package main

import "fmt"

func main() {
	//counter := 1
	//
	//for counter <= 10 {
	//	fmt.Println("Perulangan ke", counter)
	//	counter++
	//}

	// initial-statement;, kondisi, post-statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	slice := []string{"Sony", "Melli", "Friady"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	names := []string{"Sony", "Melli", "Friady"}

	for index, value := range names {
		fmt.Println("index", index, "=", value)
		fmt.Println(value)
	}

	//	perulangan dengan map
	person := make(map[string]string)
	person["name"] = "Sony"
	person["hobby"] = "Ngoding"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
