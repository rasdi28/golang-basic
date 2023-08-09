package main

import "fmt"

func main() {

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	names := []string{"rasdi", "abdul", "rohman"}

	for i, name := range names {
		fmt.Println("index", i, "=", name)
	}

	persion := make(map[string]string)
	persion["nama"] = "rasdi"
	persion["alamat"] = "bogor"

	for a, value := range persion {
		fmt.Println("key", a, "value =", value)
	}

}
