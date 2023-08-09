package main

import "fmt"

type Blacklist func(string) bool

func registerusers(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {

	blackedlistdata := func(name string) bool {
		return name == "admin"
	}

	registerusers("admin", blackedlistdata)
	registerusers("root", blackedlistdata)

}
