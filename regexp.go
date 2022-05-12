package main

import (
	"fmt"
	"regexp"
)

func main() {

	var regex *regexp.Regexp = regexp.MustCompile(`e([a-z])o`)

	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("eGo"))
	fmt.Println(regex.MatchString("eto"))

	search := regex.FindAllString("eko eka edo eyo", -1)
	fmt.Println(search)
}
