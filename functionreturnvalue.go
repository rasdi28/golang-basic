package main

import "fmt"

func getFullName() (string, string) {

	namaDepan := "rasdi"
	namaBelakang := "abdulrohman"

	return namaDepan, namaBelakang

}

func main() {

	nama1, nama2 := getFullName()

	fmt.Println(nama1)
	fmt.Println(nama2)

	//jika ingin mengambil return 1 value
	ceknama, _ := getFullName()

	fmt.Println(ceknama)

}
