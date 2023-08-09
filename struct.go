package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {

	var rasdi Customer

	rasdi.Age = 27
	rasdi.Address = "bogor"
	rasdi.Name = "rasdi"

	fmt.Println(rasdi)

	joko := Customer{
		Name:    "joko",
		Address: "jakarta",
		Age:     32,
	}

	fmt.Println(joko)
}
