package main

import "fmt"

func main() {
	var months = [...]string{
		"jan", "feb", "mar", "apr", "mei", "jun", "jul", "agus", "sept", "okt", "nov", "des",
	}

	var slice1 = months[4:7]

	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	months[5] = "diubah"

	fmt.Println(slice1)
}
