package main

import "fmt"

func random1() interface{} {
	return "hello"
}

func main() {

	result := random1()

	switch value := result.(type) {
	case string:
		fmt.Println("string:", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown")
	}
}
