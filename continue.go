package main

import "fmt"

func main() {
	//	Continue
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("perulangan ke ", i)
	}
}
