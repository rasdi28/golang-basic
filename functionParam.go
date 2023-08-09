package main

import "fmt"

func main() {
	sayHello("rasdi", "abdulrohman")
	sayHello("joko", "anwar")
}

func sayHello(firstName string, lastName string) {

	fmt.Println("Hai, ", firstName, " ", lastName)
}
