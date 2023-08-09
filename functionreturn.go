package main

import "fmt"

func main() {

	result := getHello("rasdi")

	fmt.Println(result)

}

func getHello(name string) string {

	return "Hello " + name
}
