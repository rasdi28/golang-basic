package main

import "fmt"

func getFullName1() (firstName, middleName, lastName string) {

	firstName = "rasdi"
	lastName = "rohman"
	middleName = "abdul"

	return
}

func main() {
	a, b, c := getFullName1()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
