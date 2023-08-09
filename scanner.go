package main

import "fmt"

func main() {
	inputname()

}

func inputname() {
	fmt.Println("Masukkan nama depan :")
	var first string
	fmt.Scanln(&first)
	fmt.Println("Masukkan nama belakang:")
	var second string
	fmt.Scanln(&second)
	fmt.Println("Nama :", first, second)
	footer()
}

func footer() {
	fmt.Println("Apakah anda ingin kembali ? (Y/N)")
	var isback string
	fmt.Scanln(&isback)
	if isback == "Y" {
		inputname()
	} else {
		fmt.Println("Terima Kasih")
	}

}
