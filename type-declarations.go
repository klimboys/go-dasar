package main

import "fmt"

func main() {
	type NoKTP string
	type Merried bool

	var noKtpSony NoKTP = "1234567"
	var merriedStatus Merried = false
	fmt.Println(noKtpSony)
	fmt.Println(merriedStatus)
}
