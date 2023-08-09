package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "Rasdi",
		"address": "Bogor",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book map[string]string = make(map[string]string)
	book["title"] = "belajar golang"
	book["author"] = "Rasdi"
	book["ups"] = "salah"
	fmt.Println(book)
	fmt.Println(len(book))
	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))

}
