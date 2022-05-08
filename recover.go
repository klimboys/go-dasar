package main

import "fmt"

func endApp1() {
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message", message)
	}
	fmt.Println("Applikasi selesai")
}

func runApp1(error bool) {
	defer endApp1()
	if error {
		panic("Applikasi error")
	}

	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp1(true)
	fmt.Println("Sony")
}
