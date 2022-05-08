package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int8
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello ", name, "My name is", customer.Name)
}

func (a Customer) sayHuuu() {
	fmt.Println("Huuuuuu from", a.Name)
}

func main() {
	var sony Customer
	sony.Name = "Sony"
	sony.Address = "Indonesia"
	sony.Age = 18

	sony.sayHi("Joko ")
	sony.sayHuuu()
	//fmt.Println(sony)
	//
	//joko := Customer{
	//	Name:    "joko",
	//	Address: "Cirebon",
	//	Age:     35,
	//}
	//
	//fmt.Println(joko)
	//
	//budi := Customer{"Budi", "Jakarta", 35}
	//fmt.Println(budi)
}
