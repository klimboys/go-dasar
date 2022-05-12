package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Lorem Ipsum", "lorem"))
	fmt.Println(strings.Split("Lorem Ipsum", " "))
	fmt.Println(strings.ToLower("Lorem Ipsum"))
	fmt.Println(strings.ToUpper("Lorem Ipsum"))
	fmt.Println(strings.ToTitle("Lorem Ipsum"))
	fmt.Println(strings.Trim("              Lorem Ipsum  ", " "))
	fmt.Println(strings.ReplaceAll("lorem Ipsum Lorem ", "Lorem", "Lola"))
}
