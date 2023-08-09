package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "rasdi"
	names[1] = "abdulrohman"

	fmt.Println(names[1])
	fmt.Println(len(names))

	var values = [3]int{
		10, 20, 30,
	}

	fmt.Println(values)
}
