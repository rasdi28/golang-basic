package main

import "fmt"

func Newmap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	person := Newmap("rasdi")
	fmt.Println(person["name"])
}
