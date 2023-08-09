package main

import "fmt"

func random() interface{} {
	return "hello"
}

//buat typeassertion kita dapat merubah type data asalkan interface nya jelas type datanya
func main() {
	var result interface{} = random()
	var resultString string = result.(string)
	fmt.Println(resultString)
}
