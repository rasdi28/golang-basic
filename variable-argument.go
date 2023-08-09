package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}

	return total
}

func main() {
	jumlah := sumAll(20, 30, 10)

	fmt.Println(jumlah)

	nilais := []int{20, 30, 40, 26}
	jumlah = sumAll(nilais...)
	fmt.Println(jumlah)
}
