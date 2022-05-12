package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args
	fmt.Println(arg)

	name, err := os.Hostname()

	if err == nil {
		fmt.Println(name)
	} else {
		fmt.Println(err.Error())
	}
}
