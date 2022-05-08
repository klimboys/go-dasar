package main

import "fmt"

func main() {
	name := "Sony"
	counter := 0

	increment := func() {
		name = "zoon"
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(name)
}
